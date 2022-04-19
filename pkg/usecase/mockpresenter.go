package usecase

type MockPresenter[TOut any] struct {
	out TOut
}

func (p *MockPresenter[TOut]) Present(out TOut) {
	p.out = out
}

func (p *MockPresenter[TOut]) Received() TOut {
	return p.out
}
