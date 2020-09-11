package domain

//Usecase books use case interface
type Usecase interface {
	AddOne(title string, year int) (Book, error)
	GetByID(id string) (Book, error)
}

type usecaseStruct struct {
	repository Repository
}

//NewBookUsecase create a new book use case
func NewBookUsecase(repository Repository) Usecase {

	return &usecaseStruct{
		repository: repository,
	}
}

//AddOne method
func (u *usecaseStruct) AddOne(title string, year int) (Book, error) {

	book, err := u.repository.Save(title, year)

	return book, err
}

//GetByID method
func (u *usecaseStruct) GetByID(id string) (Book, error) {

	book, err := u.repository.Get(id)

	return book, err
}
