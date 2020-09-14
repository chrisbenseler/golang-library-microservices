package main

import (
	"database/sql"
	"fmt"
	"librarymanager/books/domain"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type bookPayload struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func main() {

	fmt.Print("Books process")

	service := domain.NewBookService()

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	repository := domain.NewBookRepository(service, database)

	usecase := domain.NewBookUsecase(repository)

	r := gin.Default()
	r.GET("/api/books/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/books/:id", func(c *gin.Context) {
		book, _ := usecase.GetByID(c.Param("id"))
		c.JSON(200, book)
	})

	r.GET("/api/books", func(c *gin.Context) {
		books, _ := usecase.All()
		c.JSON(200, books)
	})

	r.POST("/api/books", func(c *gin.Context) {

		bookPayload := bookPayload{}
		if err := c.BindJSON(&bookPayload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book, err := usecase.AddOne(bookPayload.Title, bookPayload.Year)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, book)

	})

	r.Run(":3000")
}
