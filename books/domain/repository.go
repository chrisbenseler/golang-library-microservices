package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//Repository book repository (persistence)
type Repository interface {
	Save(title string, year int) error
}

type repositoryStruct struct {
	service Service
}

//NewBookRepository create a new book repository
func NewBookRepository(service Service) Repository {

	return &repositoryStruct{
		service: service,
	}
}

//Save book
func (r *repositoryStruct) Save(title string, year int) error {

	book := NewBook(r.service.GenerateID(), title, year)

	fmt.Print("Will persist new book", book.Title)

	bookJSON, _ := json.Marshal(book)

	fileName := book.ID + ".json"

	ioutil.WriteFile("./data/"+fileName, bookJSON, 0644)

	return errors.New("New error")
}
