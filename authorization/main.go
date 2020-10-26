package main

import (
	"fmt"
	"librarymanager/authorization/common"
	"librarymanager/authorization/controllers"
	"librarymanager/authorization/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Print("Authorization process")

	broker := common.NewBroker()

	authorizationService := services.NewAuthorizationService(broker)

	authorizationController := controllers.NewAuthorizationController(authorizationService)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	config.AddExposeHeaders("Authorization")
	router.Use(cors.New(config))

	apiRoutes := controllers.MapUrls(router, authorizationController)

	apiRoutes.Use(cors.New(config))

	router.Run(":3000")
}
