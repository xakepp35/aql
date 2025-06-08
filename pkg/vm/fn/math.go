package fn

import (
	"github.com/xakepp35/aql/pkg/vmi"
)

// Add pops two numbers, adds them, and pushes the sum.
func Add(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			this.Push(l + r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			this.Push(l + r)
			return
		}
	case string:
		if r, ok := a[1].(string); ok {
			this.Push(l + r)
			return
		}
	case []byte:
		if r, ok := a[1].([]byte); ok {
			this.Push(append(l, r...))
			return
		}
	case []any:
		if r, ok := a[1].([]any); ok {
			this.Push(append(l, r...))
			return
		}
	case map[string]any:
		if r, ok := a[1].(map[string]any); ok {
			res := make(map[string]any, len(l)+len(r))
			for k, v := range l {
				res[k] = v
			}
			for k, v := range r {
				res[k] = v
			}
			this.Push(res)
			return
		}
	case bool:
		if r, ok := a[1].(bool); ok {
			this.Push(l || r)
			return
		}
	}
	this.SetErr(StackUnsupported(a...))
}

// Sub pops two numbers, subtracts them, and pushes the difference.
func Sub(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
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
	this.SetErr(StackUnsupported(a...))
}

// Mul pops two numbers, multiplies them, and pushes the product.
func Mul(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			this.Push(l * r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			this.Push(l * r)
			return
		}
	}
	this.SetErr(StackUnsupported(a...))
}

// Div pops two numbers, divides them, and pushes the quotient.
func Div(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			if r == 0 {
				this.SetErr(ErrDivisionByZero)
				return
			}
			this.Push(l / r)
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			if r == 0 {
				this.SetErr(ErrDivisionByZero)
				return
			}
			this.Push(l / r)
			return
		}
	}
	this.SetErr(StackUnsupported(a...))
}

// Mod pops two numbers, computes modulus, and pushes the result.
func Mod(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(StackUnderflow(this.Dump()...))
		return
	}
	rf, ok1 := a[1].(int64)
	lf, ok2 := a[0].(int64)
	if !ok1 || !ok2 {
		this.SetErr(StackUnsupported(a...))
		return
	}
	if rf == 0 {
		this.SetErr(ErrModuloByZero)
		return
	}
	this.Push(int64(int64(lf) % int64(rf)))
}
