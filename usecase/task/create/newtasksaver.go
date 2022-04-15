package create

import "toporet/hop/goclean/entity"

type NewTaskSaver interface {
	SaveNewTask(t *entity.Task) (*entity.TaskId, error)
}
