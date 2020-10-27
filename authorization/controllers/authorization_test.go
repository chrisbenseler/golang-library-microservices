package controllers

import (
	"librarymanager/authorization/common"
	"librarymanager/authorization/controllers"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_Authorization_SignIn(t *testing.T) {

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	authorizationService := new(MockService)
	controller := controllers.NewAuthorizationController(authorizationService)

	controller.SignIn(c)
}

type MockService struct {
}

func (u *MockService) Authenticate(email string, password string) (map[string]string, common.CustomError) {

	value := map[string]string{}
	return value, nil
}
