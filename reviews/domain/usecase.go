package domain

import (
	"encoding/json"
	"fmt"
)

//Usecase reviews use case interface
type Usecase interface {
	AddBookReview(bookID string, content string) (Review, error)
	AllFromBook(bookID string) ([]Review, error)
	Subscriptions()
}

type usecaseStruct struct {
	repository Repository
	broker     Broker
}

//NewReviewUsecase a new book use case
func NewReviewUsecase(repository Repository, broker Broker) Usecase {

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
func (u *usecaseStruct) AddBookReview(bookID string, content string) (Review, error) {
	return u.repository.Save(bookID, "book", content)
}

//AllFromBook method
func (u *usecaseStruct) AllFromBook(bookID string) ([]Review, error) {
	return u.repository.FindAll(bookID, "book")
}

type payloadBroker struct {
	ID string
}
