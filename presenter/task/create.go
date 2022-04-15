package task

import (
	"fmt"
	"net/http"
	"net/url"

	"toporet/hop/goclean/usecase/task/create"
)

type CreateTaskPresenter struct {
	w http.ResponseWriter
}

func NewCreateTaskPresenter(w http.ResponseWriter) *CreateTaskPresenter {
	return &CreateTaskPresenter{w: w}
}

func (p *CreateTaskPresenter) Present(o create.CreateTaskOut) {
	w := p.w

	statusCode, err, taskId := func() (int, error, *string) {
		t, err := o.TaskId()
		if err == nil {
			return http.StatusCreated, nil, t
		}
		if o.IsInputError(err) {
			return http.StatusBadRequest, err, nil
		}
		if o.IsDbGatewayError(err) {
			return http.StatusBadGateway, err, nil
			// or http.StatusInternalServerError
		}
		return http.StatusInternalServerError, err, nil
		// or panic("I don't know how to present this use case output")
	}()

	if err != nil {

		w.WriteHeader(statusCode)
		w.Write([]byte(err.Error()))

	} else {

		id := url.PathEscape(*taskId)

		// TODO: the route /tasks can be a constant
		// declared somewhere in the common web api related code
		w.Header().Set("Location", fmt.Sprintf("/tasks/%s", id))

		w.WriteHeader(statusCode)
	}
}
