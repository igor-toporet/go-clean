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
) (
	controller.CreateTaskUseCaseFactory,
	controller.GetAllTasksUseCaseFactory,
) {
	store := gateway.NewTaskStore(db)

	return func(w http.ResponseWriter, r *http.Request) create.CreateTaskUseCase {

			return create.NewCreateTaskUseCase(store, presenter.NewCreateTaskPresenter(w))

		},

		// func(w http.ResponseWriter, r *http.Request) getall.GetAllTasksUseCase {

		// 	return getall.NewGetAllTasksUseCase(store, presenter.NewGetAllTasksPresenter(w))
		// }

		toFactory[getall.GetAllTasksUseCase, getall.AllTasksFetcher](
			presenter.NewGetAllTasksPresenter,
			getall.NewGetAllTasksUseCase,
			store,
		)
	/*

	   cannot use getall.NewGetAllTasksUseCase (value of type

	   func(
	   	s getall.AllTasksFetcher,
	   	p getall.presenter,
	   ) getall.GetAllTasksUseCase

	   ) as

	   func(
	   	s getall.AllTasksFetcher,
	   	p *"toporet/hop/goclean/cmd/web/presenter/task".GetAllTasksPresenter,
	   ) getall.GetAllTasksUseCase

	   value in argument to

	   toFactory[getall.GetAllTasksUseCase, getall.AllTasksFetcher, presenter.GetAllTasksPresenter]

	*/

}

func toFactory[TUseCase any, TStore any, TPresenter any](
	pf func(w http.ResponseWriter) TPresenter,
	f func(s TStore, p TPresenter) TUseCase,
	store TStore,
) func(w http.ResponseWriter, r *http.Request) TUseCase {

	return func(w http.ResponseWriter, r *http.Request) TUseCase {
		p := pf(w)
		return f(store, p)
	}
}
