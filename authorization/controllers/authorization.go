package controllers

import (
	"librarymanager/authorization/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Authorization controller interface
type Authorization interface {
	SignIn(*gin.Context)
	SignUp(*gin.Context)
}

type controllerStruct struct {
	service services.Authorization
}

//NewAuthorizationController create new controller
func NewAuthorizationController(service services.Authorization) Authorization {
	return &controllerStruct{
		service: service,
	}
}

//SignIn sign in user
func (r *controllerStruct) SignIn(c *gin.Context) {
	authorizationPayload := services.AuthorizationDTO{}
	if err := c.BindJSON(&authorizationPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := r.service.Authenticate(authorizationPayload)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}

//SignUp sign up user
func (r *controllerStruct) SignUp(c *gin.Context) {
	userPayload := services.UserDTO{}
	if err := c.BindJSON(&userPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := r.service.CreateUser(userPayload)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, user)
}
