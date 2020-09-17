package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type authorizationPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	fmt.Print("Authroization process")

	router := gin.Default()

	router.Use(cors.Default())

	apiRoutes := router.Group("/api/authorization")
	{

		apiRoutes.POST("/signin", func(c *gin.Context) {

			authorizationPayload := authorizationPayload{}
			if err := c.BindJSON(&authorizationPayload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if authorizationPayload.Email != "root@gmail.com" || authorizationPayload.Password != "root" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
				return
			}

			token, err := CreateToken("root")
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})

		})

	}

	apiRoutes.Use(cors.Default())

	router.Run(":3000")
}

//CreateToken Create json webtoken
func CreateToken(userKey string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userKey
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
