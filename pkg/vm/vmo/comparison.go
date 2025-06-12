package vmo

import (
	"reflect"
)

// FnEq pops two values, compares equality, and pushes result.
//
//go:inline
func Eq(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	this.Push(reflect.DeepEqual(a[0], a[1]))
}

// FnNeq pops two values, compares inequality, and pushes result.
//
//go:inline
func Neq(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	this.Push(!reflect.DeepEqual(a[0], a[1]))
}

// FnLt pops two numbers, compares <, and pushes result.
//
//go:inline
func Lt(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			this.Push(l < r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			this.Push(l < r)
			return
		}
	case string:
		if r, ok := a[1].(string); ok {
			this.Push(l < r)
			return
		}
	}
	this.Fail(StackUnsupported(a...))
}

// FnLe pops two numbers, compares <=, and pushes result.
//
//go:inline
func Le(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			this.Push(l <= r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			this.Push(l <= r)
			return
		}
	case string:
		if r, ok := a[1].(string); ok {
			this.Push(l <= r)
			return
		}
	}
	this.Fail(StackUnsupported(a...))
}

// FnGt pops two numbers, compares >, and pushes result.
//
//go:inline
func Gt(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			this.Push(l > r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			this.Push(l > r)
			return
		}
	case string:
		if r, ok := a[1].(string); ok {
			this.Push(l > r)
			return
		}
	}
	this.Fail(StackUnsupported(a...))
}

// FnGe pops two numbers, compares >=, and pushes result.
//
//go:inline
func Ge(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			this.Push(l >= r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			this.Push(l >= r)
			return
		}
	case string:
		if r, ok := a[1].(string); ok {
			this.Push(l >= r)
			return
		}
	}
	this.Fail(StackUnsupported(a...))
}
