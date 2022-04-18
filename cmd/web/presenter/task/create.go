package task

import (
	"fmt"
	"net/http"
	"net/url"

	"toporet/hop/goclean/pkg/usecase/task/create"
)

type CreateTaskPresenter struct {
	w http.ResponseWriter
}

func NewCreateTaskPresenter(w http.ResponseWriter) *CreateTaskPresenter {
	return &CreateTaskPresenter{w: w}
}

func (p *CreateTaskPresenter) Present(o create.CreateTaskOut) {
	w := p.w

	statusCode, taskId, err := func() (int, *string, error) {
		t, err := o.TaskId()
		if err == nil {
			return http.StatusCreated, t, nil
		}
		if o.IsInputError(err) {
			return http.StatusBadRequest, nil, err
		}
		if o.IsDbGatewayError(err) {
			return http.StatusBadGateway, nil, err
			// or http.StatusInternalServerError
		}
		return http.StatusInternalServerError, nil, err
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
