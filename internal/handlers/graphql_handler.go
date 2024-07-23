package handlers

import (
	"graphql-sample/internal/services"
	"graphql-sample/pkg/graphql"

	"github.com/gin-gonic/gin"
)

func GraphQLHandler(userService services.UserService) gin.HandlerFunc {
    return graphql.NewGraphQLHandler(userService)
}

func PlaygroundHandler() gin.HandlerFunc {
    return graphql.NewPlaygroundHandler()
}
