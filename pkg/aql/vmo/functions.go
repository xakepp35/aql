package vmo

// External function calls, can be managed by caller
//
//go:inline
func NewFunctions() Functions {
	return make(Functions)
}

type Functions map[string]Fn

func (f Functions) SetCalls(calls map[string]Fn) {
	for name, call := range calls {
		f[name] = call
	}
}

func (f Functions) SetCall(name string, call Fn) {
	f[name] = call
}

func (f Functions) GetCall(name string) Fn {
	return f[name]
}

func (f Functions) GetCalls() map[string]Fn {
	return f
}
