package aql

import (
	"context"

	"github.com/xakepp35/aql/pkg/aql/vmc"
	"github.com/xakepp35/aql/pkg/aql/vmo"
	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/ast"
)

type VM = vmo.VM

func Run(src string) *VM {
	return NewSrc([]byte(src)).Run()
}

func New(ctx context.Context, cancel context.CancelFunc) *VM {
	return &VM{
		Executor:   vmc.New(ctx, cancel),
		Variables:  make(vmc.Variables),
		SendStream: make(vmc.SendStream),
		RecvStream: make(vmc.RecvStream),
		Functions:  vmo.Builtins,
		Table:      &vmo.Default,
	}
}

//go:inline
func NewSrc(src []byte) *VM {
	res := New(context.Background(), nil)
	res.Emit = MustCompile(src)
	return res
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
