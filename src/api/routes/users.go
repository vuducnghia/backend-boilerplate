package routes

import (
	"backend-boilerplate/src/api/handlers"
	"backend-boilerplate/src/api/middleware"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", handlers.Handler(middleware.PagedWrapper(handlers.GetUsers)))
		users.POST("", handlers.Handler(handlers.CreateUser))
		users.GET("/:user_id", handlers.Handler(handlers.GetUser))
		users.PUT("/:user_id", handlers.Handler(handlers.UpdateUser))
		users.DELETE("/:user_id", handlers.Handler(handlers.DeleteUser))
	}
}
