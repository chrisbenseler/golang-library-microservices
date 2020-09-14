package domain

import (
	"errors"
	"fmt"
)

//Usecase books use case interface
type Usecase interface {
	AddOne(title string, year int) (Book, error)
	GetByID(id string) (Book, error)
	All() ([]Book, error)
	Destroy(id string) error
}

type usecaseStruct struct {
	repository Repository
	rdb        Broker
}

//NewBookUsecase create a new book use case
func NewBookUsecase(repository Repository, broker Broker) Usecase {

	broker.Subscribe("book_destroy", func(data interface{}) {
		fmt.Print("Book destroyed msg from broker")
		fmt.Print(data)
	})

	broker.Subscribe("book_getall", func(data interface{}) {
		fmt.Print("Book get all msg from broker")
		fmt.Print(data)
	})

	return &usecaseStruct{
		repository: repository,
		rdb:        broker,
	}
}

//AddOne method
func (u *usecaseStruct) AddOne(title string, year int) (Book, error) {

	if len(title) == 0 {
		return Book{}, errors.New("Invalid parameters provided")
	}

	return u.repository.Save(title, year)
}

//GetByID method
func (u *usecaseStruct) GetByID(id string) (Book, error) {
	return u.repository.Get(id)
}

//All method
func (u *usecaseStruct) All() ([]Book, error) {

	go u.rdb.Publish("book_getall", "all")

	return u.repository.All()
}

//Destroy destroy a book
func (u *usecaseStruct) Destroy(id string) error {

	err := u.repository.Destroy(id)

	if err == nil {
		go u.rdb.Publish("book_destroy", id)
	}

	return err
}
