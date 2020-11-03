package services

import (
	"encoding/json"
	"librarymanager/books/common"
	"librarymanager/books/domain"
)

//Book use case interface
type Book interface {
	AddOne(book domain.BookDTO, createdByID string) (*domain.Book, common.CustomError)
	GetByID(id string) (*domain.Book, common.CustomError)
	All() (*[]domain.Book, common.CustomError)
	Destroy(id string, createdByID string) common.CustomError
	Update(id string, book domain.BookDTO, updatedByID string) (*domain.Book, common.CustomError)
}

type serviceStruct struct {
	repository domain.BookRepository
	rdb        common.Broker
}

//NewBooksService create a new book use case
func NewBooksService(repository domain.BookRepository, broker common.Broker) Book {

	return &serviceStruct{
		repository: repository,
		rdb:        broker,
	}
}

//AddOne method
func (u *serviceStruct) AddOne(book domain.BookDTO, createdByID string) (*domain.Book, common.CustomError) {

	if len(book.Title) == 0 {
		return nil, common.NewBadRequestError("invalid title")
	}

	return u.repository.Save(book.Title, book.Year, createdByID)
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
		payload := common.BrokerPayloadDTO{
			ID: id,
		}
		b, _ := json.Marshal(payload)
		go u.rdb.Publish("book_destroy", b)
	}

	return err
}

//Update method
func (u *serviceStruct) Update(id string, book domain.BookDTO, updatedByID string) (*domain.Book, common.CustomError) {

	savedBook, err := u.repository.Get(id)

	if err != nil {
		return nil, err
	}

	if savedBook.CreatedByID != updatedByID {
		err := common.NewUnauthorizedError("User cannot update this book")
		return nil, err
	}

	return u.repository.Update(id, book.Title, book.Year)
}
