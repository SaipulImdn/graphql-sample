package main

import (
	"graphql-sample/internal/config"
	"graphql-sample/internal/handlers"
	"graphql-sample/internal/repositories"
	"graphql-sample/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()

    userRepository := repositories.NewUserRepository()
    userService := services.NewUserService(userRepository)

    router := gin.Default()
    router.POST("/query", handlers.GraphQLHandler(userService))
    router.GET("/", handlers.PlaygroundHandler())

    router.Run(":5000")
}
