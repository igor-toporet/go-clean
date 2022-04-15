package create

import (
	"errors"
	"testing"
	"toporet/hop/goclean/entity"

	"gotest.tools/assert"
)

func TestNewCreateTaskOutDbGatewayError(t *testing.T) {
	out := NewCreateTaskOutDbGatewayError(
		errors.New("save failed"))

	tid, err := out.TaskId()

	assert.Check(t, tid == nil)
	assert.Check(t, out.IsDbGatewayError(err))
	assert.Check(t, !out.IsInputError(err))
	assert.Error(t, err, "save failed")
}

func TestNewCreateTaskOutInputError(t *testing.T) {
	out := NewCreateTaskOutInputError(
		errors.New("invalid input"))

	tid, err := out.TaskId()

	assert.Check(t, tid == nil)
	assert.Check(t, out.IsInputError(err))
	assert.Check(t, !out.IsDbGatewayError(err))
	assert.Error(t, err, "invalid input")
}

func TestNewCreateTaskOutSuccess(t *testing.T) {
	tid, err := entity.NewTaskId("task-id")
	assert.NilError(t, err)
	out := NewCreateTaskOutSuccess(tid)

	gotTid, err := out.TaskId()

	assert.Check(t, *gotTid == tid.String())
	assert.NilError(t, err)
}

func TestTaskId_panic_if_struct_was_not_properly_initialized(t *testing.T) {
	out := out{}

	arg := shouldPanic(t, func() {
		out.TaskId()
	})

	s, ok := arg.(string)
	assert.Check(t, ok)
	assert.Equal(t, s,
		"one of the properties must be initialized "+
			"(taskId, inputErr or dbErr) using corresponding "+
			"constuctor function (NewCreateTaskOut<...>)")
}

func shouldPanic(t *testing.T, f func()) (panicArg any) {
	t.Helper()
	defer func() {
		panicArg = recover()
	}()
	f()
	t.Errorf("should have panicked")
	return
}
