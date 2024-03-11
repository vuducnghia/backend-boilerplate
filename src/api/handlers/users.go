package handlers

import (
	"backend-boilerplate/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// GetUser		godoc
// @Summary		return a user
// @Tags		users
// @Param		user_id path string true "user_id"
// @Success		200
// @Router		/users/{user_id} [get]
// @Security 	Bearer
func GetUser(c *gin.Context) *gin.Error {
	u := &models.User{}
	u.Id = GetUUIDFromPath("user_id", c)
	if u.Id == "" {
		return BadParameterError(BadRequestParameter, "user_id", c)
	}
	if err := u.GetById(c); err != nil {
		return DatabaseError(err, "the user could not be found", c)
	}

	c.JSON(http.StatusOK, u)
	return nil
}

// GetUsers 	godoc
// @Summary 	return list users
// @Tags 		users
// @Param       limit    		query     string  false  "limit"
// @Param       page_number    	query     string  false  "page_number"
// @Param       search_query    query     string  false  "search_query"
// @Success 	200
// @Router		/users [get]
// @Security 	Bearer
func GetUsers(c *gin.Context) *gin.Error {
	u := &models.Users{}
	w := GetPaginationVariables(c)
	if count, err := u.GetAll(c); err != nil {
		return c.Error(DatabaseError(err, "could not retrieve user list", c))
	} else {
		w = models.NewPaginationWrapper(u, count, c)
	}

	c.JSON(http.StatusOK, w)
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
// @Security 	Bearer
func UpdateUser(c *gin.Context) *gin.Error {
	u := &models.User{}
	u.Id = c.Param("user_id")
	if u.Id == "" {
		return BadParameterError(BadRequestParameter, "user_id", c)
	}

	if err := u.GetById(c); err != nil {
		return EntityNotFoundError(err, "user", c)
	}
	if err := c.ShouldBindJSON(u); err != nil {
		return ValidatorError(err, "error validating user entity", c)
	}
	if err := u.Update(c); err != nil {
		return DatabaseError(err, "", c)
	}
	c.JSON(http.StatusOK, u)
	return nil
}

// CreateUser 	godoc
// @Summary		create a user
// @Tags		auth
// @Accept		json
// @Param		user body models.User true "user"
// @Success 	200
// @Router		/auth/register [post]
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
	if err := u.Create(c); err != nil {
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
// @Security 	Bearer
func DeleteUser(c *gin.Context) *gin.Error {
	u := &models.User{}
	u.Id = c.Param("user_id")
	if u.Id == "" {
		return BadParameterError(BadRequestParameter, "user_id", c)
	}

	if err := u.Delete(c); err != nil {
		return DatabaseError(err, "", c)
	}

	c.Status(http.StatusNoContent)
	return nil
}
