package middleware

import (
	application "backend-boilerplate/config"
	"backend-boilerplate/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"net/http"
)

func validateToken(c *gin.Context) (*jwt.Token, error) {
	return request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, keyLookupFunc, request.WithClaims(&models.AuthClaims{}))
}

func keyLookupFunc(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected token signing method")
	}
	return []byte(application.GetConfig().ApplicationConfig.AccessToken), nil
}

func AuthorizationHandler(c *gin.Context) {
	if t, err := validateToken(c); err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	} else {
		claims := t.Claims.(*models.AuthClaims)
		if claims.UserId == "" {
			_ = c.AbortWithError(http.StatusUnauthorized, errors.New("invalid key type used as authorization key"))
			return
		}
		c.Set("current_user_id", claims.UserId)
	}
	c.Next()
}
