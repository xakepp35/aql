package vm

import "github.com/xakepp35/aql/pkg/vmi"

////////////////////////////////////////////////////////////////
// Stack manipulation

type Stack []any

//go:inline
func NewStack() vmi.Stacker {
	return make(Stack, 0, 16)
}

//go:inline
func (s Stack) Pushs(vals ...any) vmi.Stacker {
	return append(s, vals...)
}

//go:inline
func (s Stack) Push(v any) vmi.Stacker {
	return append(s, v)
}

//go:inline
func (s Stack) Pops(dst *[]any, count uint) vmi.Stacker {
	l := int(count)
	if l > len(s) {
		*dst = nil
		return s
	}
	n := len(s) - l
	*dst = s[:n]
	return s[n:]
}

//go:inline
func (s Stack) Pop(dst *any) vmi.Stacker {
	n := len(s)
	if n == 0 {
		return nil
	}
	n--
	*dst = s[n]
	return s[:n]
}

//go:inline
func (s Stack) Depth() int {
	return len(s)
}

//go:inline
func (s Stack) Dump() []any {
	return s
}
