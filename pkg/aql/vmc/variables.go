package vmc

import "github.com/xakepp35/aql/pkg/vmi"

////////////////////////////////////////////////////////////////
// variables hashmap - dictionary for literals

type Variables map[string]any

//go:inline
func NewVariables() vmi.Variabler {
	return make(Variables)
}

//go:inline
func (s Variables) SetVars(dict map[string]any) {
	for k, v := range dict {
		s[k] = v
	}
}

//go:inline
func (s Variables) Vars() map[string]any {
	return s
}

//go:inline
func (s Variables) Set(k string, v any) {
	s[k] = v
}

//go:inline
func (s Variables) Get(k string) (any, bool) {
	v, ok := s[k]
	return v, ok
}

//go:inline
func (s Variables) Del(k string) {
	delete(s, k)
}

//go:inline
func (s Variables) Clear() {
	for k := range s {
		delete(s, k)
	}
}
