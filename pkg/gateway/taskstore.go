package gateway

import (
	"database/sql"
	"fmt"

	"toporet/hop/goclean/pkg/entity"
)

// type IGetAllBooks func() ([]entity.Book, error)

type TaskStore struct {
	db *sql.DB
}

func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{db}
}

func (s *TaskStore) SaveNewTask(t *entity.Task) (*entity.TaskId, error) {

	row := s.db.QueryRow(
		"INSERT INTO tasks (name, done) VALUES($1, $2) RETURNING id",
		t.Name().String(), t.Done())

	err := row.Err()
	if err != nil {
		return nil, err
	}

	lastInsertId := 0
	err = row.Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	taskId, err := entity.NewTaskId(fmt.Sprint(lastInsertId))

	return taskId, err
}

func (s *TaskStore) SaveTask(t *entity.Task) error {
	//
	// TODO: implement saving for realz
	//
	return nil
}

// func (s TaskStore) RetrieveAll() ([]entity.Book, error) {
// 	rows, err := s.db.Query("SELECT * FROM books")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var books []entity.Book

// 	for rows.Next() {
// 		var bk entity.Book

// 		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
// 		if err != nil {
// 			return nil, err
// 		}

// 		books = append(books, bk)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return books, nil
// }
