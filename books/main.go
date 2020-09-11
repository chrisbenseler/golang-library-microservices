package main

import (
	"fmt"
	"librarymanager/books/domain"
)

func main() {

	fmt.Print("Books process")

	service := domain.NewBookService()

	repository := domain.NewBookRepository(service)

	usecase := domain.NewBookUsecase(repository)

	//usecase.AddOne("some title", 2020)

	book, _ := usecase.GetByID("w7eo7Lq7jgykVvVS")
	fmt.Print(book)
}
