package vm

import "github.com/xakepp35/aql/pkg/vmi"

////////////////////////////////////////////////////////////////
// Stack manipulation

type Stack []any

//go:inline
func NewStack() vmi.Stack {
	res := make(Stack, 0, 16)
	return &res
}

func (s *Stack) PushArgs(vals ...any) {
	for _, v := range vals {
		s.Push(v)
	}
}

func (s *Stack) Args(n uint) []any {
	r, l := *s, int(n)
	if l < len(r) {
		return nil
	}
	pos := len(r) - l
	*s = r[:pos]
	return r[pos:]
}

// Push кладёт значение в стек
//
//go:inline
func (s *Stack) Push(v any) {
	*s = append(*s, v)
}

// Pop снимает и возвращает верхнее значение стека.
// Если стек пуст, возвращает nil.
//
//go:inline
func (s *Stack) Pop() any {
	b := *s
	n := len(b)
	if n == 0 {
		return nil
	}
	v := b[n-1]
	*s = b[:n-1]
	return v
}

//go:inline
func (s *Stack) Len() int {
	return len(*s)
}
