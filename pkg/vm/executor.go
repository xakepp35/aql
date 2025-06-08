package vm

import (
	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/vm/fn"
	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Executor struct {
	asf.Program
	Stack
	ErrStatus
}

// Run main program entrypoint
//
//go:inline
func (s *Executor) Run(this vmi.VM) {
	s.Head = 0
	for {
		s.step(this)
	}
}

// Next is a slower option for debugging
//
//go:inline
func (s *Executor) Next(this vmi.VM) bool {
	s.step(this)
	if s.err != nil {
		return false
	}
	return true
}

//go:inline
func (s *Executor) step(this vmi.VM) {
	dLen := atf.PC(len(s.Emit))
	if s.Head >= dLen {
		s.err = vmi.ErrFinished
		return
	}
	o := op.Code(s.Emit[s.Head])
	s.Head++
	v, err := s.Load(asf.Type(o))
	if err != nil {
		f := fn.Func[o]
		f(this)
		return
	}
	s.Stack = append(s.Stack, v)
	return
}

func (s *Executor) exec()

//go:inline
func (s *Executor) Jmp(pc atf.PC) {
	s.Head = pc
}

//go:inline
func (s *Executor) PC() atf.PC {
	return s.Head
}
