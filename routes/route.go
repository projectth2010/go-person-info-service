package routes

import (
	"go-person-info-service/controllers"
	"go-person-info-service/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	// Public routes
	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.RegisterUser)
			auth.POST("/login", controllers.LoginUser)
		}

		// Protected routes
		users := api.Group("/users")
		users.Use(middleware.JWTAuthMiddleware())
		{
			// User profile management
			users.GET("/me", controllers.GetUserAccount)
			users.PUT("/me", controllers.UpdateUserProfile)
			users.DELETE("/me", controllers.DeleteUserAccount)
			
			// Password management
			users.POST("/me/change-password", controllers.ChangePassword)
		}
	} 
}
