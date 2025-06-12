package vmo

import (
	"github.com/xakepp35/aql/pkg/vm/vmc"
)

type VM struct {
	vmc.Executor
	vmc.Variables
	vmc.Stream
	vmc.Context
	*Table
	Functions
	UserData any
}

// Run main program entrypoint
//
//go:inline
func (s *VM) Run() *VM {
	s.Reset()
	for s.Active() {
		s.Next()
	}
	return s
}

// Next is a slower option for debugging
//
//go:inline
func (s *VM) Next() {
	o := s.Op()
	f := s.Table[o]
	f(s)
}
