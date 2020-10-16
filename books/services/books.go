package services

import (
	"encoding/json"
	"errors"
	"librarymanager/books/common"
	"librarymanager/books/domain"
)

//Book use case interface
type Book interface {
	AddOne(title string, year int, createdByID string) (*domain.Book, common.CustomError)
	GetByID(id string) (*domain.Book, common.CustomError)
	All() (*[]domain.Book, error)
	Destroy(id string, createdByID string) error
}

type serviceStruct struct {
	repository domain.Repository
	rdb        Broker
}

//NewBooksService create a new book use case
func NewBooksService(repository domain.Repository, broker Broker) Book {

	return &serviceStruct{
		repository: repository,
		rdb:        broker,
	}
}

//AddOne method
func (u *serviceStruct) AddOne(title string, year int, createdByID string) (*domain.Book, common.CustomError) {

	if len(title) == 0 {
		return nil, common.NewBadRequestError("invalid title")
	}

	return u.repository.Save(title, year, createdByID)
}

//GetByID method
func (u *serviceStruct) GetByID(id string) (*domain.Book, common.CustomError) {
	return u.repository.Get(id)
}

//All method
func (u *serviceStruct) All() (*[]domain.Book, error) {
	return u.repository.All()
}

//Destroy destroy a book
func (u *serviceStruct) Destroy(id string, createdByID string) error {

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
