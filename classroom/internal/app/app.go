package app

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lucasd-coder/classroom/internal/graphql/graph/generated"
	"github.com/lucasd-coder/classroom/internal/graphql/resolvers"
	"github.com/lucasd-coder/classroom/internal/middlewares"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/lucasd-coder/classroom/internal/tools"
)

func graphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers:  &resolvers.Resolver{},
			Directives: generated.DirectiveRoot{},
			Complexity: generated.ComplexityRoot{},
		},
	))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:   1024,
			WriteBufferSize:  1024,
			HandshakeTimeout: 5 * time.Second,
		},
		KeepAlivePingInterval: 10 * time.Second,
		PingPongInterval:      time.Second,
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			token, err := tools.EnsureValidToken(initPayload.Authorization(), ctx)
			if err != nil {
				logger.Log.Errorf("unable to initialise websocket connection", err)
				return nil, tools.ErrUnAuthorized
			}

			claims, ok := token.(*validator.ValidatedClaims)
			if !ok {
				logger.Log.Warn("unexpected token format")
				return nil, tools.ErrUnAuthorized
			}

			return context.WithValue(ctx, jwtmiddleware.ContextKey{}, claims), nil
		},
	})

	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	srv := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func Run(port string) {
	logger.SetUpLog()
	srv := gin.Default()
	srv.Use(gin.Recovery())
	srv.Use(middlewares.GinContextToContextMiddleware())
	srv.POST("/graphql", graphqlHandler())
	srv.GET("/playground", playgroundHandler())
	err := srv.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
