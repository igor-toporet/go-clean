package controller

import (
	"net/http"
)

type UseCaseFactory[TUseCase any] func(
	http.ResponseWriter,
	*http.Request,
) TUseCase
