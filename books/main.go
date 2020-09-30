package main

import (
	"database/sql"
	"fmt"
	"librarymanager/books/domain"
	"net/http"

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

	createNewBookRoute := func(c *gin.Context) {

		bookPayload := bookPayload{}
		if err := c.BindJSON(&bookPayload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, _ := c.Get("user_id")

		book, err := usecase.AddOne(bookPayload.Title, bookPayload.Year, userID.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, book)

	}

	apiRoutes := router.Group("/api/books")
	{

		apiRoutes.GET("/:id", func(c *gin.Context) {
			book, err := usecase.GetByID(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, book)
		})

		apiRoutes.GET("/", func(c *gin.Context) {
			books, _ := usecase.All()
			c.JSON(200, books)
		})

		apiRoutes.POST("/", middleware.CheckJWTToken, createNewBookRoute)

		apiRoutes.DELETE("/:id", func(c *gin.Context) {
			usecase.Destroy(c.Param("id"))
			c.JSON(200, gin.H{})
		})

	}

	apiRoutes.Use(cors.New(config))

	router.Run(":3000")
}
