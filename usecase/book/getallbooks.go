package book

import (
	"toporet/hop/goclean/entity"
	"toporet/hop/goclean/usecase"
)

type GetAllBooksUseCase struct {
	bookStore BookStore
	presenter Presenter
}

func NewGetAllBooksUseCase(
	s BookStore,
	p Presenter,
) GetAllBooksUseCase {
	return GetAllBooksUseCase{s, p}
}

func (u GetAllBooksUseCase) Handle(in GetAllBooksIn) {
	//
	// potentially more IO interactions and business logic here
	//
	b, e := u.bookStore.RetrieveAll()

	u.presenter.Present(GetAllBooksOut{b, e})
}

type GetAllBooksIn struct {
}

type Presenter = usecase.Presenter[GetAllBooksOut]

type BookStore interface {
	RetrieveAll() ([]entity.Book, error)
}

type GetAllBooksOut struct {
	books []entity.Book
	err   error
}

func (o GetAllBooksOut) Books() []entity.Book {
	return o.books
}

func (o GetAllBooksOut) Err() error {
	return o.err
}
