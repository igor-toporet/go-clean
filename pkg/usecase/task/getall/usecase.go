package getall

import (
	"toporet/hop/goclean/pkg/usecase"
)

type GetAllTasksUseCase usecase.UseCase[GetAllTasksIn]

type Presenter usecase.Presenter[GetAllTasksOut]

type uc struct {
	fetcher   AllTasksFetcher
	presenter Presenter
}

func NewGetAllTasksUseCase(
	f AllTasksFetcher,
	p Presenter,
) GetAllTasksUseCase {
	return &uc{f, p}
}

func (u *uc) Handle(in GetAllTasksIn) {
	out := func() GetAllTasksOut {
		tasks, err := u.fetcher.FetchAll()
		if err != nil {
			return NewGetAllTasksOutDbGatewayError(err)
		}

		return NewGetAllTasksOutSuccess(tasks)
	}()

	u.presenter.Present(out)
}
