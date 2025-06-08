package fn

import (
	"github.com/xakepp35/aql/pkg/vmi"
)

// Not pops a boolean and pushes logical NOT. Or negates if its a number
func Not(this vmi.VM) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	switch x := a[0].(type) {
	case bool:
		this.Push(!x)
	case int64:
		this.Push(-x)
	case float64:
		this.Push(-x)
	default:
		this.SetErr(StackUnsupported(a...))
	}
}

// And logical AND.
func And(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	switch a1 := a[1].(type) {
	case bool:
		a0, ok := a[0].(bool)
		if !ok {
			this.SetErr(StackUnsupported(a...))
			return
		}
		this.Push(a0 && a1)
	case int64:
		a0, ok := a[0].(int64)
		if ok {
			this.SetErr(StackUnsupported(a...))
			return
		}
		this.Push(a0 & a1)
	default:
		this.SetErr(StackUnsupported(a...))
	}
}

// Or
func Or(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	switch a1 := a[1].(type) {
	case bool:
		lb, ok := a[0].(bool)
		if !ok {
			this.SetErr(StackUnsupported(a...))
			return
		}
		this.Push(lb || a1)
	case int64:
		lb, ok := a[0].(int64)
		if ok {
			this.SetErr(StackUnsupported(a...))
			return
		}
		this.Push(lb | a1)
	default:
		this.SetErr(StackUnsupported(a...))
	}
}

// Xor
func Xor(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	switch a1 := a[1].(type) {
	case bool:
		a0, ok := a[0].(bool)
		if !ok {
			this.SetErr(StackUnsupported(a...))
			return
		}
		this.Push(a0 != a1)
	case int64:
		a0, ok := a[0].(int64)
		if ok {
			this.SetErr(StackUnsupported(a...))
			return
		}
		this.Push((a0 | a1) & ^(a0 & a1))
	default:
		this.SetErr(StackUnsupported(a...))
	}
}
