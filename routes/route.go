package routes

import (
	"go-person-info-service/controllers"
	"go-person-info-service/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/user", middleware.JWTAuthMiddleware(), controllers.GetUserAccount)
}
