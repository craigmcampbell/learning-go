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

	controllers.SetupUserRoutes(router)

	return router
}
