package vmo

// Not pops a boolean and pushes logical NOT. Or negates if its a number
//go:inline
func Not(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
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
		this.Fail(StackUnsupported(a...))
	}
}

// And logical AND.
//go:inline
func And(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch a1 := a[1].(type) {
	case bool:
		a0, ok := a[0].(bool)
		if !ok {
			this.Fail(StackUnsupported(a...))
			return
		}
		this.Push(a0 && a1)
	case int64:
		a0, ok := a[0].(int64)
		if ok {
			this.Fail(StackUnsupported(a...))
			return
		}
		this.Push(a0 & a1)
	default:
		this.Fail(StackUnsupported(a...))
	}
}

// Or
//go:inline
func Or(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch a1 := a[1].(type) {
	case bool:
		lb, ok := a[0].(bool)
		if !ok {
			this.Fail(StackUnsupported(a...))
			return
		}
		this.Push(lb || a1)
	case int64:
		lb, ok := a[0].(int64)
		if ok {
			this.Fail(StackUnsupported(a...))
			return
		}
		this.Push(lb | a1)
	default:
		this.Fail(StackUnsupported(a...))
	}
}

// Xor
//go:inline
func Xor(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch a1 := a[1].(type) {
	case bool:
		a0, ok := a[0].(bool)
		if !ok {
			this.Fail(StackUnsupported(a...))
			return
		}
		this.Push(a0 != a1)
	case int64:
		a0, ok := a[0].(int64)
		if ok {
			this.Fail(StackUnsupported(a...))
			return
		}
		this.Push((a0 | a1) & ^(a0 & a1))
	default:
		this.Fail(StackUnsupported(a...))
	}
}
