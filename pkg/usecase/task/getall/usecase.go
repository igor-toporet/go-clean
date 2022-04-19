package getall

import (
	"toporet/hop/goclean/pkg/usecase"
)

// type presenter usecase.Presenter[GetAllTasksOut]

type GetAllTasksUseCase struct {
	fetcher AllTasksFetcher
	// presenter presenter
	presenter usecase.Presenter[GetAllTasksOut]
}

func NewGetAllTasksUseCase(
	s AllTasksFetcher,
	// p presenter,
	p usecase.Presenter[GetAllTasksOut],
) GetAllTasksUseCase {
	return GetAllTasksUseCase{s, p}
}

func (u *GetAllTasksUseCase) Handle(in GetAllTasksIn) {
	out := func() GetAllTasksOut {
		tasks, err := u.fetcher.FetchAll()
		if err != nil {
			return NewGetAllTasksOutDbGatewayError(err)
		}

		return NewGetAllTasksOutSuccess(tasks)
	}()

	u.presenter.Present(out)
}
