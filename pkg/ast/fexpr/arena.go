package fexpr

import (
	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/asf/atf"
)

type Arena []Compiler

//go:inline
func NewArena(capHint int) *Arena {
	res := make(Arena, 0, capHint)
	return &res
}

// allocate возвращает *сохранённый* экземпляр
//
//go:inline
func (a *Arena) Alloc(c Compiler) Compiler {
	*a = append(*a, c)
	return (*a)[len(*a)-1]
}

//go:inline
func (a *Arena) I64(v int64) Compiler {
	return a.Alloc(func(e *asf.Emitter) error {
		e.I64(v)
		return nil
	})
}

//go:inline
func (a *Arena) Binary(l, r Compiler, o op.Code) Compiler {
	return a.Alloc(func(e *asf.Emitter) error {
		if err := l(e); err != nil {
			return nil
		}
		if err := r(e); err != nil {
			return nil
		}
		e.U8(atf.U8(o))
		return nil
	})
}
