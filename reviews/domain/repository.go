package domain

import (
	"database/sql"
	"strconv"
)

//Repository book repository (persistence)
type Repository interface {
	Save(content string, entityKey string, entityID string) (Review, error)
	FindAll(entityID string, entityKey string) ([]Review, error)
}

type repositoryStruct struct {
	db *sql.DB
}

//NewReviewRepository create a new book repository
func NewReviewRepository(database *sql.DB) Repository {

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS review (id INTEGER PRIMARY KEY, content TEXT, entityID TEXT, entityKey TEXT)")
	statement.Exec()

	return &repositoryStruct{
		db: database,
	}
}

//Save review
func (r *repositoryStruct) Save(content string, entityID string, entityKey string) (Review, error) {

	review := NewReview("", content, entityID, entityKey)

	statement, _ := r.db.Prepare("INSERT INTO review (content, entityID, entityKey) VALUES (?, ?, ?)")

	result, err := statement.Exec(review.Content, review.EntityID, review.EntityKey)

	id, _ := result.LastInsertId()

	review.ID = strconv.FormatInt(int64(id), 10)

	return *review, err
}

//FindAll reviews
func (r *repositoryStruct) FindAll(entityID string, entityKey string) ([]Review, error) {

	reviews := []Review{}

	rows, _ := r.db.Query("SELECT id, content FROM review WHERE entityID='" + entityID + "' AND entityKey='" + entityKey + "'")

	for rows.Next() {
		var id string
		var content string
		rows.Scan(&id, &content)
		review := NewReview(id, content, entityID, entityKey)
		reviews = append(reviews, *review)
	}

	return reviews, nil
}
