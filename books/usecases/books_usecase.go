package usecases

import (
	"encoding/json"
	"errors"
	"librarymanager/books/domain"
)

//Book use case interface
type Book interface {
	AddOne(title string, year int, createdByID string) (domain.Book, error)
	GetByID(id string) (domain.Book, error)
	All() ([]domain.Book, error)
	Destroy(id string, createdByID string) error
}

type usecaseStruct struct {
	repository domain.Repository
	rdb        domain.Broker
}

//NewBooksUsecase create a new book use case
func NewBooksUsecase(repository domain.Repository, broker domain.Broker) Book {

	return &usecaseStruct{
		repository: repository,
		rdb:        broker,
	}
}

//AddOne method
func (u *usecaseStruct) AddOne(title string, year int, createdByID string) (domain.Book, error) {

	if len(title) == 0 {
		return domain.Book{}, errors.New("Invalid parameters provided")
	}

	return u.repository.Save(title, year, createdByID)
}

//GetByID method
func (u *usecaseStruct) GetByID(id string) (domain.Book, error) {
	return u.repository.Get(id)
}

//All method
func (u *usecaseStruct) All() ([]domain.Book, error) {
	return u.repository.All()
}

//Destroy destroy a book
func (u *usecaseStruct) Destroy(id string, createdByID string) error {

	book, _ := u.repository.Get(id)

	if book.CreatedByID != createdByID {
		return errors.New("This user cannot perform this action")
	}

	err := u.repository.Destroy(id)

	if err == nil {
		payload := brokerPayload{
			ID: id,
		}
		b, _ := json.Marshal(payload)
		u.rdb.Publish("book_destroy", b)
	}

	return err
}

type brokerPayload struct {
	ID    string
	Extra string
}
