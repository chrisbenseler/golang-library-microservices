package services

import (
	"librarymanager/authorization/common"
	"librarymanager/authorization/domain"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func Test_Authorization_CreateToken(t *testing.T) {

	td, err := CreateToken("somekey")

	if err != nil {
		t.Error("Error when creating token details")
	}

	if len(td.AccessToken) == 0 {
		t.Error("Invalid token length")
	}

	if len(td.RefreshToken) == 0 {
		t.Error("Must have refresh token")
	}

}

func Test_Authorization_CreateAuth(t *testing.T) {

	td := &TokenDetails{}

	broker := new(MockBroker)

	err := CreateAuth("somekey", td, broker)

	if err != nil {
		t.Error("Error when creating authentication")
	}
}

func Test_Authorization_Auth(t *testing.T) {

	broker := new(MockBroker)

	userRepository := new(MockUserRepository)

	service := NewAuthorizationService(userRepository, broker)

	authorizationPayload := AuthorizationDTO{
		Email:    "root@gmail.com",
		Password: "root",
	}
	td, err := service.Authenticate(authorizationPayload)

	r := *td //

	if err != nil {
		t.Error("Error when authenticate root")
	}

	token := r["access_token"]

	if len(token) == 0 {
		t.Error("Invalid access token")
	}

}

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

func (r *MockUserRepository) Get(id int) (*domain.User, common.CustomError) {
	return nil, nil
}

func (r *MockUserRepository) GetByCredentials(email string, password string) (*domain.User, common.CustomError) {
	return &domain.User{ID: 20, Email: "teste@gmail.com"}, nil
}

func (r *MockUserRepository) GetByEmail(email string) (*domain.User, common.CustomError) {
	return nil, nil
}
