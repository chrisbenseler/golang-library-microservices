package main

import (
	"fmt"
	"librarymanager/authorization/domain"
	"net/http"

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

	broker := domain.NewBroker()

	usecase := domain.NewUsecase(broker)

	apiRoutes := router.Group("/api/authorization")
	{

		apiRoutes.POST("/signin", func(c *gin.Context) {

			authorizationPayload := authorizationPayload{}
			if err := c.BindJSON(&authorizationPayload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			tokens, err := usecase.Authenticate(authorizationPayload.Email, authorizationPayload.Password)

			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"tokens": tokens})

		})

	}

	apiRoutes.Use(cors.Default())

	router.Run(":3000")
}
