package users

import (
	"fmt"
	"net/http"

	"cc.tech/models"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{Id: "1", Name: "John", Email: "john@example.com"},
	{Id: "2", Name: "Doe", Email: "doe@example.com"},
}

func AddUser(c *gin.Context) {
	var newUser models.User

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&newUser); err != nil {
			return
	}

	// Add the new album to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, u := range users {
			if u.Id == id {
					c.IndentedJSON(http.StatusOK, u)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("user with id %s not found", id)})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}