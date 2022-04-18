package entity

import (
	"errors"
)

type TaskId struct {
	val string
}

func NewTaskId(s string) (*TaskId, error) {
	if s == "" {
		return nil, errors.New("task ID cannot be empty")
	}

	return &TaskId{s}, nil
}

func (ti *TaskId) String() string {
	return ti.val
}
