package domain

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//Repository book repository (persistence)
type Repository interface {
	Save(title string, year int) (Book, error)
	Get(id string) (Book, error)
	All() ([]Book, error)
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

func readFile(path string) (Book, error) {
	book := Book{}

	byteValue, err := ioutil.ReadFile(path)

	if err != nil {
		return book, err
	}

	json.Unmarshal(byteValue, &book)

	return book, nil
}

//Get get a book by its id
func (r *repositoryStruct) Get(id string) (Book, error) {

	fileName := id + ".json"

	return readFile("./data/" + fileName)

}

//All list all books
func (r *repositoryStruct) All() ([]Book, error) {

	books := []Book{}

	var files []string

	err := filepath.Walk("./data/", func(path string, info os.FileInfo, err error) error {

		if strings.Contains(path, ".json") == true {
			files = append(files, path)
		}
		return nil

	})

	if err != nil {
		return []Book{}, err
	}

	for _, file := range files {
		book, _ := readFile(file)
		books = append(books, book)
	}

	return books, nil

}
