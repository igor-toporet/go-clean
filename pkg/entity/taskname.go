package entity

import (
	"fmt"
	"strings"
)

type TaskName struct {
	val string
}

func NewTaskName(s string) (*TaskName, error) {
	trim := strings.TrimSpace(s)

	if trim == "" {
		return nil, fmt.Errorf("task name is empty or whitespace (%q)", s)
	}
	return &TaskName{s}, nil
}

func (tn *TaskName) String() string {
	return tn.val
}
