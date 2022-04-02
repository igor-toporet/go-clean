package controller

import (
	"net/http"
	"toporet/hop/goclean/presenter"
	"toporet/hop/goclean/usecase"
)

type UseCaseFactory[
	TIn any,
	TOut any,
	TUseCase usecase.UseCase[TIn],
	TPresenter presenter.HttpPresenter[TOut],
] func(
	http.ResponseWriter,
	*http.Request,
) (
	TUseCase,
	TPresenter,
)
