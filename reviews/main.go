package main

import (
	"database/sql"
	"fmt"
	"librarymanager/reviews/common"
	"librarymanager/reviews/domain"
	"librarymanager/reviews/services"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type reviewPayload struct {
	Content string `json:"content"`
}

func main() {

	fmt.Print("Reviews process")

	broker := common.NewBroker()

	database, _ := sql.Open("sqlite3", "./data/tmp.db")

	repository := domain.NewReviewRepository(database)

	//usecase := domain.NewReviewUsecase(repository, broker)

	reviewsService := services.NewReviewsService(repository, broker)

	//usecase.Subscriptions()
	reviewsService.Subscriptions()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	config.AddExposeHeaders("Authorization")

	router.Use(cors.New(config))

	router.GET("/api/reviewsping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	keysService := services.NewKeysService()

	checkJWTToken := func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) == 2 {

			token, err := keysService.VerifyToken(strArr[1], os.Getenv("ACCESS_SECRET"))
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

	apiRoutes := router.Group("/api/reviews")
	{
		apiRoutes.GET("/books/:id", func(c *gin.Context) {
			bookID := c.Param("id")
			reviews, _ := reviewsService.AllFromBook(bookID)
			c.JSON(200, reviews)
		})

		apiRoutes.POST("/books/:id", checkJWTToken, func(c *gin.Context) {

			bookID := c.Param("id")

			payload := reviewPayload{}
			if err := c.BindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			userID, _ := c.Get("user_id")

			book, err := reviewsService.AddBookReview(bookID, payload.Content, userID.(string))
			if err != nil {
				c.JSON(err.Status(), gin.H{"error": err.Message()})
				return
			}

			c.JSON(201, book)

		})

	}

	apiRoutes.Use(cors.New(config))

	router.Run(":3000")
}
