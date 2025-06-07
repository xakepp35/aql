package vm

import (
	"context"
	"fmt"

	"github.com/xakepp35/aql/pkg/vm/fn"
	"github.com/xakepp35/aql/pkg/vmi"
)

type State struct {
	pc uint
	vmi.Stack
	vmi.Variables
	fns vmi.NamedFuncs
	ctx context.Context
	err error
}

//go:inline
func NewState() vmi.State {
	return &State{
		Stack:     NewStack(),
		Variables: NewVariables(),
		fns:       NewNamedFuncs(),
	}
}

//go:inline
func (s *State) Run(p vmi.Program) {
	s.pc = 0
	pLen := uint(len(p))
	for s.pc < pLen {
		o := p[s.pc]
		f := fn.Func[o]
		f(s)
		if s.err != nil {
			return
		}
		s.pc++
	}
}

//go:inline
func (s *State) Runf(fns vmi.JIT) {
	s.pc = 0
	pLen := uint(len(fns))
	for s.pc < pLen {
		f := fns[s.pc]
		f(s)
		if s.err != nil {
			return
		}
		s.pc++
	}
}

// slower option for debugging
//
//go:inline
func (s *State) Next(p vmi.Program) bool {
	pLen := uint(len(p))
	if s.pc >= pLen {
		return false
	}
	o := p[s.pc]
	f := fn.Func[o]
	f(s)
	if s.err != nil {
		return false
	}
	s.pc++
	return true
}

//go:inline
func (s *State) SetPC(pc uint) {
	s.pc = pc
}

//go:inline
func (s *State) AddPC(pc int) {
	s.pc += uint(pc)
}

//go:inline
func (s *State) PC() uint {
	return s.pc
}

//go:inline
func (s *State) Err() error {
	return s.err
}

//go:inline
func (s *State) SetErr(err error) {
	s.err = err
}

////////////////////////////////////////////////////////////////
// External func block

//go:inline
func NewNamedFuncs() vmi.NamedFuncs {
	return make(vmi.NamedFuncs)
}

//go:inline
func (s *State) Call(fnName string) {
	fn, ok := s.fns[fnName]
	if !ok {
		s.SetErr(fmt.Errorf("undefined function: %s", fnName))
		return
	}
	fn(s)
}

//go:inline
func (s *State) SetCall(fnName string, fn vmi.OpFunc) {
	s.fns[fnName] = fn
}

////////////////////////////////////////////////////////////////
// context for io/network calls

func (s *State) Context() context.Context {
	return s.ctx
}

func (s *State) SetContext(ctx context.Context) {
	s.ctx = ctx
}
