package book

import (
	"encoding/json"
	"log"
	"net/http"

	"toporet/hop/goclean/usecase/book"
)

type GetAllBooksPresenter struct {
	w http.ResponseWriter
}

func NewGetAllBooksPresenter(w http.ResponseWriter) GetAllBooksPresenter {
	return GetAllBooksPresenter{w: w}
}

func (p GetAllBooksPresenter) Present(o book.GetAllBooksOut) {
	books := o.Books()
	err := o.Err()
	w := p.w

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	} else {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := make(map[string]any)
		resp["status"] = http.StatusOK
		resp["statusText"] = "Status OK"
		resp["data"] = books

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	}
}
