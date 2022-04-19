package task

import (
	"net/http"
	"strings"
)

func Handle(
	f CreateTaskUseCaseFactory,
	g GetAllTasksUseCaseFactory,

) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case http.MethodPost:

			f.create(w, r)

		case http.MethodGet:

			if r.URL.Path == "/tasks/" {
				g.getAll(w, r)
				break
			}
			fallthrough

		default:
			http.NotFound(w, r)
		}
	}
}
