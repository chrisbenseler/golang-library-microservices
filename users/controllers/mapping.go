package controllers

import (
	"github.com/gin-gonic/gin"
)

//MapUrls map routes to controller
func MapUrls(router *gin.Engine, usersController Users) *gin.RouterGroup {

	apiRoutes := router.Group("/api/users")
	{
		apiRoutes.GET("/:id", usersController.Get)
	}

	return apiRoutes
}
