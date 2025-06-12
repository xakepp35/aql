package vm

import (
	"context"

	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/ast"
	"github.com/xakepp35/aql/pkg/vm/vmc"
	"github.com/xakepp35/aql/pkg/vm/vmo"
)

type VM = vmo.VM

func Run(src string) *VM {
	return NewSrc([]byte(src)).Run()
}

func New() *VM {
	return &VM{
		Executor: vmc.Executor{
			Stack: make(vmc.Stack, 0, 16),
			Program: asf.Program{
				Emit: asf.Emitter{},
				Head: 0,
			},
			Err: nil,
		},
		Variables: make(vmc.Variables),
		Stream:    make(vmc.Stream),
		Functions: vmo.Builtins,
		Table:     &vmo.Default,
		Context: vmc.Context{
			Context:    context.Background(),
			CancelFunc: func() {},
		},
	}
}

//go:inline
func NewSrc(src []byte) *VM {
	return &VM{
		Executor: vmc.Executor{
			Stack: make(vmc.Stack, 0, 16),
			Program: asf.Program{
				Emit: MustCompile(src),
				Head: 0,
			},
			Err: nil,
		},
		Variables: make(vmc.Variables),
		Stream:    make(vmc.Stream),
		Functions: vmo.Builtins,
		Table:     &vmo.Default,
		Context: vmc.Context{
			Context:    context.Background(),
			CancelFunc: func() {},
		},
	}
}

func MustCompile(src []byte) asf.Emitter {
	res, err := Compile(src)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func Compile(src []byte) (asf.Emitter, error) {
	a, err := ast.Parse(src)
	if err != nil {
		return nil, err
	}
	res := make(asf.Emitter, 0, 256)
	a.P0(&res)
	a.P1(&res)
	a.P2(&res)
	return res, nil
}
