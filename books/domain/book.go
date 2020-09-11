package domain

//Book book model
type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

//NewBook new book entity
func NewBook(id string, title string, year int) *Book {
	return &Book{
		ID:    id,
		Title: title,
		Year:  year,
	}
}
