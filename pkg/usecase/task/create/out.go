package create

import "toporet/hop/goclean/pkg/entity"

type CreateTaskOut interface {
	TaskId() (*string, error)
	IsDbGatewayError(e error) bool
	IsInputError(e error) bool
}

type out struct {
	taskId      *string
	inputErr    error
	databaseErr error
}

func NewCreateTaskOutSuccess(id *entity.TaskId) CreateTaskOut {
	i := id.String()
	return &out{taskId: &i}
}

func NewCreateTaskOutDbGatewayError(err error) CreateTaskOut {
	return &out{databaseErr: err}
}

func NewCreateTaskOutInputError(err error) CreateTaskOut {
	return &out{inputErr: err}
}

func (o *out) TaskId() (*string, error) {
	if o.taskId != nil {
		return o.taskId, nil
	}
	if o.inputErr != nil {
		return nil, o.inputErr
	}
	if o.databaseErr != nil {
		return nil, o.databaseErr
	}
	panic("one of the properties must be initialized " +
		"(taskId, inputErr or dbErr) using corresponding " +
		"constuctor function (NewCreateTaskOut<...>)")
}

func (o *out) IsDbGatewayError(e error) bool {
	return o.databaseErr == e
}

func (o *out) IsInputError(e error) bool {
	return o.inputErr == e
}
