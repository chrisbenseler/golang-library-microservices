package main

import (
	"database/sql"
	"fmt"
	"librarymanager/books/controllers"
	"librarymanager/books/domain"
	"librarymanager/books/middlewares"
	"librarymanager/books/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Print("Books process with broker")

	keysService := services.NewKeysService()

	broker := services.NewBroker()

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	repository := domain.NewBookRepository(database)

	booksService := services.NewBooksService(repository, broker)
	middleware := middlewares.NewMiddleware(keysService)
	booksController := controllers.NewBooksController(booksService)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	config.AddExposeHeaders("Authorization")
	router.Use(cors.New(config))

	apiRoutes := controllers.MapUrls(router, booksController, middleware)

	apiRoutes.Use(cors.New(config))

	router.Run(":3000")
}
