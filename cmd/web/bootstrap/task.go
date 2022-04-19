package bootstrap

import (
	"database/sql"
	"net/http"

	controller "toporet/hop/goclean/cmd/web/controller/task"
	presenter "toporet/hop/goclean/cmd/web/presenter/task"

	"toporet/hop/goclean/pkg/gateway"
	"toporet/hop/goclean/pkg/usecase/task/create"
	"toporet/hop/goclean/pkg/usecase/task/getall"
)

func Task(
	db *sql.DB,
	mux *http.ServeMux,
) {
	store := gateway.NewTaskStore(db)

	create := func(w http.ResponseWriter, r *http.Request) create.CreateTaskUseCase {
		return create.NewCreateTaskUseCase(store, presenter.NewCreateTaskPresenter(w))
	}

	getAll := func(w http.ResponseWriter, r *http.Request) getall.GetAllTasksUseCase {
		return getall.NewGetAllTasksUseCase(store, presenter.NewGetAllTasksPresenter(w))
	}

	mux.HandleFunc(
		controller.RoutePath,
		controller.Handle(
			create,
			getAll,
		),
	)
}
