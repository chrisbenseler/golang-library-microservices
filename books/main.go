package main

import (
	"fmt"
	"librarymanager/books/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookPayload struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func main() {

	fmt.Print("Books process")

	service := domain.NewBookService()

	repository := domain.NewBookRepository(service)

	usecase := domain.NewBookUsecase(repository)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/books/:id", func(c *gin.Context) {
		book, _ := usecase.GetByID(c.Param("id"))
		c.JSON(200, book)
	})

	r.GET("/books", func(c *gin.Context) {
		books, _ := usecase.All()
		c.JSON(200, books)
	})

	r.POST("/books", func(c *gin.Context) {

		bookPayload := BookPayload{}
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

	r.Run()
}
