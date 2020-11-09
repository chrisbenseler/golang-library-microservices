package domain

//User user model
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser new user entity
func NewUser(email string, password string) *User {

	return &User{
		Email:    email,
		Password: password,
	}
}
