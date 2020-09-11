package tests

import (
	"librarymanager/books/domain"
	"testing"
)

func Test_Service(t *testing.T) {

	service := domain.NewBookService()

	id := service.GenerateID()

	if len(id) != 16 {
		t.Error("Invalid length for new ID")
	}
}
