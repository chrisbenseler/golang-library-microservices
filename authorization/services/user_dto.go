package services

//UserDTO auth payload dto
type UserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
