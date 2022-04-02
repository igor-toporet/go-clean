package usecase

type Presenter[TOut any] interface {
	Present(out TOut)
}

type UseCase[TIn any] interface {
	Handle(in TIn)
}
