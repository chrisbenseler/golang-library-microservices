package domain

import (
	"database/sql"
	"errors"
)

//Repository book repository (persistence)
type Repository interface {
	Save(title string, year int, createdByID string) (Book, error)
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

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS book (id STRING PRIMARY KEY, title TEXT, year INTEGER, createdByID TEXT)")
	statement.Exec()

	return &repositoryStruct{
		service: service,
		db:      database,
	}
}

//Save book
func (r *repositoryStruct) Save(title string, year int, createdByID string) (Book, error) {

	book := NewBook(r.service.GenerateID(), title, year, createdByID)

	statement, _ := r.db.Prepare("INSERT INTO book (id, title, year, createdByID) VALUES (?, ?, ?, ?)")

	_, err := statement.Exec(book.ID, book.Title, book.Year, book.CreatedByID)

	return *book, err
}

//Get get a book by its id
func (r *repositoryStruct) Get(id string) (Book, error) {

	book := &Book{}

	rows, err := r.db.Query("SELECT 1 title, year, createdByID FROM book WHERE id = '" + id + "' LIMIT 1")

	if err != nil {
		return *book, err
	}

	for rows.Next() {

		var title string
		var year int
		var createdByID string
		rows.Scan(&title, &year, &createdByID)
		book = NewBook(id, title, year, createdByID)

	}

	if book.ID == "" {
		return *book, errors.New("No book found for the given ID")
	}

	return *book, nil

}

//All list all books
func (r *repositoryStruct) All() ([]Book, error) {

	books := []Book{}

	rows, _ := r.db.Query("SELECT id, title, year, createdByID FROM book")

	for rows.Next() {
		var id string
		var title string
		var year int
		var createdByID string
		rows.Scan(&id, &title, &year, &createdByID)
		book := NewBook(id, title, year, createdByID)
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
