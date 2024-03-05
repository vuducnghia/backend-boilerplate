package handlers

import (
	application "backend-boilerplate/config"
	"backend-boilerplate/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Login 		godoc
// @Summary		create a user
// @Tags		auth
// @Accept		json
// @Param		user body models.UserCredentials true "user"
// @Success 	200
// @Router		/auth/login [post]
func Login(c *gin.Context) *gin.Error {
	u := &models.User{}
	credentials := &models.UserCredentials{}
	if err := c.ShouldBindJSON(credentials); err != nil {
		return BadRequestError(err, "", c)
	}
	u.Username = credentials.Username
	if err := u.GetByUsername(c.Request.Context()); err != nil {
		return InternalError(err, "invalid username or password", c)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(credentials.Password)); err != nil {
		return InternalError(err, "invalid username or password", c)
	}
	at, err := generateAccessToken(u, 24*time.Hour)
	if err != nil {
		return InternalError(err, "error signing JWT token", c)
	}
	rt, err := generateRefreshToken(u, 24*time.Hour)
	if err != nil {
		return InternalError(err, "error signing JWT token", c)
	}

	a := &models.Auth{
		UserId:       u.Id,
		RefreshToken: models.RefreshToken{RefreshToken: rt},
		AccessToken:  models.AccessToken{AccessToken: at},
	}
	if err := a.Upsert(c); err != nil {
		return DatabaseError(err, "", c)
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  at,
		"refresh_token": rt,
	})
	return nil
}

// RefreshToken godoc
// @Summary		get new token
// @Tags		auth
// @Accept		json
// @Param		user body models.RefreshToken true "user"
// @Success 	200
// @Router		/auth/refresh [post]
func RefreshToken(c *gin.Context) *gin.Error {
	t := &models.RefreshToken{}
	if err := c.ShouldBindJSON(t); err != nil {
		return ValidatorError(err, "error validating refresh token payload", c)
	}

	var claims *models.AuthClaims

	if r, err := validateRefreshToken(t.RefreshToken); err != nil {
		return BadRequestError(err, "invalid refresh token", c)
	} else {
		claims = r.Claims.(*models.AuthClaims)
	}

	u := &models.User{}
	u.Id = claims.UserId
	if err := u.GetById(c.Request.Context()); err != nil {
		return InternalError(err, "invalid username or password", c)

	}
	at, err := generateAccessToken(u, 24*time.Hour)
	if err != nil {
		return InternalError(err, "error signing JWT token", c)
	}
	rt, err := generateRefreshToken(u, 24*time.Hour)
	if err != nil {
		return InternalError(err, "error signing JWT token", c)
	}

	a := &models.Auth{
		UserId:       u.Id,
		RefreshToken: models.RefreshToken{RefreshToken: rt},
		AccessToken:  models.AccessToken{AccessToken: at},
	}
	if err := a.Upsert(c); err != nil {
		return DatabaseError(err, "", c)
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  at,
		"refresh_token": rt,
	})
	return nil
}

func generateRefreshToken(u *models.User, life time.Duration) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.Id,
		"exp":     time.Now().Add(life * 30).Unix(),
	})
	return refreshToken.SignedString([]byte(application.GetConfig().ApplicationConfig.Secret))
}

func generateAccessToken(u *models.User, life time.Duration) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.Id,
		"exp":     time.Now().Add(life).Unix(),
	})
	return accessToken.SignedString([]byte(application.GetConfig().ApplicationConfig.Secret))
}

func validateRefreshToken(t string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(t, &models.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected token signing method")
		}
		return []byte(application.GetConfig().ApplicationConfig.Secret), nil
	})
}
