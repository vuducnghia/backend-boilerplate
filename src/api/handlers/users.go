package handlers

import (
	"backend-boilerplate/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
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
	if err := u.GetById(c.Request.Context()); err != nil {
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
	u := &models.Users{}

	if err := u.GetAll(c.Request.Context()); err != nil {
		return c.Error(DatabaseError(err, "could not retrieve user list", c))
	}

	c.JSON(http.StatusOK, u)
	return nil
}

// UpdateUser	godoc
// @Summary		update a user
// @Tags 		users
// @Accept      json
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

	if err := u.GetById(c.Request.Context()); err != nil {
		return EntityNotFoundError(err, "user", c)
	}
	if err := c.ShouldBindJSON(u); err != nil {
		return ValidatorError(err, "error validating user entity", c)
	}
	if err := u.Update(c.Request.Context()); err != nil {
		return DatabaseError(err, "", c)
	}
	c.JSON(http.StatusOK, u)
	return nil
}

// CreateUser 	godoc
// @Summary		create a user
// @Tags		users
// @Accept		json
// @Param		user body models.User true "user"
// @Success 	200
// @Router		/users [post]
func CreateUser(c *gin.Context) *gin.Error {
	u := &models.User{}
	p := &models.UserPassword{}
	if err := c.ShouldBindBodyWith(u, binding.JSON); err != nil {
		return ValidatorError(err, "error validating user entity", c)
	}
	if err := c.ShouldBindBodyWith(p, binding.JSON); err != nil {
		return ValidatorError(err, "error validating user entity", c)
	}
	if hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), 10); err != nil {
		return InternalError(err, "error creating the hashed password", c)
	} else {
		u.Password = string(hash)
	}
	if err := u.Create(c.Request.Context()); err != nil {
		return DatabaseError(err, "", c)
	}
	c.JSON(http.StatusOK, u)
	return nil
}

// DeleteUser 	godoc
// @Summary		delete a user
// @Tags 		users
// @Accept      json
// @Param		user_id path string true "user_id"
// @Success 	200
// @Router		/users/{user_id} [delete]
func DeleteUser(c *gin.Context) *gin.Error {
	u := &models.User{}
	u.Id = c.Param("user_id")
	if u.Id == "" {
		return BadParameterError(BadRequestParameter, "user_id", c)
	}

	if err := u.Delete(c.Request.Context()); err != nil {
		return DatabaseError(err, "", c)
	}

	c.Status(http.StatusNoContent)
	return nil
}
