package domain

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_UserRepository_Get(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository := NewUsersRepository(db)
	userRepository.Initialize()

	rows := sqlmock.NewRows([]string{"id", "email", "fullName"}).
		AddRow(1, "random@email", "random fullname")
	mock.ExpectQuery("SELECT (.+) FROM user").WithArgs(1).WillReturnRows(rows)

	user, getError := userRepository.Get(1)

	if getError != nil {
		t.Fatalf("an error '%s' was not expected when querying", getError)
		return
	}

	if user.Email != "random@email" {
		t.Fatalf("did not return user with valid email")
		return
	}

	_, getError = userRepository.Get(1111)

	if getError == nil {
		t.Fatalf("should raise an error")
	}

}

func Test_UserRepository_GetByEmail(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository := NewUsersRepository(db)
	userRepository.Initialize()

	rows := sqlmock.NewRows([]string{"id", "email", "fullName"}).
		AddRow(1, "random@email", "randmon fullname")
	mock.ExpectQuery("SELECT (.+) FROM user").WithArgs("random@email").WillReturnRows(rows)

	user, getError := userRepository.GetByEmail("random@email")

	fmt.Println("user", user)

	if getError != nil {
		t.Fatalf("an error '%s' was not expected when querying", getError)
		return
	}

	if user.Email != "random@email" {
		t.Fatalf("did not return user with expected id")
		return
	}

}
