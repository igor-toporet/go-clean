package presenter

import (
	"toporet/hop/goclean/usecase"
)

type HttpPresenter[TOut any] interface {
	usecase.Presenter[TOut]
}
