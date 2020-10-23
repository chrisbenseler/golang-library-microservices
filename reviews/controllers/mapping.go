package controllers

import (
	"librarymanager/reviews/middlewares"

	"github.com/gin-gonic/gin"
)

//MapUrls map routes to controller
func MapUrls(router *gin.Engine, reviewsController Reviews, middleware middlewares.Middleware) *gin.RouterGroup {

	apiRoutes := router.Group("/api/reviews")
	{

		apiRoutes.GET("/books/:id", reviewsController.GetFromBook)

		apiRoutes.POST("/books/:id", middleware.CheckJWTToken, reviewsController.CreateInBook)

	}

	return apiRoutes
}
