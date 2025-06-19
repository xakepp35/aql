package fexpr

import (
	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/asf/atf"
)

type Compiler func(e *asf.Emitter) error

func Nop(e *asf.Emitter) error {
	return nil
}

func Dup(e *asf.Emitter) error {
	e.U8(atf.U8(op.Dup))
	return nil
}

func Swap(e *asf.Emitter) error {
	e.U8(atf.U8(op.Swap))
	return nil
}

func This(e *asf.Emitter) error {
	e.U8(atf.U8(op.This))
	return nil
}

func Literal(name []byte) Compiler {
	return func(e *asf.Emitter) error {
		e.StringBytes(name)
		e.U8(atf.U8(op.Ident))
		return nil
	}
}

func Ident(name []byte) Compiler {
	return func(e *asf.Emitter) error {
		e.StringBytes(name)
		e.U8(atf.U8(op.Ident))
		return nil
	}
}

func Unary(r Compiler, o op.Code) Compiler {
	return func(e *asf.Emitter) error {
		if err := r(e); err != nil {
			return nil
		}
		e.U8(atf.U8(o))
		return nil
	}
}

func Binary(l, r Compiler, o op.Code) Compiler {
	return func(e *asf.Emitter) error {
		if err := l(e); err != nil {
			return nil
		}
		if err := r(e); err != nil {
			return nil
		}
		e.U8(atf.U8(o))
		return nil
	}
}

func Ternary(a, b, c Compiler, o op.Code) Compiler {
	return func(e *asf.Emitter) error {
		if err := a(e); err != nil {
			return nil
		}
		if err := b(e); err != nil {
			return nil
		}
		if err := c(e); err != nil {
			return nil
		}
		e.U8(atf.U8(o))
		return nil
	}
}

func Call(args []Compiler, name []byte) Compiler {
	return func(e *asf.Emitter) error {
		for _, a := range args {
			a(e)
		}
		e.StringBytes(name)
		e.U8(atf.U8(op.Call))
		return nil
	}
}

func Field(base Compiler, name []byte) Compiler {
	return func(e *asf.Emitter) error {
		base(e)
		e.StringBytes(name)
		e.U8(atf.U8(op.Field))
		return nil
	}
}

func Pipe(base Compiler, name []byte) Compiler {
	return func(e *asf.Emitter) error {
		base(e)
		e.StringBytes(name)
		e.U8(atf.U8(op.Field))
		return nil
	}
}

func Over(base Compiler, expr Compiler) Compiler {
	return func(e *asf.Emitter) error {
		base(e)
		expr(e)
		e.U8(atf.U8(op.Over))
		return nil
	}
}
