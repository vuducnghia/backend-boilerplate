package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addHeartbeat(r *gin.RouterGroup) {
	app := r.Group("/application")
	{
		app.GET("/heartbeat", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}
}
