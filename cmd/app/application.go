package app

import (
	"context"
	"ecobake/cmd/config"
	"ecobake/cmd/internal"
	"ecobake/ent"
	"ecobake/internal/controllers"
	"ecobake/internal/graph"
	"ecobake/internal/graph/generated"
	"ecobake/internal/services"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Output   output
	Database database
	Minio    minio
	Secrets  secrets
	Redis    redis
	Server   server
}

type database struct {
	Server   string
	Port     string
	Database string
	User     string
	Password string
}

type minio struct {
	Server   string
	Port     string
	Name     string
	User     string
	Password string
}

type redis struct {
	Server string
	Port   string
	DB     int
}

type server struct {
	Address string
	Port    string
}

type output struct {
	Directory string
	Format    string
}

type secrets struct {
	TTL int
	Jwt string
}

func initConfig(logger *log.Logger) (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		logger.Println(err)
		return conf, err
	}
	return conf, nil
}

func StartApplication() {
	// configure logger
	logger := log.New(os.Stdout, "[ECOBAKE] [DEBUG] ", log.LstdFlags|log.Lshortfile)
	conf, err := initConfig(logger)
	if err != nil {
		log.Println(err)
		return
	}
	minioURL := fmt.Sprintf("%s:%s", conf.Minio.Server, conf.Minio.Port)
	// Get file storage connection
	minioClient, bucketName, minioErr := internal.NewMinioConnection(
		conf.Minio.User,
		conf.Minio.Password,
		minioURL,
		conf.Minio.Name,
		logger)
	if minioErr != nil {
		logger.Println(minioErr)
		return
	}

	logger.Printf("mino endpoint is 🍉 %s", minioClient.EndpointURL())

	redisURL := fmt.Sprintf("%s:%s", conf.Redis.Server, conf.Redis.Port)

	redisConn := internal.GetRedisClient(redisURL, conf.Redis.DB, logger)

	searchClient := internal.GetMeiliConn()
	searchCli := internal.NewMeiliSearch(searchClient)

	pgConn := config.PostgresConn{
		URL:    conf.Database.Server,
		DBName: conf.Database.Database,
		DBUser: conf.Database.User,
		DBPass: conf.Database.Password,
		Port:   conf.Database.Port,
	}

	// connect to nats
	jetStreamContext := internal.GetJetStream("", logger)

	zincChan := make(chan any)
	zincRcvChan := make(chan any)
	cfg := config.NewAppConfig(
		logger,
		minioClient.EndpointURL(),
		bucketName,
		20,
		false,
		20,
		redisConn,
		jetStreamContext,
		pgConn,
		zincChan,
		zincRcvChan,
	)

	pgPool, err := internal.NewPostgreSQL(pgConn)
	if err != nil {
		logger.Println(err)
		return
	}
	defer pgPool.Close()
	logger.Println(err)

	// initiate services
	client, err := internal.EntConn(pgConn)
	if err != nil {
		logger.Fatal(err)
	}
	mailService := services.NewMailService()
	storageService := services.NewFileStorageService(bucketName, minioClient)
	userService := services.NewUsersService(pgPool, cfg, client)
	tokenService, err := services.NewTokenService(conf.Secrets.Jwt, cfg)
	if err != nil {
		logger.Println(err)
		return
	}
	natsService := services.NewNatsService(cfg)
	searchService := services.NewSearchService(cfg, searchCli.MeiliClient)
	categoryService := services.NewCategoriesService(pgPool, cfg)
	//paymentService := services.NewPaymentsService(cfg, pgPool)

	allServices := controllers.NewRepo(
		mailService,
		cfg,
		storageService,
		natsService,
		userService,
		tokenService,
		searchService,
		categoryService,
		client,
		//paymentService,
	)

	r := allServices.SetupRouter()

	//userService.CleanDB()
	// Run the server
	Run(
		client,
		logger,
		conf.Server.Port,
		conf.Server.Address,
		r,
		storageService,
		tokenService,
		natsService,
		userService,
		searchService,
		categoryService,
	)

}

func Run(
	client *ent.Client,
	logger *log.Logger,
	port string,
	address string,
	r http.Handler,
	storageService services.FileStorageService,
	tokenService services.TokenService,
	natsService services.NatsService,
	userService services.UsersService,
	searchService services.SearchService,
	categoryService services.CategoriesService,
) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	resolver := graph.NewResolver(
		client,
		storageService,
		natsService,
		userService,
		tokenService,
		searchService,
		categoryService,
	)

	//httpServerURL := fmt.Sprintf("%s:%s", conf.Server.Address, conf.Server.Port)
	gsrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	// Configure WebSocket with CORS
	gsrv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},

		KeepAlivePingInterval: 1 * time.Second,
		PingPongInterval:      1 * time.Second,
	})

	gsrv.AddTransport(transport.Options{})
	gsrv.AddTransport(transport.Websocket{})
	gsrv.AddTransport(transport.GET{})
	gsrv.AddTransport(transport.POST{})
	gsrv.AddTransport(transport.MultipartForm{})
	gsrv.SetQueryCache(lru.New(1000))
	gsrv.Use(extension.Introspection{})
	gsrv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	srvURL := fmt.Sprintf("%s:%s", address, port)

	srv := &http.Server{
		Addr:    srvURL,
		Handler: r,
		//ReadHeaderTimeout: time.Second,
		//WriteTimeout:      time.Second,
		//IdleTimeout:       time.Second,
		ErrorLog: logger,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)

		}
	}()
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.AllowAll().Handler(gsrv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	// Listen for the interrupt signal.
	<-ctx.Done()

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logger.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	etime := time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), etime)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: ", err)
	}
	logger.Println("Server exiting")

}
