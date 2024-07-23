package graphql

import (
	"graphql-sample/internal/services"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func NewGraphQLHandler(userService services.UserService) gin.HandlerFunc {
    h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{UserService: userService}}))
    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}

func NewPlaygroundHandler() gin.HandlerFunc {
    h := playground.Handler("GraphQL", "/query")
    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}
