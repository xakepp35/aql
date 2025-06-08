package fn

import (
	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

// Func is a direct lookup table from OpCode to handler func.
var Func = [256]vmi.Func{

	// flow

	op.Nop:  Nop,
	op.Call: Call,
	op.Over: Over,
	op.Loop: Loop,
	// op.Break: Break,
	op.Halt: Halt,

	// stack

	op.Nil:  Nil,
	op.Pop:  Pop,
	op.Dup:  Dup,
	op.Swap: Swap,
	op.Id:   Id,

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

func init() {
	for i, f := range Func {
		if f == nil {
			Func[i] = Unimplemented
		}
	}
}
