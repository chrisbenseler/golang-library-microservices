package main

import (
	"fmt"
	"librarymanager/authorization/common"
	"librarymanager/authorization/controllers"
	"librarymanager/authorization/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type authorizationPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	fmt.Print("Authorization process")

	broker := common.NewBroker()

	usecase := domain.NewUsecase(broker)

	authorizationController := controllers.NewAuthorizationController(usecase)

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
