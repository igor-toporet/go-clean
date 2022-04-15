package entity

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestDefaultTask(t *testing.T) {
	task := new(Task)

	assert.Assert(t, is.Nil(task.Id()))
	assert.Assert(t, is.Nil(task.Name()))
	assert.Equal(t, task.Done(), false)
}

func TestNewTask(t *testing.T) {
	name, _ := NewTaskName("Buy milk")

	task := NewTask(name)

	assert.Assert(t, is.Nil(task.Id()))
	assert.Equal(t, task.Name().String(), "Buy milk")
	assert.Equal(t, task.Done(), false)
}

func TestNewFromExisting_MissingId_Error(t *testing.T) {
	name, _ := NewTaskName("dummy")

	_, err := NewTaskFromExisting(nil, name, true)

	assert.Error(t, err, "nil task id")
}

func TestNewFromExisting_MissingName_Error(t *testing.T) {
	id, _ := NewTaskId("id")

	_, err := NewTaskFromExisting(id, nil, true)

	assert.Error(t, err, "nil task name")
}

func TestNewFromExisting_AllProps_Success(t *testing.T) {
	id, _ := NewTaskId("id-123")
	name, _ := NewTaskName("buy milk")

	task, err := NewTaskFromExisting(id, name, true)

	assert.NilError(t, err)
	assert.Equal(t, task.Id().String(), "id-123")
	assert.Equal(t, task.Name().String(), "buy milk")
	assert.Equal(t, task.Done(), true)
}

func TestMarkComplete(t *testing.T) {
	name, _ := NewTaskName("Buy milk")
	task := NewTask(name)

	task.MarkComplete()

	assert.Equal(t, task.Done(), true)
}

func TestMarkIncomplete(t *testing.T) {
	name, _ := NewTaskName("Buy milk")
	task := NewTask(name)
	task.MarkComplete()

	task.MarkIncomplete()

	assert.Equal(t, task.Done(), false)
}

func TestString_NewTask(t *testing.T) {
	name, err := NewTaskName("foo")
	assert.NilError(t, err)
	task := NewTask(name)

	assert.Equal(t, task.String(), "[_] foo <new>")
}

func TestString_NewTask_Done(t *testing.T) {
	name, err := NewTaskName("foo")
	assert.NilError(t, err)
	task := NewTask(name)
	task.MarkComplete()

	assert.Equal(t, task.String(), "[✓] foo <new>")
}

func TestString_ExistingTask(t *testing.T) {
	id, err := NewTaskId("task-id")
	assert.NilError(t, err)
	name, err := NewTaskName("task name")
	assert.NilError(t, err)
	task, err := NewTaskFromExisting(id, name, false)
	assert.NilError(t, err)

	assert.Equal(t, task.String(), "[_] task name (task-id)")
}

func TestString_ExistingTask_Done(t *testing.T) {
	id, err := NewTaskId("task-id")
	assert.NilError(t, err)
	name, err := NewTaskName("task name")
	assert.NilError(t, err)
	task, err := NewTaskFromExisting(id, name, false)
	assert.NilError(t, err)
	task.MarkComplete()

	assert.Equal(t, task.String(), "[✓] task name (task-id)")
}
