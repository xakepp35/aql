package fn

import (
	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

// Func is a direct lookup table from OpCode to handler func.
var Func = [256]vmi.OpFunc{
	op.Nop:     Nop,
	op.Add:     Add,
	op.Sub:     Sub,
	op.Mul:     Mul,
	op.Div:     Div,
	op.Mod:     Mod,
	op.And:     And,
	op.Or:      Or,
	op.Not:     Not,
	op.Eq:      Eq,
	op.Neq:     Neq,
	op.Lt:      Lt,
	op.Le:      Le,
	op.Gt:      Gt,
	op.Ge:      Ge,
	op.PushNil: PushNil,
	op.Halt:    Halt,
}

func init() {
	for i, f := range Func {
		if f == nil {
			Func[i] = Unimplemented
		}
	}
}
