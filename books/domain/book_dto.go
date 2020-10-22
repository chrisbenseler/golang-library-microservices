package domain

//BookDTO book data transfer object
type BookDTO struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}
