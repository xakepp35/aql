package fn

import (
	"github.com/xakepp35/aql/pkg/vmi"
)

// FnNop does nothing.
func Nop(this vmi.State) {}

// Halt stops operation
func Halt(this vmi.State) {
	this.SetErr(vmi.ErrHalted)
}

// Unimplemented
func Unimplemented(this vmi.State) {
	this.SetErr(vmi.ErrUnimplemented)
}

// Over
func Over(this vmi.State) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
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
func Loop(this vmi.State) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
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
