package users

import (
	"net/http"

	"../../domain/users"
	"../../services"
	"../../utils/errors"

	"github.com/gin-gonic/gin"
)

// GetUser searches and returns a matching user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// CreateUser creates a new user and returns the result
func CreateUser(c *gin.Context) {
	var user users.User

	// attempt to create a new user structure
	// using the json body
	if err := c.ShouldBindJSON(&user); err != nil {
		// handle json error
		restError := errors.NewBadRequestError(
			"Invalid json body",
		)

		c.JSON(restError.Status, restError)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		// handle user creation error
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result)
}
