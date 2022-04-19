package task

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	presenter "toporet/hop/goclean/cmd/web/presenter/task"
	"toporet/hop/goclean/pkg/entity"
	"toporet/hop/goclean/pkg/usecase/task/create"
	"toporet/hop/goclean/pkg/usecase/task/getall"

	"gotest.tools/assert"
)

const (
	contentType string = "content-type"
)

func TestRoot_UnrecognizedVerb_NotFound(t *testing.T) {

	handler := Handle(bootstrap(nil, nil))

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("PATCH", "/tasks/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusNotFound)
	assert.Equal(t, rr.Header().Get(contentType), "text/plain; charset=utf-8")
	assert.Equal(t, rr.Body.String(), "404 page not found\n")
}

func TestCreate_InvalidRequest(t *testing.T) {
	handler := Handle(bootstrap(nil, nil))

	cases := []struct {
		contentType, reqBody, expectedRespBody string
	}{
		{"", "ignore", `invalid request Content-Type (expected "application/json", received [""])`},
		{"text", "ignore", `invalid request Content-Type (expected "application/json", received ["text"])`},
		{"application/json", "", "missing request body"},
	}
	for _, c := range cases {

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/tasks", strings.NewReader(c.reqBody))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add(contentType, c.contentType)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, rr.Header().Get(contentType), "text/plain; charset=utf-8")
		assert.Equal(t, rr.Body.String(), c.expectedRespBody+"\n")
	}
}

func TestCreate_Success(t *testing.T) {
	mockDb := &create.MockNewTaskSaver{}
	id, err := entity.NewTaskId("42")
	if err != nil {
		t.Fatal(err.Error())
	}
	mockDb.SetupSuccess(id)
	handler := Handle(bootstrap(mockDb, nil))

	rr := httptest.NewRecorder()
	body := `{"name": "a task"}`
	req, err := http.NewRequest("POST", "/tasks", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add(contentType, "application/json")

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)
	assert.Equal(t, rr.Header().Get("Location"), "/tasks/42")
	assert.Equal(t, rr.Header().Get(contentType), "")
	assert.Equal(t, rr.Body.String(), "")
}

func TestGetAll_Success(t *testing.T) {
	mockDb := &getall.MockAllTasksFetcher{}
	mockDb.SetupSuccess(createDummyTasks([]string{"foo", "bar"}))
	handler := Handle(bootstrap(nil, mockDb))

	rr := httptest.NewRecorder()

	// Note that the trailing slash at the end of the path simulates the runtime behavior.
	// Runtime requests without a trailing slash still hit the router with the slash.
	req, err := http.NewRequest("GET", "/tasks/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, rr.Header().Get(contentType), "application/json")
	assert.Equal(t, rr.Body.String(),
		`{"data":[`+
			`{"id":"t-1","name":"foo","done":false},`+
			`{"id":"t-2","name":"bar","done":false}`+
			"]}\n")
}

func createDummyTasks(names []string) []*entity.Task {

	createTask := func(i int, name string) (*entity.Task, error) {
		id, err := entity.NewTaskId(fmt.Sprintf("t-%d", i+1))
		if err == nil {
			name, err := entity.NewTaskName(name)
			if err == nil {
				return entity.NewTaskFromExisting(id, name, false)
			}
		}
		return nil, err
	}

	var tasks []*entity.Task

	for i, n := range names {
		t, err := createTask(i, n)
		if err == nil {
			tasks = append(tasks, t)
		} else {
			panic(err)
		}
	}

	return tasks
}

// TODO: db failure test

// TODO: invalid route / not found test

func bootstrap(
	saver *create.MockNewTaskSaver,
	fetcher *getall.MockAllTasksFetcher,
) (
	CreateTaskUseCaseFactory,
	GetAllTasksUseCaseFactory,
) {
	return func(w http.ResponseWriter, r *http.Request) create.CreateTaskUseCase {
			return create.NewCreateTaskUseCase(
				saver, presenter.NewCreateTaskPresenter(w))
		},
		func(w http.ResponseWriter, r *http.Request) getall.GetAllTasksUseCase {
			return getall.NewGetAllTasksUseCase(
				fetcher, presenter.NewGetAllTasksPresenter(w))
		}
}
