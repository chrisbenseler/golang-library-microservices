package middlewares

import (
	"librarymanager/reviews/services"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Middleware middlewares for http server
type Middleware interface {
	CheckJWTToken(c *gin.Context)
}

type middlewareStruct struct {
	service services.Keys
}

//NewMiddleware ccreate new middleware
func NewMiddleware(service services.Keys) Middleware {

	return &middlewareStruct{
		service: service,
	}
}

//CheckJWTToken check JWT token
func (m *middlewareStruct) CheckJWTToken(c *gin.Context) {
	bearToken := c.GetHeader("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {

		token, err := m.service.VerifyToken(strArr[1], os.Getenv("ACCESS_SECRET"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["user_id"])

		return

	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
	return
}
