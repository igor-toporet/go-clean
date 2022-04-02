package bootstrap

import (
	"database/sql"
	"net/http"

	"toporet/hop/goclean/controller"
	"toporet/hop/goclean/gateway"

	. "toporet/hop/goclean/presenter/book"
	. "toporet/hop/goclean/usecase/book"
)

func Book(db *sql.DB) controller.UseCaseFactory[GetAllBooksIn, GetAllBooksOut, GetAllBooksUseCase, GetAllBooksPresenter] {
	return func(w http.ResponseWriter, r *http.Request) (GetAllBooksUseCase, GetAllBooksPresenter) {
		bookStore := gateway.NewBookStore(db)
		prsntr := NewGetAllBooksPresenter(w)
		return NewGetAllBooksUseCase(bookStore, prsntr), prsntr
	}
}
