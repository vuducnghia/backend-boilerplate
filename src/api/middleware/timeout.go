package middleware

import (
	application "backend-boilerplate/config"
	"backend-boilerplate/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Timeout(c *gin.Context) {
	if strings.Contains(c.Request.RequestURI, "swagger") {
		c.Next()
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(),
		time.Duration(application.GetConfig().ApplicationConfig.RequestTimeout)*time.Second)
	defer cancel()
	c.Request = c.Request.Clone(ctx)

	go func() {
		select {
		case <-ctx.Done():
			processError(&models.InternalError{
				Status:  http.StatusRequestTimeout,
				Type:    "request_timeout",
				Message: "request timeout",
				Details: fmt.Sprintf("%s", c.Request.Context().Err()),
			}, c)
		}
	}()
	c.Next()
}
