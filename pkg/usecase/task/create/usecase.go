package create

import (
	"toporet/hop/goclean/pkg/entity"
	"toporet/hop/goclean/pkg/usecase"
)

type CreateTaskUseCase usecase.UseCase[CreateTaskIn]

type Presenter usecase.Presenter[CreateTaskOut]

type uc struct {
	taskSaver NewTaskSaver
	presenter Presenter
}

func NewCreateTaskUseCase(
	s NewTaskSaver,
	p Presenter,
) CreateTaskUseCase {
	return &uc{s, p}
}

func (u *uc) Handle(in CreateTaskIn) {
	out := func() CreateTaskOut {
		tn, err := entity.NewTaskName(in.TaskName())
		if err != nil {
			return NewCreateTaskOutInputError(err)
		}

		task := entity.NewTask(tn)
		id, err := u.taskSaver.SaveNewTask(task)
		if err != nil {
			return NewCreateTaskOutDbGatewayError(err)
		}

		return NewCreateTaskOutSuccess(id)
	}()

	u.presenter.Present(out)
}
