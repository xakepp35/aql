package fn

import (
	"github.com/xakepp35/aql/pkg/vmi"
)

// FnNop does nothing.
func Nop(this vmi.VM) {}

// Halt stops operation
func Halt(this vmi.VM) {
	this.SetErr(vmi.ErrHalted)
}

// Unimplemented
func Unimplemented(this vmi.VM) {
	this.SetErr(vmi.ErrUnimplemented)
}

func Call(this vmi.VM) {
	// stack-layout: ... <name string> <argc int64> <args ...any>
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	name, ok1 := a[0].(string)
	argc, ok2 := a[1].(int64)
	if !ok1 || !ok2 {
		this.SetErr(StackUnsupported(a...))
		return
	}
	this.Push(argc)
	if f := Builtins[name]; f != nil {
		f(this)
		return
	}
	if f := this.Call(name); f != nil {
		f(this)
		return
	}
	this.Push(name)
	this.SetErr(FunctionUndefined(name))
}

// Over
func Over(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	endPC, ok := a[1].(uint)
	if !ok {
		this.SetErr(ErrStackMissingPC)
		return
	}
	it, ok := a[0].(vmi.Iterator)
	if !ok {
		this.SetErr(ErrStackUniterable)
	}
	if !it.Next() {
		this.SetPC(endPC)
		return
	}
	this.PushArgs(
		this.PC()-1, // адрес собственного OpLoop
		it,
		it.Item(),
	)
}

// Loop
func Loop(this vmi.VM) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	loopPC, ok := a[0].(uint)
	if !ok {
		this.SetErr(ErrStackMissingPC)
		return
	}
	this.SetPC(loopPC)
	this.PushArgs(a...)
}
