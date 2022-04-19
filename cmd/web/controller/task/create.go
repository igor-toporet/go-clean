package task

import (
	"net/http"
	"toporet/hop/goclean/cmd/web/controller"
	"toporet/hop/goclean/cmd/web/controller/parser"
	"toporet/hop/goclean/pkg/usecase/task/create"
)

type payload struct {
	Name string
}

type CreateTaskUseCaseFactory controller.UseCaseFactory[create.CreateTaskUseCase]

func (f CreateTaskUseCaseFactory) handle(w http.ResponseWriter, r *http.Request) {

	toUseCaseInput := func(p payload) (*create.CreateTaskIn, error) {
		in, err := create.NewCreateTaskIn(p.Name)

		return &in, err
	}

	in, err := parser.ParseAndTranslate(r, toUseCaseInput)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uc := f(w, r)
	uc.Handle(*in)
}
