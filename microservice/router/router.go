package router

import (
	"net/http"

	"cc.tech/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add this line to handle favicon requests
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.GetUserById)

	return router
}
