package app

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/classroom/internal/graphql/graph/generated"
	"github.com/lucasd-coder/classroom/internal/graphql/resolvers"
)

func graphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers:  &resolvers.Resolver{},
			Directives: generated.DirectiveRoot{},
			Complexity: generated.ComplexityRoot{},
		},
	))

	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	srv := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func Run() {
	// Setting up Gin
	srv := gin.Default()
	srv.Use(gin.Recovery())
	srv.POST("/graphql", graphqlHandler())
	srv.GET("/", playgroundHandler())
	err := srv.Run(":" + "3334")
	if err != nil {
		panic(err)
	}
}
