package create

import (
	"toporet/hop/goclean/entity"
	"toporet/hop/goclean/usecase"
)

type presenter usecase.Presenter[CreateTaskOut]

type CreateTaskUseCase struct {
	taskSaver NewTaskSaver
	presenter presenter
}

func NewCreateTaskUseCase(
	s NewTaskSaver,
	p presenter,
) CreateTaskUseCase {
	return CreateTaskUseCase{s, p}
}

func (u *CreateTaskUseCase) Handle(in CreateTaskIn) {
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
