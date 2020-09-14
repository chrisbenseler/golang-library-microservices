package domain

//Usecase reviews use case interface
type Usecase interface {
	AddBookReview(bookID string, content string) (Review, error)
	AllFromBook(bookID string) ([]Review, error)
}

type usecaseStruct struct {
	repository Repository
}

//NewReviewUsecase a new book use case
func NewReviewUsecase(repository Repository) Usecase {

	return &usecaseStruct{
		repository: repository,
	}
}

//AddBookReview method
func (u *usecaseStruct) AddBookReview(bookID string, content string) (Review, error) {
	return u.repository.Save(content, "book", bookID)
}

//AllFromBook method
func (u *usecaseStruct) AllFromBook(bookID string) ([]Review, error) {
	return u.repository.FindAll("book", bookID)
}
