package entity

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestNewTaskId_EmptyInput_Err(t *testing.T) {
	id, err := NewTaskId("")

	assert.Assert(t, is.Nil(id))
	assert.Error(t, err, "task ID cannot be empty")
}

func TestNewTaskId_NonEmptyInput_Success(t *testing.T) {
	id, err := NewTaskId("123")

	assert.Equal(t, id.String(), "123")
	assert.NilError(t, err)
}
