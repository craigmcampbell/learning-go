package controllers

import (
	"cc.tech/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/v1/users")
	{
		userGroup.GET("", services.GetUsers)
		userGroup.GET("/:id", services.GetUserById)
	}
}