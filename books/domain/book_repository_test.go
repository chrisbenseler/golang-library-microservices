package domain

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_BookRepository_Create(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	bookRepository := NewBookRepository(db)
	bookRepository.Initialize()

	rows := sqlmock.NewRows([]string{"id", "title", "year", "createdByID"}).
		AddRow("random_id", "title", "2020", "someid")
	mock.ExpectQuery("SELECT (.+) FROM book (.+) LIMIT 1").WithArgs("random_id").WillReturnRows(rows)

	_, getError := bookRepository.Get("random_id")

	if getError != nil {
		t.Fatalf("an error '%s' was not expected when querying", getError)
	}

	_, getError = bookRepository.Get("any_other_random_id")

	if getError == nil {
		t.Fatalf("should raise an error")
	}

}

/*

type MockBroker struct {
}

func (m *MockBroker) Subscribe(channel string, cb func(string)) {

}

func (m *MockBroker) Set(key string, message interface{}, time time.Duration) error {

	return nil
}

func (m *MockBroker) Publish(channel string, message interface{}) *redis.IntCmd {

	intCmd := &redis.IntCmd{}
	return intCmd
}

type MockUserRepository struct {
}

func (r *MockUserRepository) Save(user *domain.User) (*domain.User, common.CustomError) {
	return nil, nil
}
*/
