package domain

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Service books service interface
type Service interface {
	GenerateID() string
	VerifyToken(tokenString string, secretKey string) (*jwt.Token, error)
}

type usecaseService struct {
	//repository Repository
}

//NewBookService create a new book service
func NewBookService() Service {

	return &usecaseService{
		//repository: repository,
	}
}

//GenerateID method
func (u *usecaseService) GenerateID() string {

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)

}

//VerifyToken check is token is valid
func (u *usecaseService) VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
