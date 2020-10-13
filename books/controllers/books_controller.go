package controllers

import (
	"librarymanager/books/usecases"
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
	usecase usecases.Book
}

//NewBooksController create new controller
func NewBooksController(usecase usecases.Book) Books {
	return &controllerStruct{
		usecase: usecase,
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

	book, err := r.usecase.AddOne(bookPayload.Title, bookPayload.Year, userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, book)

}

func (r *controllerStruct) All(c *gin.Context) {
	books, _ := r.usecase.All()
	c.JSON(200, books)
}

func (r *controllerStruct) Delete(c *gin.Context) {
	userID, _ := c.Get("user_id")

	err := r.usecase.Destroy(c.Param("id"), userID.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func (r *controllerStruct) GetByID(c *gin.Context) {
	book, err := r.usecase.GetByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, book)
}
