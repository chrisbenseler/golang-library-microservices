package domain

import (
	"database/sql"
	"errors"
	"librarymanager/authorization/common"
)

//UserRepository user repository (persistence)
type UserRepository interface {
	Save(email string, password string) (*User, common.CustomError)
}

type userRepositoryStruct struct {
	db *sql.DB
}

//NewUserRepository create a new user repository
func NewUserRepository(database *sql.DB) UserRepository {

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (id STRING PRIMARY KEY, email TEXT, password TEXT)")
	statement.Exec()

	return &userRepositoryStruct{
		db: database,
	}
}

//Save book
func (r *userRepositoryStruct) Save(email string, password string) (*User, common.CustomError) {

	user := NewUser(email, password)

	statement, _ := r.db.Prepare("INSERT INTO user (email, password) VALUES (?, ?)")

	if _, err := statement.Exec(user.Email, user.Password); err != nil {
		return nil, common.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	return user, nil

}
