package main

import (
	"go-person-info-service/config"
	_ "go-person-info-service/docs"
	"go-person-info-service/routes"

	 "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	routes.UserRoutes(r)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
