package book

import (
	"toporet/hop/goclean/entity"
)

type BookStore interface {
	RetrieveAll() ([]entity.Book, error)
}

type GetAllBooksUseCase struct {
	bookStore BookStore
}

func NewGetAllBooksUseCase(s BookStore) GetAllBooksUseCase {
	return GetAllBooksUseCase{s}
}

func (u GetAllBooksUseCase) GetAllBooks(
//
// pass an input model later
//
) ([]entity.Book, error) {
	//
	// potentially more IO interactions and business logic here
	//
	return u.bookStore.RetrieveAll()
}
