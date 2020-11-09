package domain

import (
	"database/sql"
	"errors"
	"librarymanager/authorization/common"
)

//UserRepository user repository (persistence)
type UserRepository interface {
	Save(*User) (*User, common.CustomError)
	Get(int) (*User, common.CustomError)
	GetByEmail(string) (*User, common.CustomError)
}

type userRepositoryStruct struct {
	db *sql.DB
}

//NewUserRepository create a new user repository
func NewUserRepository(database *sql.DB) UserRepository {

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, email TEXT, password TEXT)")
	statement.Exec()

	return &userRepositoryStruct{
		db: database,
	}
}

//Save user
func (r *userRepositoryStruct) Save(user *User) (*User, common.CustomError) {

	statement, _ := r.db.Prepare("INSERT INTO user (email, password) VALUES (?, ?)")

	result, err := statement.Exec(user.Email, user.Password)

	if err != nil {
		return nil, common.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	idInt64, _ := result.LastInsertId()

	user.ID = int(idInt64)

	return user, nil

}

//Get user
func (r *userRepositoryStruct) Get(id int) (*User, common.CustomError) {

	user := NewUser("", "")

	row := r.db.QueryRow(`SELECT id, email FROM user WHERE id=$1`, id)

	err := row.Scan(&user.Email, id)

	if err != nil {
		return nil, common.NewNotFoundError("No user found for the given email")
	}

	return user, nil

}

//GetByEmail get a user by its email
func (r *userRepositoryStruct) GetByEmail(email string) (*User, common.CustomError) {

	user := NewUser("", "")

	row := r.db.QueryRow(`SELECT id, email FROM user WHERE email=$1`, email)

	err := row.Scan(&user.ID, &user.Email)

	if err != nil {
		return nil, common.NewNotFoundError("No user found for the given email")
	}

	return user, nil

}
