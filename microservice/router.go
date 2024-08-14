package main

import (
	"net/http"

	"cc.tech/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add this line to handle favicon requests
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	// router.GET("/users", users.GetUsers)
	// router.GET("/users/:id", users.GetUserById)

	controllers.SetupUserRoutes(router)

	return router
}
