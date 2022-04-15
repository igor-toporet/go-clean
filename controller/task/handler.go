package task

import (
	"net/http"
	"strings"
)

func Handle(

	f CreateTaskUseCaseFactory,

	// TODO: more use case factories

) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case http.MethodPost:

			f.create(w, r)

		default:
			http.NotFound(w, r)
		}
	}
}
