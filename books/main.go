package main

import (
	"database/sql"
	"fmt"
	"librarymanager/books/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type bookPayload struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func main() {

	fmt.Print("Books process with broker")

	service := domain.NewBookService()

	broker := domain.NewBroker()

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	repository := domain.NewBookRepository(service, database)

	usecase := domain.NewBookUsecase(repository, broker)
	middleware := domain.NewMiddleware(service)
	controller := domain.NewController(usecase)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	config.AddExposeHeaders("Authorization")
	router.Use(cors.New(config))

	router.GET("/api/booksping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRoutes := router.Group("/api/books")
	{

		apiRoutes.GET("/:id", controller.GetByID)

		apiRoutes.GET("/", controller.All)

		apiRoutes.POST("/", middleware.CheckJWTToken, controller.Create)

		apiRoutes.DELETE("/:id", middleware.CheckJWTToken, controller.Delete)

	}

	apiRoutes.Use(cors.New(config))

	router.Run(":3000")
}
