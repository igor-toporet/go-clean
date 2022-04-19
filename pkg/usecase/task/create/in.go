package create

import "fmt"

type CreateTaskIn interface {
	TaskName() string
}

type in struct {
	taskName string
}

func NewCreateTaskIn(taskName string) (CreateTaskIn, error) {
	if taskName == "" {
		return nil, fmt.Errorf("task name is required but got empty")
	}
	return &in{taskName}, nil
}

func (in *in) TaskName() string {
	return in.taskName
}
