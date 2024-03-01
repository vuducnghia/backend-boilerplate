package routes

import (
	"backend-boilerplate/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", handlers.Handler(handlers.GetUsers))
		users.POST("", handlers.Handler(handlers.CreateUser))
		users.GET("/:user_id", handlers.Handler(handlers.GetUser))
		users.PUT("/:user_id", handlers.Handler(handlers.UpdateUser))
		users.DELETE("/:user_id", handlers.Handler(handlers.DeleteUser))
	}
}
