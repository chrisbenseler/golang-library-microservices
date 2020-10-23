package domain

import (
	"encoding/json"
	"fmt"
	"librarymanager/reviews/common"
)

//Usecase reviews use case interface
type Usecase interface {
	AddBookReview(bookID string, content string, createdByID string) (*Review, common.CustomError)
	AllFromBook(bookID string) (*[]Review, common.CustomError)
	Subscriptions()
}

type usecaseStruct struct {
	repository Repository
	broker     common.Broker
}

//NewReviewUsecase a new book use case
func NewReviewUsecase(repository Repository, broker common.Broker) Usecase {

	return &usecaseStruct{
		repository: repository,
		broker:     broker,
	}
}

//Subscriptions subscriptions
func (u *usecaseStruct) Subscriptions() {
	fmt.Print("\nSubscriptions in review use case")

	u.broker.Subscribe("book_destroy", func(data string) {
		fmt.Println("\nBook destroyed msg from broker", data)

		payload := payloadBroker{}

		json.Unmarshal([]byte(data), &payload)

		u.repository.DestroyByType(payload.ID, "book")
	})
}

//AddBookReview method
func (u *usecaseStruct) AddBookReview(bookID string, content string, createdByID string) (*Review, common.CustomError) {
	return u.repository.Save(bookID, "book", content, createdByID)
}

//AllFromBook method
func (u *usecaseStruct) AllFromBook(bookID string) (*[]Review, common.CustomError) {
	return u.repository.FindAll(bookID, "book")
}

type payloadBroker struct {
	ID string
}
