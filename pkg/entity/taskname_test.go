package entity

import (
	"testing"

	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
)

func TestNewTask_InvalidName(t *testing.T) {
	tn, err := NewTaskName(" \t ")

	assert.Error(t, err, "task name is empty or whitespace (\" \\t \")")

	assert.Assert(t, cmp.Nil(tn))
}

func TestNewTask_ValidName(t *testing.T) {
	n, err := NewTaskName("Buy milk")

	assert.NilError(t, err)

	assert.Equal(t, n.String(), "Buy milk")
}
