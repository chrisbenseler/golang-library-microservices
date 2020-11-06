package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"librarymanager/users/common"
)

//UserRepository user repository (persistence)
type UserRepository interface {
	Get(int) (*User, common.CustomError)
	GetByEmail(string) (*User, common.CustomError)
	Save(email string, fullName string) (*User, common.CustomError)
	Initialize() common.CustomError
}

type repositoryStruct struct {
	db *sql.DB
}

//NewUsersRepository create a new user repository
func NewUsersRepository(database *sql.DB) UserRepository {

	return &repositoryStruct{
		db: database,
	}
}

func (r *repositoryStruct) Initialize() common.CustomError {
	statement, err := r.db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, fullName TEXT, email TEXT)")
	if err != nil {
		return common.NewInternalServerError("error when trying to create database table", errors.New("database error"))
	}
	statement.Exec()
	return nil
}

//Get get a user by its id
func (r *repositoryStruct) Get(id int) (*User, common.CustomError) {

	user := NewUser("", "")

	row := r.db.QueryRow(`SELECT email, fullName FROM user WHERE id=$1`, id)

	err := row.Scan(&user.Email, &user.FullName)

	if err != nil {
		return nil, common.NewNotFoundError("No user found for the given ID")
	}

	return user, nil

}

//Get get a user by its email
func (r *repositoryStruct) GetByEmail(email string) (*User, common.CustomError) {

	user := NewUser("", "")

	row := r.db.QueryRow(`SELECT email, fullName, id FROM user WHERE email=$1`, email)

	err := row.Scan(&user.Email, &user.FullName, &user.ID)

	if err != nil {
		fmt.Println(err)
		return nil, common.NewNotFoundError("No user found for the given email")
	}

	return user, nil

}

//Save user
func (r *repositoryStruct) Save(email string, fullName string) (*User, common.CustomError) {

	statement, dbErr := r.db.Prepare("INSERT INTO user (email, fullName) VALUES (?, ?)")

	if dbErr != nil {

		fmt.Println(dbErr)
		return nil, common.NewInternalServerError("error when trying to prepare statement", dbErr)
	}

	_, err := statement.Exec(email, fullName)

	if err != nil {
		return nil, common.NewInternalServerError("error when tying to save user", err)
	}

	return r.GetByEmail(email)

}
