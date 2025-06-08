package vm

import "github.com/xakepp35/aql/pkg/vmi"

// External function calls, can be managed by caller
//
//go:inline
func NewFunctions() vmi.Functioner {
	return make(Functions)
}

type Functions map[string]vmi.Func

func (f Functions) SetCalls(calls map[string]vmi.Func) {
	for name, call := range calls {
		f[name] = call
	}
}

func (f Functions) SetCall(name string, call vmi.Func) {
	f[name] = call
}

func (f Functions) GetCall(name string) vmi.Func {
	return f[name]
}

func (f Functions) GetCalls() map[string]vmi.Func {
	return f
}
