package controllers

import (
	"librarymanager/users/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Users controller interface
type Users interface {
	Get(*gin.Context)
}

type controllerStruct struct {
	usersService services.Users
}

//NewUsersController create new controller
func NewUsersController(service services.Users) Users {
	return &controllerStruct{
		usersService: service,
	}
}

func (u *controllerStruct) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := u.usersService.GetByID(id)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, user)
}
