package getall

import "toporet/hop/goclean/pkg/entity"

type GetAllTasksOut interface {
	AllTasks() ([]*entity.Task, error)
	IsDbGatewayError(e error) bool
}

type out struct {
	allTasks    []*entity.Task
	databaseErr error
}

func NewGetAllTasksOutSuccess(tasks []*entity.Task) GetAllTasksOut {
	return &out{allTasks: tasks}
}

func NewGetAllTasksOutDbGatewayError(err error) GetAllTasksOut {
	return &out{databaseErr: err}
}

func (o *out) AllTasks() ([]*entity.Task, error) {
	if o.allTasks != nil {
		return o.allTasks, nil
	}
	if o.databaseErr != nil {
		return nil, o.databaseErr
	}
	panic("one of the properties must be initialized " +
		"(allTasks or dbErr) using corresponding " +
		"constuctor function (NewGetAllTasksOut<...>)")
}

func (o *out) IsDbGatewayError(e error) bool {
	return o.databaseErr == e
}
