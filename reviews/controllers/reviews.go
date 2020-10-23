package controllers

import (
	"librarymanager/reviews/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Reviews controller interface
type Reviews interface {
	GetFromBook(c *gin.Context)
	CreateInBook(c *gin.Context)
}

type controllerStruct struct {
	reviewsService services.Review
}

//NewReviewsController create new controller
func NewReviewsController(reviewsService services.Review) Reviews {
	return &controllerStruct{
		reviewsService: reviewsService,
	}
}

func (r *controllerStruct) GetFromBook(c *gin.Context) {
	bookID := c.Param("id")
	reviews, err := r.reviewsService.AllFromBook(bookID)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}
	c.JSON(200, reviews)
}

func (r *controllerStruct) CreateInBook(c *gin.Context) {
	bookID := c.Param("id")

	payload := reviewPayload{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	book, err := r.reviewsService.AddBookReview(bookID, payload.Content, userID.(string))
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(201, book)
}

type reviewPayload struct {
	Content string `json:"content"`
}
