package domain

import (
	"math/rand"
	"time"
)

//Service books service interface
type Service interface {
	GenerateID() string
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

	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)

}
