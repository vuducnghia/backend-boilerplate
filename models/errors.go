package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ErrorInt interface {
	Error() string
	GetStatus() int
	Response() gin.H
}
type InternalError struct {
	BaseModel
	Status    int       `json:"status"`
	Method    string    `json:"method"`
	Endpoint  string    `json:"endpoint"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
}
type ValidatorError struct {
	*InternalError
	ValidationErrors map[string]string `json:"validation_errors"`
}

func (e *InternalError) TableName() string {
	return "errors"
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%s for more details refer to: %d", e.Message, e.Id)
}

func (e *InternalError) GetStatus() int {
	return e.Status
}

func (e *InternalError) Response() gin.H {
	return gin.H{"error_id": e.Id, "error": e.Error(), "error_type": e.Type}
}

func (e *ValidatorError) Response() gin.H {
	return gin.H{"error_id": e.Id, "error": e.Error(), "error_type": e.Type, "fields": e.ValidationErrors}
}

func newBadRequest(c string, d string, m string) *InternalError {
	return &InternalError{
		Status:  http.StatusBadRequest,
		Type:    c,
		Details: d,
		Method:  m,
	}
}

func newServerError(c string, d string, m string) *InternalError {
	return &InternalError{
		Status:  http.StatusInternalServerError,
		Type:    c,
		Details: d,
		Method:  m,
	}
}

func NewBadParameterError(e error, m string) *InternalError {
	return newBadRequest("bad_parameter", e.Error(), m)
}

func NewBadRequestError(e error, m string) *InternalError {
	return newBadRequest("request_error", e.Error(), m)
}

func NewDatabaseError(e error, m string) *InternalError {
	return newServerError("database_error", e.Error(), m)
}

func NewEntityNotFoundError(e error, m string) *InternalError {
	return &InternalError{
		Status:  http.StatusNotFound,
		Type:    "entity_not_found",
		Message: m,
		Details: e.Error(),
	}
}

func NewAuthenticationError(e error, m string) *InternalError {
	return &InternalError{
		Status:  http.StatusUnauthorized,
		Type:    "authentication_error",
		Message: m,
		Details: e.Error(),
	}
}

func NewAuthorizationError(e error, m string) *InternalError {
	return &InternalError{
		Status:  http.StatusUnauthorized,
		Type:    "authorization_error",
		Message: m,
		Details: e.Error(),
	}
}

func NewInternalError(e error, m string) *InternalError {
	return newServerError("internal_error", e.Error(), m)
}

func NewValidatorError(e error, m string, fields map[string]string) *ValidatorError {
	return &ValidatorError{
		newBadRequest("validation_error", e.Error(), m), fields,
	}
}
