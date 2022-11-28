package app

import (
	"context"
	"ecobake/cmd/config"
	"ecobake/cmd/internal"
	"ecobake/internal/services"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database database
	Minio    minio
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

type server struct {
	Address string
	Port    string
}

func initConfig(logger *log.Logger) (Config, error) {
	var conf Config
	wordPtr := flag.String("c", ".", "config file location")
	flag.Parse()
	if _, err := toml.DecodeFile(*wordPtr, &conf); err != nil {
		return conf, err
	}
	return conf, nil
}

func StartApplication() error {
	// configure logger
	logger := log.New(os.Stdout, "[ECOBAKE] [DEBUG] ", log.LstdFlags|log.Lshortfile)
	conf, err := initConfig(logger)
	if err != nil {
		return err
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
		return err
	}

	logger.Printf("mino endpoint is 🍉 %s", minioClient.EndpointURL())

	pgConn := config.PostgresConn{
		URL:    conf.Database.Server,
		DBName: conf.Database.Database,
		DBUser: conf.Database.User,
		DBPass: conf.Database.Password,
		Port:   conf.Database.Port,
	}

	// connect to nats

	cfg := config.NewAppConfig(
		logger,
		minioClient.EndpointURL(),
		bucketName,
		20,
		false,
		20,
		pgConn,
	)
	// initiate services
	client, err := internal.EntConn(pgConn)
	if err != nil {
		return err
	}
	storageService := services.NewFileStorageService(bucketName, minioClient)
	userService := services.NewUsersService(cfg, client)
	categoryService := services.NewCategoriesService(cfg)
	tokenService, err := services.NewTokenService("12345")
	if err != nil {
		return err
	}
	allServices := handlers.NewRepo(
		cfg,
		storageService,
		userService,
		tokenService,
		categoryService,
		client,
	)

	r := allServices.SetupRouter()
	// Run the server
	Run(
		logger,
		conf.Server.Port,
		conf.Server.Address,
		r,
	)
	return nil
}

func Run(
	logger *log.Logger,
	port string,
	address string,
	r http.Handler,

) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srvURL := fmt.Sprintf("%s:%s", address, port)

	srv := &http.Server{
		Addr:              srvURL,
		Handler:           r,
		ErrorLog:          logger,
		ReadTimeout:       2 * time.Second,
		WriteTimeout:      2 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)

		}
	}()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logger.Println("shutting down gracefully, press Ctrl+C again to force")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	etime := time.Second * 5
	ctx, cancel := context.WithTimeout(ctx, etime)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: ", err)
	}
	logger.Println("Server exiting")

}
