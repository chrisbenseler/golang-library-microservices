package domain

import (
	"encoding/json"
	"io/ioutil"
)

//Repository book repository (persistence)
type Repository interface {
	Save(title string, year int) (Book, error)
	Get(id string) (Book, error)
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
func (r *repositoryStruct) Save(title string, year int) (Book, error) {

	book := NewBook(r.service.GenerateID(), title, year)

	bookJSON, _ := json.Marshal(book)

	fileName := book.ID + ".json"

	err := ioutil.WriteFile("./data/"+fileName, bookJSON, 0644)

	return *book, err

}

//Get book
func (r *repositoryStruct) Get(id string) (Book, error) {

	book := Book{}

	fileName := id + ".json"

	byteValue, err := ioutil.ReadFile("./data/" + fileName)

	if err != nil {
		return book, err
	}

	json.Unmarshal(byteValue, &book)

	return book, nil

}
