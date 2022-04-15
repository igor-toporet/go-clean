package create

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewCreateTaskIn_Error(t *testing.T) {
	in, err := NewCreateTaskIn("")

	assert.Assert(t, in == nil)
	assert.Error(t, err, "task name is required but got empty")
}

func TestNewCreateTaskIn_Success(t *testing.T) {
	in, err := NewCreateTaskIn("new task")

	assert.NilError(t, err)
	assert.Assert(t, in != nil)
	assert.Equal(t, in.TaskName(), "new task")
}
