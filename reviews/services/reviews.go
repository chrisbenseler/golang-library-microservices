package services

import (
	"encoding/json"
	"fmt"
	"librarymanager/reviews/common"
	"librarymanager/reviews/domain"
)

//Service reviews service interface
type Service interface {
	AddBookReview(bookID string, content string, createdByID string) (*domain.Review, common.CustomError)
	AllFromBook(bookID string) (*[]domain.Review, common.CustomError)
	Subscriptions()
}

type serviceStruct struct {
	repository domain.Repository
	broker     common.Broker
}

//NewReviewsService new reviews service
func NewReviewsService(repository domain.Repository, broker common.Broker) Service {

	return &serviceStruct{
		repository: repository,
		broker:     broker,
	}
}

//Subscriptions subscriptions
func (u *serviceStruct) Subscriptions() {
	fmt.Print("\nSubscriptions in review use case")

	u.broker.Subscribe("book_destroy", func(data string) {
		fmt.Println("\nBook destroyed msg from broker", data)

		payload := payloadBroker{}

		json.Unmarshal([]byte(data), &payload)

		u.repository.DestroyByType(payload.ID, "book")
	})
}

//AddBookReview method
func (u *serviceStruct) AddBookReview(bookID string, content string, createdByID string) (*domain.Review, common.CustomError) {
	return u.repository.Save(bookID, "book", content, createdByID)
}

//AllFromBook method
func (u *serviceStruct) AllFromBook(bookID string) (*[]domain.Review, common.CustomError) {
	return u.repository.FindAll(bookID, "book")
}

type payloadBroker struct {
	ID string
}
