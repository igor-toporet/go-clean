package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"toporet/hop/goclean/cmd/web/bootstrap"
	"toporet/hop/goclean/cmd/web/controller/task"

	_ "github.com/lib/pq"
)

func main() {

	connStr := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	createTask := bootstrap.Task(db)

	mux.HandleFunc("/tasks", task.Handle(createTask))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
