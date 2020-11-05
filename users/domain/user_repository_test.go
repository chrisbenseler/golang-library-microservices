package domain

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_UserRepository_Create(t *testing.T) {

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

	fmt.Println("rows")
	fmt.Println(rows)

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
