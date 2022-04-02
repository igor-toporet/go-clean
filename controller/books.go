package controller

import (
	"fmt"
	"net/http"
	"strings"
	"toporet/hop/goclean/presenter/book"
	. "toporet/hop/goclean/usecase/book"
)

// type GetAllBooksFactory func {
// 	UseCaseFactory[GetAllBooksIn, GetAllBooksOut, GetAllBooksUseCase, book.GetAllBooksPresenter]
// }

func Books(f UseCaseFactory[GetAllBooksIn, GetAllBooksOut, GetAllBooksUseCase, book.GetAllBooksPresenter]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case http.MethodGet:

			u, p := f(w, r)

			getAll(u)

			fmt.Printf("Presenter: %v", p)

		default:
			http.NotFound(w, r)
		}
	}
}

func getAll(u GetAllBooksUseCase) {
	u.Handle(GetAllBooksIn{})
}
