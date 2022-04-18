package task

import "toporet/hop/goclean/pkg/entity"

type TaskSaver interface {
	SaveTask(t *entity.Task) error
}
