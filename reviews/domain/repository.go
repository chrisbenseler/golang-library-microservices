package domain

import (
	"database/sql"
	"errors"
	"librarymanager/reviews/common"
	"strconv"
)

//Repository book repository (persistence)
type Repository interface {
	Save(entityID string, entityKey string, content string, createdByID string) (*Review, common.CustomError)
	FindAll(entityID string, entityKey string) (*[]Review, common.CustomError)
	DestroyByType(entityID string, entityKey string) common.CustomError
}

type repositoryStruct struct {
	db *sql.DB
}

//NewReviewRepository create a new book repository
func NewReviewRepository(database *sql.DB) Repository {

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS review (id INTEGER PRIMARY KEY, content TEXT, entityID TEXT, entityKey TEXT, createdByID TEXT)")
	statement.Exec()

	return &repositoryStruct{
		db: database,
	}
}

//Save review
func (r *repositoryStruct) Save(entityID string, entityKey string, content string, createdByID string) (*Review, common.CustomError) {

	review := NewReview("", content, entityID, entityKey, createdByID)

	statement, _ := r.db.Prepare("INSERT INTO review (content, entityID, entityKey, createdByID) VALUES (?, ?, ?, ?)")

	result, err := statement.Exec(review.Content, review.EntityID, review.EntityKey, review.CreatedByID)

	if err != nil {
		return nil, common.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	id, _ := result.LastInsertId()

	review.ID = strconv.FormatInt(int64(id), 10)

	return review, nil
}

//FindAll reviews
func (r *repositoryStruct) FindAll(entityID string, entityKey string) (*[]Review, common.CustomError) {

	reviews := []Review{}

	rows, err := r.db.Query("SELECT id, content, createdByID FROM review WHERE entityID='" + entityID + "' AND entityKey='" + entityKey + "'")

	if err != nil {
		return nil, common.NewInternalServerError("error when tying to get all users", errors.New("database error"))
	}

	for rows.Next() {
		var id string
		var content string
		var createdByID string
		rows.Scan(&id, &content, &createdByID)
		review := NewReview(id, content, entityID, entityKey, createdByID)
		reviews = append(reviews, *review)
	}

	return &reviews, nil
}

//Destroy destroy a book by its id
func (r *repositoryStruct) DestroyByType(entityID string, entityKey string) common.CustomError {

	statement, _ := r.db.Prepare("DELETE FROM review WHERE entityID = ? AND entityKey = ?")

	_, err := statement.Exec(entityID, entityKey)
	if err != nil {
		return common.NewInternalServerError("error when tying to get all users", errors.New("database error"))
	}
	return nil

}
