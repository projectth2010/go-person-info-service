// main.go
package main

import (
	"log"
	"os"

	"go-person-info-service/config"
	"go-person-info-service/middleware"
	"go-person-info-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Person Information Service API
// @version 1.0
// @description A RESTful API for managing person information with authentication
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize Gin
	app := gin.Default()

	// Middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.CORSMiddleware())

	// Connect to database
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup routes
	routes.UserRoutes(app)
	// routes.SetupRoutes(app)

	// Swagger documentation
	if os.Getenv("ENABLE_SWAGGER") != "false" {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
