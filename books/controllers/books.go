package controllers

import (
	"librarymanager/books/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Books controller interface
type Books interface {
	Create(c *gin.Context)
	All(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
}

type controllerStruct struct {
	booksService services.Book
}

//NewBooksController create new controller
func NewBooksController(booksService services.Book) Books {
	return &controllerStruct{
		booksService: booksService,
	}
}

type bookPayload struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func (r *controllerStruct) Create(c *gin.Context) {

	bookPayload := bookPayload{}
	if err := c.BindJSON(&bookPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	book, err := r.booksService.AddOne(bookPayload.Title, bookPayload.Year, userID.(string))
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(201, book)

}

func (r *controllerStruct) All(c *gin.Context) {
	books, _ := r.booksService.All()
	c.JSON(200, books)
}

func (r *controllerStruct) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")

	err := r.booksService.Destroy(c.Param("id"), userID.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func (r *controllerStruct) GetByID(c *gin.Context) {

	id := c.Param("id")

	if id == "ping" {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		return
	}

	book, err := r.booksService.GetByID(id)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, book)
}
