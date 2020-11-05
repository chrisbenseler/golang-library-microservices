package domain

import (
	"database/sql"
	"errors"
	"librarymanager/books/common"
)

//UserRepository user repository (persistence)
type UserRepository interface {
	Get(string) (*User, common.CustomError)
	Initialize() common.CustomError
}

type repositoryStruct struct {
	db *sql.DB
}

//NewUserRepository create a new user repository
func NewUserRepository(database *sql.DB) UserRepository {

	return &repositoryStruct{
		db: database,
	}
}

func (r *repositoryStruct) Initialize() common.CustomError {
	statement, err := r.db.Prepare("CREATE TABLE IF NOT EXISTS user (id STRING PRIMARY KEY, fullName TEXT, email TEXT)")
	if err != nil {
		return common.NewInternalServerError("error when trying to create database table", errors.New("database error"))
	}
	statement.Exec()
	return nil
}

//Get get a user by its id
func (r *repositoryStruct) Get(id string) (*User, common.CustomError) {

	user := &User{}

	var email string
	var fullName string
	var receivedID string

	row := r.db.QueryRow(`SELECT email, fullName FROM user WHERE id=$1`, id)

	err := row.Scan(&receivedID, &email, &fullName)

	if err != nil {
		return nil, common.NewNotFoundError("No user found for the given ID")
	}

	user = NewUser(email, fullName)
	user.ID = receivedID

	return user, nil

}
