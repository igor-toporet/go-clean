package getall

import (
	"toporet/hop/goclean/pkg/entity"
)

type MockAllTasksFetcher struct {
	tasks []*entity.Task
	err   error
}

func (s *MockAllTasksFetcher) SetupFailure(err error) {
	s.tasks = nil
	s.err = err
}

func (s *MockAllTasksFetcher) SetupSuccess(tasks []*entity.Task) {
	s.err = nil
	s.tasks = tasks
}

func (s *MockAllTasksFetcher) FetchAll() ([]*entity.Task, error) {
	if s.tasks != nil {
		return s.tasks, nil
	}
	return nil, s.err
}
