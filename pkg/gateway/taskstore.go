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

type task struct {
	id   int
	name string
	done bool
}

func (s *TaskStore) FetchAll() ([]*entity.Task, error) {

	var tasks []*entity.Task

	rows, err := s.db.Query(
		"SELECT id, name, done FROM tasks")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t task
		if err := rows.Scan(&t.id, &t.name, &t.done); err != nil {
			return nil, err
		}
		task, err := from(t)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func from(t task) (*entity.Task, error) {
	id, err := entity.NewTaskId(fmt.Sprint(t.id))
	if err == nil {
		name, err := entity.NewTaskName(t.name)
		if err == nil {
			task, err := entity.NewTaskFromExisting(id, name, t.done)
			return task, err
		}
	}
	return nil, err
}
func (s *TaskStore) SaveTask(t *entity.Task) error {
	//
	// TODO: implement saving for realz
	//
	return nil
}
