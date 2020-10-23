package services

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

//Keys service interface
type Keys interface {
	VerifyToken(tokenString string, secretKey string) (*jwt.Token, error)
}

type keysService struct {
}

//NewKeysService create a new keys service
func NewKeysService() Keys {
	return &keysService{}
}

//VerifyToken check is token is valid
func (u *keysService) VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {

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
