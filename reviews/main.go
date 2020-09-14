package main

import (
	"database/sql"
	"fmt"
	"librarymanager/reviews/domain"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type reviewPayload struct {
	Content string `json:"content"`
}

func main() {

	fmt.Print("Reviews process")

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	repository := domain.NewReviewRepository(database)

	usecase := domain.NewReviewUsecase(repository)

	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*
		r.GET("/api/books/:id", func(c *gin.Context) {
			book, _ := usecase.GetByID(c.Param("id"))
			c.JSON(200, book)
		})

		r.GET("/api/books", func(c *gin.Context) {
			books, _ := usecase.All()
			c.JSON(200, books)
		})
	*/

	r.GET("/api/reviews/books/:id", func(c *gin.Context) {
		bookID := c.Param("id")
		books, _ := usecase.AllFromBook(bookID)
		c.JSON(200, books)
	})

	r.POST("/api/reviews/books/:id", func(c *gin.Context) {

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

	r.Run(":3000")
}
