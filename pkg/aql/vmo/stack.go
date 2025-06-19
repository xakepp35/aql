package vmo

//go:inline
func Ident(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	name, ok := a[0].(string)
	if !ok {
		this.Fail(StackUnsupported(a...))
		return
	}
	if val, ok := this.Get(name); ok {
		this.Push(val)
		return
	}
	if f := this.GetCall(name); f != nil {
		f(this)
		return
	}
	this.Fail(IdentifierUndefined(name))
}

//go:inline
func Pop(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
}

//go:inline
func Dup(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	this.Pushs(a[0], a[0])
}

//go:inline
func Swap(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	this.Pushs(this.Pop(), this.Pop())
}
