package main

import (
	"database/sql"
	"fmt"
	"librarymanager/reviews/common"
	"librarymanager/reviews/controllers"
	"librarymanager/reviews/domain"
	"librarymanager/reviews/middlewares"
	"librarymanager/reviews/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type reviewPayload struct {
	Content string `json:"content"`
}

func main() {

	fmt.Print("Reviews process")

	broker := common.NewBroker()

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	repository := domain.NewReviewRepository(database)

	reviewsService := services.NewReviewsService(repository, broker)
	keysService := services.NewKeysService()

	middleware := middlewares.NewMiddleware(keysService)
	reviewsController := controllers.NewReviewsController(reviewsService)

	reviewsService.Subscriptions()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	config.AddExposeHeaders("Authorization")

	router.Use(cors.New(config))

	router.GET("/api/reviewsping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRoutes := controllers.MapUrls(router, reviewsController, middleware)

	apiRoutes.Use(cors.New(config))

	router.Run(":3000")
}
