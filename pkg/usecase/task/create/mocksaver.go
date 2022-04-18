package create

import (
	"toporet/hop/goclean/pkg/entity"
)

type MockNewTaskSaver struct {
	id  *entity.TaskId
	err error
}

func (s *MockNewTaskSaver) SetupFailure(err error) {
	s.id = nil
	s.err = err
}

func (s *MockNewTaskSaver) SetupSuccess(id *entity.TaskId) {
	s.err = nil
	s.id = id
}

func (s *MockNewTaskSaver) SaveNewTask(t *entity.Task) (*entity.TaskId, error) {
	if s.id != nil {
		return s.id, nil
	}
	return nil, s.err
}
