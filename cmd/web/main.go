package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"toporet/hop/goclean/cmd/web/bootstrap"

	_ "github.com/lib/pq"
)

func main() {

	connStr := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	bootstrap.Task(db, mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
