package vmo

import (
	"github.com/xakepp35/aql/pkg/vmi"
)

// Add pops two numbers, adds them, and pushes the sum.
func Add(this *VM) {
	a := this.Stack.Open(2, 1)
	defer this.Stack.Close(2, 1)
	// if a == nil {
	// 	this.Fail(StackUnderflow(this.Dump()...))
	// 	return
	// }
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			a[0] = l + r
			//this.Push(l + r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			a[0] = l + r
			// this.Push(l + r)
			return
		}
	case string:
		if r, ok := a[1].(string); ok {
			a[0] = l + r
			// this.Push(l + r)
			return
		}
	case []byte:
		if r, ok := a[1].([]byte); ok {
			a[0] = append(l, r...)
			// this.Push(append(l, r...))
			return
		}
	case []any:
		if r, ok := a[1].([]any); ok {
			a[0] = append(l, r...)
			// this.Push(append(l, r...))
			return
		}
		// case map[string]any:
		// 	if r, ok := a[1].(map[string]any); ok {
		// 		res := make(map[string]any, len(l)+len(r))
		// 		for k, v := range l {
		// 			res[k] = v
		// 		}
		// 		for k, v := range r {
		// 			res[k] = v
		// 		}
		// 		this.Push(res)
		// 		return
		// 	}
		// case bool:
		// 	if r, ok := a[1].(bool); ok {
		// 		this.Push(l || r)
		// 		return
		// 	}
	}
	this.Fail(StackUnsupported(a...))
}

// func Add(this *VM) {
// 	a, b := this.Open(2), this.Open(1)
// 	defer this.Close(2, 1)
// 	switch a.Kind {
// 	case vm.KindInt:
// 		a.SetI64(a.I64() + b.I64())
// 	case vm.KindFloat:
// 		a.SetF64(a.F64() + b.F64())
// 	default:
// 		// this.Fail(StackUnsupported(a...))
// 	}
// }

// Sub pops two numbers, subtracts them, and pushes the difference.
func Sub(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			this.Push(l - r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			this.Push(l - r)
			return
		}
	}
	this.Fail(StackUnsupported(a...))
}

// Mul pops two numbers, multiplies them, and pushes the product.
func Mul(this *VM) {
	// a := this.Pops(2)
	// if a == nil {
	// 	this.Fail(StackUnderflow(this.Dump()...))
	// 	return
	// }
	a := this.Stack.Open(2, 1)
	defer this.Stack.Close(2, 1)
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			// this.Stack = append(this.Stack, l*r)
			a[0] = l * r
			// this.Push(l * r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			a[0] = l * r
			// this.Push(l * r)
			return
		}
	}
	this.Fail(StackUnsupported(a...))
}

// func Mul(this *VM) {
// 	a, b := this.Open(2), this.Open(1)
// 	defer this.Close(2, 1)
// 	switch a.Kind {
// 	case vm.KindInt:
// 		a.SetI64(a.I64() * b.I64())
// 	case vm.KindFloat:
// 		a.SetF64(a.F64() * b.F64())
// 	default:
// 		// this.Fail(StackUnsupported(args...))
// 	}
// }

// Div pops two numbers, divides them, and pushes the quotient.
func Div(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			if r == 0 {
				this.Fail(vmi.ErrDivisionByZero)
				return
			}
			this.Push(l / r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			if r == 0 {
				this.Fail(vmi.ErrDivisionByZero)
				return
			}
			this.Push(l / r)
			return
		}
	}
	this.Fail(StackUnsupported(a...))
}

// Mod pops two numbers, computes modulus, and pushes the result.
func Mod(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}
	rf, ok1 := a[1].(int64)
	lf, ok2 := a[0].(int64)
	if !ok1 || !ok2 {
		this.Fail(StackUnsupported(a...))
		return
	}
	if rf == 0 {
		this.Fail(vmi.ErrModuloByZero)
		return
	}
	this.Push(int64(int64(lf) % int64(rf)))
}
