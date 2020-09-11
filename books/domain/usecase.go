package domain

import "errors"

//Usecase books use case interface
type Usecase interface {
	AddOne(title string, year int) error
}

type usecaseStruct struct {
	repository Repository
}

//NewBookUsecase create a new book use case
func NewBookUsecase(repository Repository) Usecase {

	return &usecaseStruct{
		repository: repository,
	}
}

//AddOne method
func (u *usecaseStruct) AddOne(title string, year int) error {

	u.repository.Save(title, year)
	return errors.New("Could not create new book")
}
