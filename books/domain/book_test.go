package domain

import (
	"testing"
)

func Test_Domain_Book(t *testing.T) {
	book := NewBook("randomid", "randomtitle", 2020, "randomcreatedid")

	if book.ID != "randomid" {
		t.Error("Invalid id")
	}
}
