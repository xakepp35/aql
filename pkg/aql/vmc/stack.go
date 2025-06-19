package vmc

import "github.com/xakepp35/aql/pkg/vmi"

////////////////////////////////////////////////////////////////
// Stack manipulation

type Stack []any

//go:inline
func NewStack() Stack {
	s := make(Stack, 0, 16)
	return s
}

//go:inline
func (s *Stack) Open(pops, pushes int) []any {
	n := len(*s)
	diff := pushes - pops
	if diff > 0 {
		*s = append(*s, make([]any, diff)...)
	}
	return (*s)[n-pops:]
}

//go:inline
func (s *Stack) Close(pops, pushes int) {
	n := len(*s)
	diff := pushes - pops
	*s = (*s)[:n+diff]
}

//go:inline
func (s *Stack) Pushs(vals ...any) vmi.Stacker {
	*s = append(*s, vals...)
	return s
}

//go:inline
func (s *Stack) Push(v any) vmi.Stacker {
	*s = append(*s, v)
	return s
}

//go:inline
func (s *Stack) Pops(count uint) []any {
	l := int(count)
	r := *s
	if l > len(r) {
		return nil
	}
	n := len(r) - l
	*s = r[:n]
	return r[n:]
}

//go:inline
func (s *Stack) Pop() any {
	r := *s
	n := len(r)
	if n == 0 {
		return nil
	}
	n--
	*s = r[:n]
	return r[n]
}

//go:inline
func (s *Stack) Depth() int {
	return len(*s)
}

//go:inline
func (s *Stack) Dump() []any {
	return *s
}
