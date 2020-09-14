package domain

import (
	"database/sql"
	"errors"
)

//Repository book repository (persistence)
type Repository interface {
	Save(title string, year int) (Book, error)
	Get(id string) (Book, error)
	All() ([]Book, error)
	Destroy(id string) error
}

type repositoryStruct struct {
	service Service
	db      *sql.DB
}

//NewBookRepository create a new book repository
func NewBookRepository(service Service, database *sql.DB) Repository {

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS book (id STRING PRIMARY KEY, title TEXT, year INTEGER)")
	statement.Exec()

	return &repositoryStruct{
		service: service,
		db:      database,
	}
}

//Save book
func (r *repositoryStruct) Save(title string, year int) (Book, error) {

	book := NewBook(r.service.GenerateID(), title, year)

	statement, _ := r.db.Prepare("INSERT INTO book (id, title, year) VALUES (?, ?, ?)")

	_, err := statement.Exec(book.ID, book.Title, book.Year)

	return *book, err
}

//Get get a book by its id
func (r *repositoryStruct) Get(id string) (Book, error) {

	book := &Book{}

	rows, err := r.db.Query("SELECT 1 id, title, year FROM book WHERE id = '" + id + "' LIMIT 1")

	if err != nil {
		return *book, err
	}

	for rows.Next() {
		var id string
		var title string
		var year int
		rows.Scan(&id, &title, &year)
		book = NewBook(id, title, year)

	}

	if book.ID == "" {
		return *book, errors.New("No book found for the given ID")
	}

	return *book, nil

}

//All list all books
func (r *repositoryStruct) All() ([]Book, error) {

	books := []Book{}

	rows, _ := r.db.Query("SELECT id, title, year FROM book")

	for rows.Next() {
		var id string
		var title string
		var year int
		rows.Scan(&id, &title, &year)
		book := NewBook(id, title, year)
		books = append(books, *book)
	}

	return books, nil

}

//Destroy destroy a book by its id
func (r *repositoryStruct) Destroy(id string) error {

	statement, _ := r.db.Prepare("DELETE FROM book WHERE id = ?")

	_, err := statement.Exec(id)
	return err

}
