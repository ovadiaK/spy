package spy

import "github.com/ovadiaK/spy/internal/app"

type Spy struct {
	GetFunc metric
	SetFunc metric
	app.Api
}

func New() Spy {
	application := app.New()
	return Spy{Api: &application}
}

type metric struct {
	called int
}

func (m metric) WasCalled() bool {
	if m.called > 0 {
		return true
	}
	return false
}
func (m *metric) call() {
	m.called++
}

//
//func (s *Spy) Set(key, value string) {
//	s.SetFunc.call()
//	s.app.Set(key, value)
//}

func (s *Spy) Get(key string) string {
	s.GetFunc.call()
	return s.Api.Get(key)
}
