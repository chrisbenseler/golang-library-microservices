package controllers

import (
	"librarymanager/authorization/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Authorization controller interface
type Authorization interface {
	SignIn(c *gin.Context)
}

type controllerStruct struct {
	service domain.Usecase
}

//NewAuthorizationController create new controller
func NewAuthorizationController(service domain.Usecase) Authorization {
	return &controllerStruct{
		service: service,
	}
}

//SignIn sign in user
func (r *controllerStruct) SignIn(c *gin.Context) {
	authorizationPayload := authorizationPayload{}
	if err := c.BindJSON(&authorizationPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := r.service.Authenticate(authorizationPayload.Email, authorizationPayload.Password)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}

type authorizationPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
