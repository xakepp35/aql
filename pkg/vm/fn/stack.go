package fn

import (
	"github.com/xakepp35/aql/pkg/vmi"
)

//go:inline
func Id(this vmi.VM) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	name, ok := a[0].(string)
	if !ok {
		this.SetErr(StackUnsupported(a...))
		return
	}
	if f := Builtins[name]; f != nil {
		f(this)
		return
	}
	val, ok := this.Get(name)
	if !ok {
		this.SetErr(VariableUndefined(name))
		return
	}
	this.Push(val)
}

//go:inline
func Pop(this vmi.VM) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
}

//go:inline
func Dup(this vmi.VM) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	this.PushArgs(a[0], a[0])
}

//go:inline
func Swap(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	this.PushArgs(this.Pop(), this.Pop())
}
