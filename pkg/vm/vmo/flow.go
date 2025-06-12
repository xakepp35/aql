package vmo

import (
	"github.com/xakepp35/aql/pkg/vmi"
)

// FnNop does nothing.
func Nop(this *VM) {}

// Halt stops operation
func Halt(this *VM) {
	this.Fail(vmi.ErrHalted)
}

// Unimplemented
func Unimplemented(this *VM) {
	this.Fail(vmi.ErrUnimplemented)
}

func Call(this *VM) {
	Ident(this)
	// stack-layout: ... <name string> <argc int64> <args ...any>
	// a := this.Pops(1)
	// if a == nil {
	// 	this.Fail(StackUnderflow(this.Dump()...))
	// 	return
	// }
	// name, ok1 := a[0].(string)
	// if !ok1 {
	// 	this.Fail(StackUnsupported(a...))
	// 	return
	// }
	// if f := Builtins[name]; f != nil {
	// 	f(this)
	// 	return
	// }
	// if f := this.GetCall(name); f != nil {
	// 	f(this)
	// 	return
	// }
	// this.Push(name)
	// this.Fail(FunctionUndefined(name))
}

// Over
func Over(this *VM) {
	// a := this.Pops(2)
	// if a == nil {
	// 	this.Fail(StackUnderflow(this.Dump()...))
	// 	return
	// }
	// endPC, ok := a[1].(uint)
	// if !ok {
	// 	this.Fail(vmi.ErrStackMissingPC)
	// 	return
	// }
	// it, ok := a[0].(vmi.Iterator)
	// if !ok {
	// 	this.Fail(vmi.ErrStackUniterable)
	// }
	// if !it.Next() {
	// 	this.SetPC(endPC)
	// 	return
	// }
	// this.Pushs(
	// 	this.PC()-1, // адрес собственного OpLoop
	// 	it,
	// 	it.Item(),
	// )
}

// Loop
func Loop(this *VM) {
	// a := this.Pops(1)
	// if a == nil {
	// 	this.Fail(StackUnderflow(this.Dump()...))
	// 	return
	// }
	// loopPC, ok := a[0].(uint)
	// if !ok {
	// 	this.Fail(vmi.ErrStackMissingPC)
	// 	return
	// }
	// this.SetPC(loopPC)
	// this.Pushs(a...)
}
