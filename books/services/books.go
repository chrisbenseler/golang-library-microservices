package services

import (
	"encoding/json"
	"librarymanager/books/common"
	"librarymanager/books/domain"
)

//Book use case interface
type Book interface {
	AddOne(title string, year int, createdByID string) (*domain.Book, common.CustomError)
	GetByID(id string) (*domain.Book, common.CustomError)
	All() (*[]domain.Book, common.CustomError)
	Destroy(id string, createdByID string) common.CustomError
}

type serviceStruct struct {
	repository domain.Repository
	rdb        common.Broker
}

//NewBooksService create a new book use case
func NewBooksService(repository domain.Repository, broker common.Broker) Book {

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
func (u *serviceStruct) All() (*[]domain.Book, common.CustomError) {
	return u.repository.All()
}

//Destroy destroy a book
func (u *serviceStruct) Destroy(id string, createdByID string) common.CustomError {

	book, _ := u.repository.Get(id)

	if book.CreatedByID != createdByID {
		return common.NewUnauthorizedError("User cannot perform this action on this resource")
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
