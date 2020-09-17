package domain

//Book book model
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
	CreatedByID string `json:"createdByID"`
}

//NewBook new book entity
func NewBook(id string, title string, year int, createdByID string) *Book {
	return &Book{
		ID:          id,
		Title:       title,
		Year:        year,
		CreatedByID: createdByID,
	}
}
