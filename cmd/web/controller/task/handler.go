package task

import (
	"net/http"
	"strings"
)

const (
	RoutePath = "/tasks/"
)

func Handle(
	create CreateTaskUseCaseFactory,
	getAll GetAllTasksUseCaseFactory,

) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case http.MethodPost:

			create.handle(w, r)

		case http.MethodGet:

			if r.URL.Path == RoutePath {
				getAll.handle(w, r)
				break
			}
			fallthrough

		default:
			http.NotFound(w, r)
		}
	}
}
