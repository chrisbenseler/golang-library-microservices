package controllers

import (
	"github.com/gin-gonic/gin"
)

//MapUrls map routes to controller
func MapUrls(router *gin.Engine, authorizationController Authorization) *gin.RouterGroup {

	apiRoutes := router.Group("/api/authorization")
	{
		apiRoutes.POST("/signin", authorizationController.SignIn)
		apiRoutes.POST("/signup", authorizationController.SignUp)
	}

	return apiRoutes
}
