package usecase

type Presenter[TOut any] interface {
	Present(out TOut)
}
