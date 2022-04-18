package create

type MockPresenter struct {
	out CreateTaskOut
}

func (p *MockPresenter) Present(out CreateTaskOut) {
	p.out = out
}

func (p *MockPresenter) Received() CreateTaskOut {
	return p.out
}
