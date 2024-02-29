package handlers

import (
	"backend-boilerplate/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser	godoc
// @Summary	return a user
// @Tags	users
// @Param	user_id path string true "user_id"
// @Success	200
// @Router	/users/{user_id} [get]
func GetUser(c *gin.Context) *gin.Error {
	u := &models.User{}
	u.Id = GetUUIDFromPath("user_id", c)
	if u.Id == "" {
		return BadParameterError(BadRequestParameter, "user_id", c)
	}
	if err := u.GetById(); err != nil {
		return DatabaseError(err, "the user could not be found", c)
	}

	c.JSON(http.StatusOK, u)
	return nil
}

// GetUsers 	godoc
// @Summary 	return list users
// @Tags 		users
// @Success 	200
// @Router		/users [get]
func GetUsers(c *gin.Context) *gin.Error {
	u := models.Users{}
	err := u.GetAll()
	if err != nil {
		return c.Error(DatabaseError(err, "could not retrieve user list", c))
	}

	c.JSON(http.StatusOK, u)
	return nil
}

// UpdateUser	godoc
// @Summary		update a user
// @Tags 		users
// @Accept       json
// @Param		user_id path string true "user_id"
// @Param		user body models.User true "user"
// @Success 	200
// @Router		/users/{user_id} [put]
func UpdateUser(c *gin.Context) *gin.Error {
	u := &models.User{}
	u.Id = c.Param("user_id")
	if u.Id == "" {
		return BadParameterError(BadRequestParameter, "user_id", c)
	}

	if err := u.GetById(); err != nil {
		return EntityNotFoundError(err, "user", c)
	}
	if err := c.ShouldBindJSON(u); err != nil {
		return ValidatorError(err, "error validating user entity", c)
	}
	if err := u.Update(); err != nil {
		return DatabaseError(err, "", c)
	}
	c.JSON(http.StatusOK, u)
	return nil
}
