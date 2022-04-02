package controller

// import (
// 	"net/http"
// 	"toporet/hop/goclean/presenter"
// 	. "toporet/hop/goclean/usecase"
// )

// type HttpWrapper[TU UseCase[TIn], TIn any, TOut, TP Presenter[]] interface {
// 	UseCase[TIn]
// 	presenter.HttpPresenter[TOut]
// }

// type Wrapper[TU UseCase[TIn], TIn any, TOut presenter.HttpPresenter] struct {
// 	UseCase[TIn]
// 	Presenter[TOut]
// }

// func (w Wrapper[TU, TIn, TOut]) Do(in TIn, resp http.ResponseWriter) {
// 	w.Handle(in)
// 	w.WriteResponse(resp)
// }
