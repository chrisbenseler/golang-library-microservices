package services

import (
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

type MockBroker struct {
	//Publish(channel string, message interface{}) *redis.IntCmd
	//Subscribe(channel string, cb func(string))
	//Set(key string, message interface{}, time time.Duration) error
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
