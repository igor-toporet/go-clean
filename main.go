package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"toporet/hop/goclean/bootstrap"
	"toporet/hop/goclean/controller/task"
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
