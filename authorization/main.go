package main

import (
	"database/sql"
	"fmt"
	"librarymanager/authorization/common"
	"librarymanager/authorization/controllers"
	"librarymanager/authorization/domain"
	"librarymanager/authorization/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Print("Authorization process")

	broker := common.NewBroker()

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	userRepository := domain.NewUserRepository(database)

	authorizationService := services.NewAuthorizationService(userRepository, broker)

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
