package middleware

import (
	"backend-boilerplate/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"
)

type InternalError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func ErrorHandler(c *gin.Context) {
	// validate required DTO
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				name = ""
			}
			return name
		})
	}
	c.Next()
	if len(c.Errors) < 1 || (c.IsAborted() && c.Writer.Written()) {
		return
	}

	errorMsg := c.Errors.Last()
	var errorInt models.ErrorInt
	switch {
	case errors.As(errorMsg.Err, &errorInt):
		processError(errorInt, c)
	default:
		processError(&models.InternalError{
			Status:  http.StatusInternalServerError,
			Type:    "unexpected_error",
			Message: "an unexpected error occurred that was not handled",
			Details: errorMsg.Error()}, c)
	}
}

func processError(e models.ErrorInt, c *gin.Context) {
	c.JSON(e.GetStatus(), e.Response())
}
