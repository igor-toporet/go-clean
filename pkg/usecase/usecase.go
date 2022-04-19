package usecase

type UseCase[TIn any] interface {
	Handle(in TIn)
}
