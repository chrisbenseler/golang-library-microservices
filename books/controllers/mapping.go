package controllers

import (
	"librarymanager/books/middlewares"

	"github.com/gin-gonic/gin"
)

//MapUrls map routes to controller
func MapUrls(router *gin.Engine, booksController Books, middleware middlewares.Middleware) *gin.RouterGroup {

	apiRoutes := router.Group("/api/books")
	{

		apiRoutes.GET("/:id", booksController.GetByID)

		apiRoutes.GET("/", booksController.All)

		apiRoutes.POST("/", middleware.CheckJWTToken, booksController.Create)

		apiRoutes.DELETE("/:id", middleware.CheckJWTToken, booksController.Delete)

	}

	return apiRoutes
}
