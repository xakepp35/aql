package fn

import (
	"github.com/xakepp35/aql/pkg/util"
	"github.com/xakepp35/aql/pkg/vmi"
)

func Count(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
		return
	}
	if c, ok := a[0].(int64); ok {
		this.Push(c + 1)
		return
	}
	this.SetErr(StackUnsupported(a...))
}

func Min(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
		return
	}

	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			if r < l {
				this.Push(r)
			} else {
				this.Push(l)
			}
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			if r < l {
				this.Push(r)
			} else {
				this.Push(l)
			}
			return
		}
	}
	this.SetErr(StackUnsupported(a...))
}

func Max(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			if r > l {
				this.Push(r)
			} else {
				this.Push(l)
			}
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			if r > l {
				this.Push(r)
			} else {
				this.Push(l)
			}
			return
		}
	}
	this.SetErr(StackUnsupported(a...))
}

// шаг на каждый элемент
func Avg(this vmi.VM) {
	a := this.Args(3)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
		return
	}

	sum, ok1 := util.ToFloat(a[0])
	cnt, ok2 := a[1].(int64)
	val, ok3 := util.ToFloat(a[2])
	if !ok1 || !ok2 || !ok3 {
		this.SetErr(StackUnsupported(a...))
		return
	}

	this.PushArgs(sum+val, cnt+1) // кладём обратно sum и cnt
}

func AvgFinal(this vmi.VM) {
	a := this.Args(2)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
		return
	}

	sum, ok1 := util.ToFloat(a[0])
	cnt, ok2 := a[1].(int64)
	if !ok1 || !ok2 || cnt == 0 {
		this.SetErr(StackUnsupported(a...))
		return
	}
	this.Push(sum / float64(cnt))
}
