package vm

import (
	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/vmi"
)

type VM struct {
	Executor
	Stream
	Variables
	Functions
	Context
}

//go:inline
func NewVM() vmi.VM {
	return &VM{
		Executor: Executor{
			Stack: make(Stack, 0, 16),
			Program: asf.Program{
				Emit: asf.Emitter{},
				Head: 0,
			},
			ErrStatus: ErrStatus{
				err: nil,
			},
		},
		Variables: make(Variables),
		Functions: make(Functions),
		Context:   Context{},
	}
}
