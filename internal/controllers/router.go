/*
 * Since: August 2022
 * Author: ianmuhia
 * Name: router.go
 * Description:
 *
 * Copyright (c) 2022
 * MirianCode. - All Rights Reserved
 * Unauthorized copying or redistribution of this file in source and binary forms via any medium
 * is strictly prohibited.
 */

package controllers

import (
	"context"
	"ecobake/internal/graph"
	"ecobake/internal/models"
	"ecobake/pkg/resterrors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func (r *Repository) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
	}))
	//router.Use(cors.New(cors.Options{
	//	AllowedOrigins:         nil,
	//	AllowOriginFunc:        nil,
	//	AllowOriginRequestFunc: nil,
	//	AllowedMethods:         nil,
	//	AllowedHeaders:         nil,
	//	ExposedHeaders:         nil,
	//	MaxAge:                 0,
	//	AllowCredentials:       false,
	//	OptionsPassthrough:     false,
	//	OptionsSuccessStatus:   0,
	//	Debug:                  false,
	//}).HandlerFunc())

	_ = router.SetTrustedProxies([]string{"*", "localhost"})

	router.Any("/query", r.graphqlHandler())
	router.GET("/", r.playgroundHandler())

	return router
}

func (r *Repository) graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ServiceProvider: r.serviceProvider,
		FsStorage:       r.storageService,
		NewUser:         make(chan *models.User),
	}}))
	// Configure WebSocket with CORS
	h.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			HandshakeTimeout: 0,
			ReadBufferSize:   1024,
			WriteBufferSize:  1024,
			WriteBufferPool:  nil,
			Subprotocols:     nil,
			Error:            nil,
			CheckOrigin: func(r *http.Request) bool {
				log.Println(r.URL.String())
				log.Println("#########")
				return true
			},
			EnableCompression: false,
		},
		InitFunc:    nil,
		InitTimeout: 0,
		ErrorFunc:   nil,
		//InitFunc:              nil,
		//InitTimeout:           0,
		//ErrorFunc:             nil,
		KeepAlivePingInterval: 1 * time.Second,
		PingPongInterval:      1 * time.Second,
	})

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(lru.New(1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (r *Repository) GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// Defining the Playground handler
func (r *Repository) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (r *Repository) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resterrors.NewError("authorization header is not provided"))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resterrors.NewError("invalid authorization header format"))
			return
		}

		authorizationType := strings.ToLower(fields[0])

		if authorizationType != authorizationTypeBearer {
			err := fmt.Sprintf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		accessToken := fields[1]
		payload, err := r.serviceProvider.VerifyToken(accessToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resterrors.NewError("invalid authorization header format"))
			log.Println(err.Error())
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
