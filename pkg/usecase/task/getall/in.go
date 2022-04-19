package getall

type GetAllTasksIn interface {
}

type getAllTasksIn struct {
}

func NewGetAllTasksIn() GetAllTasksIn {
	return &getAllTasksIn{}
}
