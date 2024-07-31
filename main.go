package main

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    _ "myapp/docs"
    "myapp/api/handlers"
)

// @title MyApp API
// @version 1.0
// @description This is a simple Gin application with Swagger documentation.
// @host localhost:5050
// @BasePath /
func main() {
    router := gin.Default()

    // Swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // API Endpoints
    router.GET("/api/hello", handlers.Hello)
    router.POST("/api/goodbye", handlers.Goodbye)

    router.Run(":5050")
}
