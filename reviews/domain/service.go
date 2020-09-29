package domain

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

//Service books service interface
type Service interface {
	VerifyToken(tokenString string, secretKey string) (*jwt.Token, error)
}

type usecaseService struct {
	//repository Repository
}

//NewReviewService create a new book service
func NewReviewService() Service {

	return &usecaseService{
		//repository: repository,
	}
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
