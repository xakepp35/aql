package vmo

// External function calls, can be managed by caller
//
//go:inline
func NewFunctions() Functions {
	return make(Functions)
}

type Functions map[string]Func

func (f Functions) SetCalls(calls map[string]Func) {
	for name, call := range calls {
		f[name] = call
	}
}

func (f Functions) SetCall(name string, call Func) {
	f[name] = call
}

func (f Functions) GetCall(name string) Func {
	return f[name]
}

func (f Functions) GetCalls() map[string]Func {
	return f
}
