package task

import (
	"encoding/json"
	"net/http"

	"toporet/hop/goclean/pkg/entity"
	"toporet/hop/goclean/pkg/usecase/task/getall"
)

type getAllTasksPresenter struct {
	w http.ResponseWriter
}

func NewGetAllTasksPresenter(w http.ResponseWriter) getall.Presenter {
	return &getAllTasksPresenter{w: w}
}

func (p *getAllTasksPresenter) Present(o getall.GetAllTasksOut) {
	w := p.w

	statusCode, tasks, err := func() (int, []*entity.Task, error) {
		tasks, err := o.AllTasks()
		if err == nil {
			return http.StatusOK, tasks, nil
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

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(statusCode)

		m := envelope[[]task]{toDtoSlice(tasks)}

		json.NewEncoder(w).Encode(m)
	}
}

//
// TODO: move to the payloads folder later
//       (to be created under the web folder)
//
type task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type envelope[TPayload any] struct {
	Data TPayload `json:"data"`
}

func toDtoSlice(tasks []*entity.Task) []task {
	var r []task

	for _, t := range tasks {
		dto := task{Id: t.Id().String(), Name: t.Name().String(), Done: t.Done()}
		r = append(r, dto)
	}
	return r
}
