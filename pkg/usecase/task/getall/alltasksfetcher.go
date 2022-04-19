package getall

import "toporet/hop/goclean/pkg/entity"

type AllTasksFetcher interface {
	FetchAll() ([]*entity.Task, error)
}
