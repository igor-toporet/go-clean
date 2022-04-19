package task

import (
	"net/http"
	"toporet/hop/goclean/cmd/web/controller"
	"toporet/hop/goclean/pkg/usecase/task/getall"
)

type GetAllTasksUseCaseFactory controller.UseCaseFactory[getall.GetAllTasksUseCase]

func (f GetAllTasksUseCaseFactory) handle(w http.ResponseWriter, r *http.Request) {

	in := getall.NewGetAllTasksIn()

	uc := f(w, r)
	uc.Handle(in)
}
