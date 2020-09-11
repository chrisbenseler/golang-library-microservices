package tests

import (
	"errors"
	"librarymanager/books/domain"
	"testing"
)

type mockRepository struct {
}

func (m *mockRepository) Save(title string, year int) (domain.Book, error) {

	if title == "repeateditem" {
		return domain.Book{}, errors.New("Book exists")
	}

	book := domain.NewBook("randomid", title, year)

	return *book, nil
}

func (m *mockRepository) Get(id string) (domain.Book, error) {
	return domain.Book{}, nil
}

func (m *mockRepository) All() ([]domain.Book, error) {
	return []domain.Book{}, nil
}

func Test_usecase(t *testing.T) {

	repository := mockRepository{}

	usecase := domain.NewBookUsecase(&repository)

	_, err := usecase.AddOne("repeateditem", 2020)

	if err == nil {
		t.Error("Did not validate repeated item")
	}

	_, err1 := usecase.AddOne("newbook", 2020)

	if err1 != nil {
		t.Error("Couldn't insert new book")
	}

}

func Test_usecaseBlank(t *testing.T) {

	repository := mockRepository{}

	usecase := domain.NewBookUsecase(&repository)

	_, err := usecase.AddOne("", 2020)

	if err == nil {
		t.Error("Did not validate blank title")
	}

}
