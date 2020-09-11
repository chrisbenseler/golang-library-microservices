package domain

//Usecase books use case interface
type Usecase interface {
	AddOne(title string, year int) (Book, error)
	GetByID(id string) (Book, error)
	All() ([]Book, error)
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

	return u.repository.Save(title, year)
}

//GetByID method
func (u *usecaseStruct) GetByID(id string) (Book, error) {

	return u.repository.Get(id)
}

//All method
func (u *usecaseStruct) All() ([]Book, error) {

	return u.repository.All()
}
