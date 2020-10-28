package domain

//AuthorizationDTO auth payload dto
type AuthorizationDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
