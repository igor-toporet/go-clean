package create

import "toporet/hop/goclean/pkg/entity"

type NewTaskSaver interface {
	SaveNewTask(t *entity.Task) (*entity.TaskId, error)
}
