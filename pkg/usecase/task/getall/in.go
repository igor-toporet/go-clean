package getall

type GetAllTasksIn interface {
}

type in struct {
}

func NewGetAllTasksIn() GetAllTasksIn {
	return &in{}
}
