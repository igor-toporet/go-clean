package controller

import (
	"net/http"
	"strings"
	. "toporet/hop/goclean/usecase/book"
)

type GetAllBooksFactory = UseCaseFactory[GetAllBooksUseCase]

func Books(f GetAllBooksFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case http.MethodGet:

			getAll(w, r, f)

		default:
			http.NotFound(w, r)
		}
	}
}

func getAll(w http.ResponseWriter, r *http.Request, f GetAllBooksFactory) {
	//
	// TODO: map request to use case input in other more complex use cases
	//
	u := f(w, r)
	u.Handle(GetAllBooksIn{})
}
