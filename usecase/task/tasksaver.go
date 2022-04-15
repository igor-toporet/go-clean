package task

import "toporet/hop/goclean/entity"

type TaskSaver interface {
	SaveTask(t *entity.Task) error
}
