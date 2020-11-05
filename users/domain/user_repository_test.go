package domain

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_UserRepository_Get(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository := NewUserRepository(db)
	userRepository.Initialize()

	rows := sqlmock.NewRows([]string{"id", "email", "fullName"}).
		AddRow("random_id", "random@email", "randmon fullname")
	mock.ExpectQuery("SELECT (.+) FROM user").WithArgs("random_id").WillReturnRows(rows)

	user, getError := userRepository.Get("random_id")

	if getError != nil {
		t.Fatalf("an error '%s' was not expected when querying", getError)
		return
	}

	if user.Email != "random@email" {
		t.Fatalf("did not return user with valid email")
		return
	}

	_, getError = userRepository.Get("any_other_random_id")

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
	userRepository := NewUserRepository(db)
	userRepository.Initialize()

	rows := sqlmock.NewRows([]string{"id", "email", "fullName"}).
		AddRow("random_id", "random@email", "randmon fullname")
	mock.ExpectQuery("SELECT (.+) FROM user").WithArgs("random@email").WillReturnRows(rows)

	user, getError := userRepository.GetByEmail("random@email")

	if getError != nil {
		t.Fatalf("an error '%s' was not expected when querying", getError)
		return
	}

	if user.ID != "random_id" {
		t.Fatalf("did not return user with expected id")
		return
	}

	_, getError = userRepository.Get("any_other_random_id")

	if getError == nil {
		t.Fatalf("should raise an error")
	}

}
