package vmc

import (
	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/asf/atf"
)

type Executor struct {
	asf.Program
	Stack
	Err error
}

// Reset state
//
//go:inline
func (s *Executor) Reset() {
	s.Head = 0
	s.Stack = s.Stack[:0]
	s.Err = nil
}

//go:inline
func (s *Executor) Active() bool {
	return s.Head < s.Emit.Len() && s.Err == nil
}

//go:inline
func (s *Executor) Op() byte {
	o := s.Emit[s.Head]
	s.Head++
	return o
}

//go:inline
func (s *Executor) Jmp(pc atf.PC) {
	s.Head = pc
}

//go:inline
func (s *Executor) PC() atf.PC {
	return s.Head
}

//go:inline
func (s *Executor) Fail(err error) {
	s.Err = err
}

//go:inline
func (s *Executor) Status() error {
	return s.Err
}
