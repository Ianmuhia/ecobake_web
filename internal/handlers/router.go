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
	"ecobake/internal/graph/directives"
	"ecobake/internal/graph/generated"
	"ecobake/internal/models"
	"ecobake/pkg/resterrors"
	"fmt"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"

	"time"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
)

func (r *Repository) SetupRouter() chi.Router {
	router := chi.NewRouter()
	//uses chi middleware logic.
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	//Sets context for all requests
	router.Use(middleware.Timeout(20 * time.Second))
	router.Use(r.AuthMiddleware)
	router.Handle("/query", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// NewExecutableSchema and Config are in the generated.go file.
		// Resolver is in the resolver.go file.
		h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
			StorageService:  r.storageService,
			UserService:     r.userService,
			TokenService:    r.tokenService,
			CategoryService: r.CategoryService,
			Client:          r.Client,
		}, Directives: struct {
			Auth    func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error)
			HasRole func(ctx context.Context, obj any, next graphql.Resolver, role models.Role) (res any, err error)
		}{Auth: directives.Auth}}))
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
					return true
				},
				EnableCompression: true,
			},
			InitFunc:    nil,
			InitTimeout: 0,
			ErrorFunc:   nil,

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
	}))
	router.Handle("/", r.playgroundHandler())

	return router
}

// func (r *Repository) graphqlHandler() func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

// 			// NewExecutableSchema and Config are in the generated.go file.
// 			// Resolver is in the resolver.go file.
// 			h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
// 				StorageService:  r.storageService,
// 				UserService:     r.userService,
// 				TokenService:    r.tokenService,
// 				CategoryService: r.CategoryService,
// 				Client:          r.Client,
// 			}, Directives: struct {
// 				Auth    func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error)
// 				HasRole func(ctx context.Context, obj any, next graphql.Resolver, role models.Role) (res any, err error)
// 			}{Auth: directives.Auth}}))
// 			// Configure WebSocket with CORS
// 			h.AddTransport(&transport.Websocket{
// 				Upgrader: websocket.Upgrader{
// 					HandshakeTimeout: 0,
// 					ReadBufferSize:   1024,
// 					WriteBufferSize:  1024,
// 					WriteBufferPool:  nil,
// 					Subprotocols:     nil,
// 					Error:            nil,
// 					CheckOrigin: func(r *http.Request) bool {

// 						return true
// 					},
// 					EnableCompression: false,
// 				},
// 				InitFunc:    nil,
// 				InitTimeout: 0,
// 				ErrorFunc:   nil,

// 				KeepAlivePingInterval: 1 * time.Second,
// 				PingPongInterval:      1 * time.Second,
// 			})

// 			h.AddTransport(transport.Options{})
// 			h.AddTransport(transport.GET{})
// 			h.AddTransport(transport.POST{})
// 			h.AddTransport(transport.MultipartForm{})

// 			h.SetQueryCache(lru.New(1000))

// 			h.Use(extension.Introspection{})
// 			h.Use(extension.AutomaticPersistedQuery{
// 				Cache: lru.New(100),
// 			})
// 			next.ServeHTTP(w, req)
// 		})
// 	}
// }

// Defining the Playground handler.

func (r *Repository) playgroundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playground.Handler("GraphQL", "/query")
	}
}

// HTTP middleware setting a value on the request context.
func (r *Repository) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get(authorizationHeaderKey)

		if len(authorizationHeader) != 0 {
			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 {
				next.ServeHTTP(w, req)
			}

			authorizationType := strings.ToLower(fields[0])

			if authorizationType != authorizationTypeBearer {
				err := fmt.Sprintf("unsupported authorization type %s", authorizationType)
				render.JSON(w, req, err)
				return
			}

			accessToken := fields[1]
			payload, err := r.tokenService.VerifyToken(accessToken)

			if err != nil {
				render.JSON(w, req, resterrors.NewError("invalid authorization header format"))
				return
			}
			// put it in context
			//ctx := context.WithValue(gin.Context, userCtxKey, user)

			// and call the next with our new context
			// r := req.WithContext(req.Context())
			// ctx := context.WithValue(re, authorizationPayloadKey, ctx)
			// //ctx.Request = ctx.Request.WithContext(ctx)

			// ctx.Set(authorizationPayloadKey, payload)
			// ctx.Next()
			ctx := context.WithValue(req.Context(), "user", payload)

			//   // call the next handler in the chain, passing the response writer and
			//   // the updated request object with the new context value.
			//   //
			//   // note: context.Context values are nested, so any previously set
			//   // values will be accessible as well, and the new `"user"` key
			//   // will be accessible from this point forward.
			next.ServeHTTP(w, req.WithContext(ctx))
		}
		// ctx.Next()

		//   // create new context from `r` request context, and assign key `"user"`
		//   // to value of `"123"`
		ctx := context.WithValue(req.Context(), "user", "123")

		//   // call the next handler in the chain, passing the response writer and
		//   // the updated request object with the new context value.
		//   //
		//   // note: context.Context values are nested, so any previously set
		//   // values will be accessible as well, and the new `"user"` key
		//   // will be accessible from this point forward.
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
