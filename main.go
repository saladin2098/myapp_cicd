package main

import (
	"myapp/api/handlers"
	_ "myapp/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyApp API
// @version 1.0
// @description This is a simple Gin application with Swagger documentation.
// @host 18.195.148.52:5050
// @BasePath /
func main() {
	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// API Endpoints
	router.GET("/api/hello", handlers.Hello)
	router.POST("/api/goodbye", handlers.Goodbye)

	router.Run(":5050")
}
