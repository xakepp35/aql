package vmo

import (
	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf"
)

type Fn = func(this *VM)
type Table = [256]func(this *VM)

var Default Table

func init() {
	InitTable(&Default)
}

//go:inline
func InitTable(ops *[256]Fn) {
	for i := range *ops {
		(*ops)[i] = Unimplemented
	}
	*ops = [256]Fn{
		// loaders
		asf.Nop:        Nop,
		asf.Nil:        Nop,
		asf.U8:         Nop,
		asf.U16:        Nop,
		asf.U32:        Nop,
		asf.U64:        Nop,
		asf.U128:       Nop,
		asf.U256:       Nop,
		asf.I8:         Nop,
		asf.I16:        Nop,
		asf.I32:        Nop,
		asf.I64:        LdI64,
		asf.I128:       Nop,
		asf.I256:       Nop,
		asf.F32:        Nop,
		asf.F64:        LdF64,
		asf.Complex64:  Nop,
		asf.Complex128: Nop,
		asf.PC:         Nop,
		asf.SP:         Nop,
		asf.Time:       Nop,
		asf.Dur:        Nop,
		asf.Bytes:      Nop,
		asf.String:     LdString,
		asf.Error:      Nop,
		asf.IBig:       Nop,
		asf.FBig:       Nop,

		// flow
		op.Call:  Call,
		op.Over:  Over,
		op.Loop:  Loop,
		op.Break: Unimplemented,
		op.Halt:  Halt,

		// stack

		op.Pop:   Pop,
		op.Dup:   Dup,
		op.Swap:  Swap,
		op.Ident: Ident,

		// logic & math

		op.Not: Not,
		op.And: And,
		op.Or:  Or,
		op.Xor: Xor,
		// op.Shl: Shl,
		// op.Shr: Shr,
		op.Add: Add,
		op.Sub: Sub,
		op.Mul: Mul,
		op.Div: Div,
		op.Mod: Mod,

		// comparison

		op.Eq:  Eq,
		op.Neq: Neq,
		op.Lt:  Lt,
		op.Le:  Le,
		op.Gt:  Gt,
		op.Ge:  Ge,

		// data

		op.Pipe: Nop,
		// op.Index1: Index1,
		// op.Index2: Index2,
		// op.Field:  Field,
	}
}
