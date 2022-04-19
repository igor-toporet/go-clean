package create

import (
	"errors"
	"testing"
	"toporet/hop/goclean/pkg/entity"
	"toporet/hop/goclean/pkg/usecase"

	"gotest.tools/assert"
)

func makeUseCase() (
	*CreateTaskUseCase,
	*MockNewTaskSaver,
	*usecase.MockPresenter[CreateTaskOut],
) {
	s := &MockNewTaskSaver{}
	p := &usecase.MockPresenter[CreateTaskOut]{}
	uc := NewCreateTaskUseCase(s, p)

	return &uc, s, p
}

func TestHandle_InputError(t *testing.T) {
	uc, _, p := makeUseCase()
	in, err := NewCreateTaskIn(" ")
	assert.NilError(t, err)

	uc.Handle(in)

	out := p.Received()
	_, err = out.TaskId()
	assert.Check(t, out.IsInputError(err))
	assert.Assert(t, len(err.Error()) > 0)
}

func TestHandle_DbError(t *testing.T) {
	uc, s, p := makeUseCase()
	s.SetupFailure(errors.New("save failure"))

	in, err := NewCreateTaskIn("foo")
	assert.NilError(t, err)

	uc.Handle(in)

	out := p.Received()
	_, err = out.TaskId()
	assert.Check(t, out.IsDbGatewayError(err))
	assert.Assert(t, len(err.Error()) > 0)
}

func TestHandle_Success(t *testing.T) {
	uc, s, p := makeUseCase()
	id, err := entity.NewTaskId("task-id")
	assert.NilError(t, err)
	s.SetupSuccess(id)

	in, err := NewCreateTaskIn("new task")
	assert.NilError(t, err)

	uc.Handle(in)

	out := p.Received()
	tid, err := out.TaskId()
	assert.Check(t, *tid == id.String())
	assert.NilError(t, err)
}
