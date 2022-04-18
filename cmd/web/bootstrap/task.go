package bootstrap

import (
	"database/sql"
	"net/http"

	controller "toporet/hop/goclean/cmd/web/controller/task"
	presenter "toporet/hop/goclean/cmd/web/presenter/task"

	"toporet/hop/goclean/pkg/gateway"
	"toporet/hop/goclean/pkg/usecase/task/create"
)

func Task(db *sql.DB) controller.CreateTaskUseCaseFactory {
	return func(w http.ResponseWriter, r *http.Request) create.CreateTaskUseCase {
		store := gateway.NewTaskStore(db)

		ucCreate := create.NewCreateTaskUseCase(store, presenter.NewCreateTaskPresenter(w))

		return ucCreate // , TODO: return more use case factories
	}
}
