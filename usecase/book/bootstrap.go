package book

import (
	"database/sql"

	"toporet/hop/goclean/gateway"
)

func Bootstrap(db *sql.DB) GetAllBooksUseCase {

	bookStore := gateway.NewBookStore(db)

	return NewGetAllBooksUseCase(bookStore)
}
