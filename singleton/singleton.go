package singleton

import "sync"

type singleton struct {
	count int
}

var s *singleton
var once sync.Once

func New() *singleton {
	once.Do(func() {
		s = &singleton{}
	})
	return s
}

func (s *singleton) AddOne() {
	s.count++
}

func (s *singleton) GetCount() int {
	return s.count
}
