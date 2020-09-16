package main

import (
	"database/sql"
	"fmt"
	"librarymanager/reviews/domain"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type reviewPayload struct {
	Content string `json:"content"`
}

func main() {

	fmt.Print("Reviews process")

	broker := domain.NewBroker()

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	repository := domain.NewReviewRepository(database)

	usecase := domain.NewReviewUsecase(repository, broker)

	usecase.Subscriptions()

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/api/reviewsping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRoutes := router.Group("/api/reviews")
	{
		apiRoutes.GET("/books/:id", func(c *gin.Context) {
			bookID := c.Param("id")
			reviews, _ := usecase.AllFromBook(bookID)
			c.JSON(200, reviews)
		})

		apiRoutes.POST("/books/:id", func(c *gin.Context) {

			bookID := c.Param("id")

			payload := reviewPayload{}
			if err := c.BindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			book, err := usecase.AddBookReview(bookID, payload.Content)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(201, book)

		})

	}

	apiRoutes.Use(cors.Default())

	router.Run(":3000")
}
