package routes

import (
	"backend-boilerplate/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func addAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", handlers.Handler(handlers.Login))
		auth.POST("/refresh", handlers.Handler(handlers.RefreshToken))
	}
}
