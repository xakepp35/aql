package vmo

import (
	"github.com/xakepp35/aql/pkg/util"
	"github.com/xakepp35/aql/pkg/vmi"
)

func Count(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(vmi.ErrStackUnderflow)
		return
	}
	if c, ok := a[0].(int64); ok {
		this.Push(c + 1)
		return
	}
	this.Fail(StackUnsupported(a...))
}

func Min(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(vmi.ErrStackUnderflow)
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
	this.Fail(StackUnsupported(a...))
}

func Max(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(vmi.ErrStackUnderflow)
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
	this.Fail(StackUnsupported(a...))
}

// шаг на каждый элемент
func Avg(this *VM) {
	a := this.Pops(3)
	if a == nil {
		this.Fail(vmi.ErrStackUnderflow)
		return
	}

	sum, ok1 := util.ToFloat(a[0])
	cnt, ok2 := a[1].(int64)
	val, ok3 := util.ToFloat(a[2])
	if !ok1 || !ok2 || !ok3 {
		this.Fail(StackUnsupported(a...))
		return
	}

	this.Pushs(sum+val, cnt+1) // кладём обратно sum и cnt
}

func AvgFinal(this *VM) {
	a := this.Pops(2)
	if a == nil {
		this.Fail(vmi.ErrStackUnderflow)
		return
	}

	sum, ok1 := util.ToFloat(a[0])
	cnt, ok2 := a[1].(int64)
	if !ok1 || !ok2 || cnt == 0 {
		this.Fail(StackUnsupported(a...))
		return
	}
	this.Push(sum / float64(cnt))
}
