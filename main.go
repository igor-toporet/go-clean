package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"toporet/hop/goclean/bootstrap"
	"toporet/hop/goclean/controller"
)

func main() {
	db, err := sql.Open("postgres",
		"postgres://postgres:Password1@localhost/bookstore?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	getAll := bootstrap.Book(db)

	mux.HandleFunc("/books", controller.Books(getAll))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
