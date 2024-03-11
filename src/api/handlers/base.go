package handlers

import (
	"backend-boilerplate/models"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

var (
	BadRequest          = errors.New("invalid request")
	BadRequestParameter = errors.New("a required path parameter is missing")
)

func Handler(h func(c *gin.Context) *gin.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			c.Abort()
		}
	}
}

func BadParameterError(e error, m string, c *gin.Context) *gin.Error {
	return c.Error(models.NewBadParameterError(e, m))
}

func BadRequestError(e error, m string, c *gin.Context) *gin.Error {
	return c.Error(models.NewBadRequestError(e, m))
}

func DatabaseError(e error, m string, c *gin.Context) *gin.Error {
	return c.Error(models.NewDatabaseError(e, m))
}

func AuthenticationError(e error, m string, c *gin.Context) *gin.Error {
	return c.Error(models.NewAuthenticationError(e, m))
}

func AuthorizationError(e error, m string, c *gin.Context) *gin.Error {
	return c.Error(models.NewAuthorizationError(e, m))
}

func EntityNotFoundError(e error, m string, c *gin.Context) *gin.Error {
	return c.Error(models.NewEntityNotFoundError(e, fmt.Sprintf("the %s could not be found", m)))
}

func InternalError(e error, message string, c *gin.Context) *gin.Error {
	return c.Error(models.NewInternalError(e, message))
}

func ValidatorError(e error, m string, c *gin.Context) *gin.Error {
	fields := make(map[string]string)
	var parseError *time.ParseError
	var validationErrors validator.ValidationErrors
	switch {
	case errors.As(e, &parseError):
		return c.Error(models.NewBadRequestError(e, "please refer to the api documentation for proper datetime formats"))
	case errors.As(e, &validationErrors):
		for _, fieldErr := range e.(validator.ValidationErrors) {
			fields[fieldErr.Field()] = fieldErr.Tag() + " - " + fieldErr.Namespace()
		}
		return c.Error(models.NewValidatorError(e, m, fields))
	default:
		return c.Error(models.NewBadRequestError(e, "unspecified error occurred with the request binding"))
	}
}

func GetIDFromPath(k string, c *gin.Context) int32 {
	if p := c.Param(k); p != "" {
		if i, err := strconv.ParseInt(p, 10, 32); err != nil {
			return 0
		} else {
			return int32(i)
		}
	}
	return 0
}

func GetUUIDFromPath(k string, c *gin.Context) string {
	return c.Param(k)
}

func GetPaginationVariables(c *gin.Context) *models.PaginationWrapper {
	var pw *models.PaginationWrapper
	context.WithValue(c, "page_number", c.GetInt("page_number"))
	context.WithValue(c, "limit", c.GetInt("limit"))
	context.WithValue(c, "search_query", c.GetInt("search_query"))
	return pw
}
