package bootstrap

import (
	"database/sql"
	"net/http"

	controller "toporet/hop/goclean/controller/task"
	presenter "toporet/hop/goclean/presenter/task"

	"toporet/hop/goclean/gateway"
	"toporet/hop/goclean/usecase/task/create"
)

func Task(db *sql.DB) controller.CreateTaskUseCaseFactory {
	return func(w http.ResponseWriter, r *http.Request) create.CreateTaskUseCase {
		store := gateway.NewTaskStore(db)

		ucCreate := create.NewCreateTaskUseCase(store, presenter.NewCreateTaskPresenter(w))

		return ucCreate // , TODO: return more use case factories
	}
}
