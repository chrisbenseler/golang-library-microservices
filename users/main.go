package main

import (
	"database/sql"
	"fmt"
	"librarymanager/users/common"
	"librarymanager/users/domain"
	"librarymanager/users/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Println("Users process")

	broker := common.NewBroker()

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	repository := domain.NewUsersRepository(database)

	usersService := services.NewUsersService(repository, broker)

	repository.Initialize()

	usersService.Subscriptions()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	config.AddExposeHeaders("Authorization")

	router.Use(cors.New(config))

	router.GET("/api/usersping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":3001")

}
