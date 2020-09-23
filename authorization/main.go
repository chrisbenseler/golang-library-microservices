package main

import (
	"fmt"
	"librarymanager/authorization/domain"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

type authorizationPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//TokenDetails token metadata struct
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

func main() {

	fmt.Print("Authroization process")

	router := gin.Default()

	router.Use(cors.Default())

	broker := domain.NewBroker()

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

			userID := "root"

			ts, err := CreateToken(userID)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			saveErr := CreateAuth(userID, ts, broker)
			if saveErr != nil {
				c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
			}

			tokens := map[string]string{
				"access_token":  ts.AccessToken,
				"refresh_token": ts.RefreshToken,
			}

			c.JSON(http.StatusOK, gin.H{"tokens": tokens})

		})

	}

	apiRoutes.Use(cors.Default())

	router.Run(":3000")
}

//CreateToken Create json webtoken
func CreateToken(userKey string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userKey
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userKey
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

//CreateAuth create auth iin broker
func CreateAuth(userID string, td *TokenDetails, broker domain.Broker) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	//broker.

	errAccess := broker.Set(td.AccessUUID, userID, at.Sub(now))
	if errAccess != nil {
		return errAccess
	}
	errRefresh := broker.Set(td.RefreshUUID, userID, rt.Sub(now))
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
