package handlers

import (
	"backend-boilerplate/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHeartbeat godoc
// @Summary		return a status
// @Tags		heartbeat
// @Success		200
// @Router		/application/heartbeat [get]
func GetHeartbeat(c *gin.Context) *gin.Error {
	if err := models.PingDatabase(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "error_type": "database_error"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return nil
}
