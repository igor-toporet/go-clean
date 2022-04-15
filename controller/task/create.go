package task

import (
	"net/http"
	"toporet/hop/goclean/controller"
	"toporet/hop/goclean/controller/parser"
	"toporet/hop/goclean/usecase/task/create"
)

type payload struct {
	Name string
}

type CreateTaskUseCaseFactory controller.UseCaseFactory[create.CreateTaskUseCase]

func (f CreateTaskUseCaseFactory) create(w http.ResponseWriter, r *http.Request) {

	in, err := parser.ParseAndTranslate(r, toUseCaseInput)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uc := f(w, r)
	uc.Handle(*in)
}

func toUseCaseInput(p payload) (*create.CreateTaskIn, error) {
	in, err := create.NewCreateTaskIn(p.Name)

	return &in, err
}
