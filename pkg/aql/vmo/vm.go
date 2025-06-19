package vmo

import (
	"github.com/xakepp35/aql/pkg/aql/vmc"
)

type Tracer = func(*VM)

type This = any
type VM struct {
	vmc.Executor
	vmc.Variables
	vmc.SendStream
	vmc.RecvStream
	*Table
	Functions
	This
	Tracer
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
	if s.Tracer != nil {
		s.Tracer(s)
	}
}
