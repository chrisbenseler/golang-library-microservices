package domain

//User user model
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

//NewUser new user entity
func NewUser(email string, fullName string) *User {
	return &User{
		Email:    email,
		FullName: fullName,
	}
}
