package services

import (
	"encoding/json"
	"fmt"
	"librarymanager/users/common"
	"librarymanager/users/domain"
)

//Users users service interface
type Users interface {
	Subscriptions()
	Create(string, string) (*domain.User, common.CustomError)
}

type serviceStruct struct {
	userRepository domain.UserRepository
	broker         common.Broker
}

//NewUsersService new users service
func NewUsersService(userRepository domain.UserRepository, broker common.Broker) Users {

	return &serviceStruct{
		userRepository: userRepository,
		broker:         broker,
	}
}

//Subscriptions subscriptions
func (u *serviceStruct) Subscriptions() {
	fmt.Println("\nSubscriptions in users service")

	u.broker.Subscribe("authorization.signup", func(data string) {

		fmt.Println("Broker - authorization.signup")

		payload := common.BrokerPayloadDTO{}

		json.Unmarshal([]byte(data), &payload)

		_, err := u.userRepository.Save(payload.Extra, "")

		if err != nil {
			fmt.Println("Broker - error when trying to create new user", err)
			return
		}
		//fmt.Println("Broker - create new user", payload.Extra)
	})

}

func (u *serviceStruct) Create(email string, fullName string) (*domain.User, common.CustomError) {

	return u.userRepository.Save(email, fullName)
}
