package task

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	presenter "toporet/hop/goclean/cmd/web/presenter/task"
	"toporet/hop/goclean/pkg/entity"
	uc "toporet/hop/goclean/pkg/usecase/task/create"

	"gotest.tools/assert"
)

const (
	contentType string = "content-type"
)

func TestCreate_InvalidRequest(t *testing.T) {
	mockDb := &uc.MockNewTaskSaver{}
	handler := Handle(bootstrap(mockDb))

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
	mockDb := &uc.MockNewTaskSaver{}
	id, err := entity.NewTaskId("42")
	if err != nil {
		t.Fatal(err.Error())
	}
	mockDb.SetupSuccess(id)
	handler := Handle(bootstrap(mockDb))

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

// TODO: db failure test

// TODO: invalid route / not found test

func bootstrap(saver *uc.MockNewTaskSaver) CreateTaskUseCaseFactory {
	return func(w http.ResponseWriter, r *http.Request) uc.CreateTaskUseCase {
		return uc.NewCreateTaskUseCase(
			saver,
			presenter.NewCreateTaskPresenter(w),
		)
	}
}
