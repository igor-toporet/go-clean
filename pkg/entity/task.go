package entity

import (
	"errors"
	"fmt"
)

type Task struct {
	id   *TaskId
	name *TaskName
	done bool
}

func NewTask(name *TaskName) *Task {
	return &Task{nil, name, false}
}

func NewTaskFromExisting(id *TaskId, name *TaskName, done bool) (*Task, error) {
	if id == nil {
		return nil, errors.New("nil task id")
	}

	if name == nil {
		return nil, errors.New("nil task name")
	}

	return &Task{id, name, done}, nil
}

func (t *Task) Id() *TaskId {
	return t.id
}

func (t *Task) Name() *TaskName {
	return t.name
}

func (t *Task) Done() bool {
	return t.done
}

func (t *Task) MarkComplete() {
	t.done = true
}

func (t *Task) MarkIncomplete() {
	t.done = false
}

//
// Used for debugging mainly
//
// Example output:
//
//  - new incomplete task     [_] a task <new>
//  - saved task              [_] a task (an-id)
//  - task marked complete    [✓] a task (an-id)
//  - task marked incomplete  [_] a task (an-id)
//
func (t *Task) String() string {
	done := "_"
	if t.done {
		done = "✓"
	}

	id := "<new>"
	if t.id != nil {
		id = fmt.Sprintf("(%s)", t.id)
	}

	return fmt.Sprintf("[%s] %s %s", done, t.name, id)
}
