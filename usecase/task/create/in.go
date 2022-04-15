package create

import "fmt"

type CreateTaskIn interface {
	TaskName() string
}

type createTaskIn struct {
	taskName string
}

func NewCreateTaskIn(taskName string) (CreateTaskIn, error) {
	if taskName == "" {
		return nil, fmt.Errorf("task name is required but got empty")
	}
	return &createTaskIn{taskName}, nil
}

func (in *createTaskIn) TaskName() string {
	return in.taskName
}
