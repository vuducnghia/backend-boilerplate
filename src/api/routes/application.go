package routes

import (
	"backend-boilerplate/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func addHeartbeat(r *gin.RouterGroup) {
	app := r.Group("/application")
	{
		app.GET("/heartbeat", handlers.Handler(handlers.GetHeartbeat))
	}
}
