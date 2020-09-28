package domain

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

//Usecase struct
type Usecase interface {
	Authenticate(email string, password string) (map[string]string, error)
}

type usecaseStruct struct {
	broker Broker
}

//NewUsecase create new use case
func NewUsecase(broker Broker) Usecase {

	return &usecaseStruct{
		broker: broker,
	}
}

func (u *usecaseStruct) Authenticate(email string, password string) (map[string]string, error) {

	if email != "root@gmail.com" || password != "root" {

		return map[string]string{}, errors.New("Credenciais inválidas")
	}

	userID := "root"

	ts, err := CreateToken(userID)
	if err != nil {
		return map[string]string{}, err
	}

	saveErr := CreateAuth(userID, ts, u.broker)
	if saveErr != nil {
		return map[string]string{}, saveErr
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	return tokens, nil

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
	os.Setenv("ACCESS_SECRET", os.Getenv("ACCESS_SECRET"))
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

//CreateAuth create auth in broker
func CreateAuth(userID string, td *TokenDetails, broker Broker) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

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

//TokenDetails token metadata struct
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}
