package fn

import (
	"github.com/xakepp35/aql/pkg/util"
	"github.com/xakepp35/aql/pkg/vmi"
)

func Count(st vmi.State) {
	a := st.Args(2)
	if a == nil {
		st.SetErr(ErrStackUnderflow)
		return
	}
	if c, ok := a[0].(int64); ok {
		st.Push(c + 1)
		return
	}
	st.SetErr(StackUnsupported(a...))
}

func Min(st vmi.State) {
	a := st.Args(2)
	if a == nil {
		st.SetErr(ErrStackUnderflow)
		return
	}

	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			if r < l {
				st.Push(r)
			} else {
				st.Push(l)
			}
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			if r < l {
				st.Push(r)
			} else {
				st.Push(l)
			}
			return
		}
	}
	st.SetErr(StackUnsupported(a...))
}

func Max(st vmi.State) {
	a := st.Args(2)
	if a == nil {
		st.SetErr(ErrStackUnderflow)
		return
	}
	switch l := a[0].(type) {
	case int64:
		if r, ok := a[1].(int64); ok {
			if r > l {
				st.Push(r)
			} else {
				st.Push(l)
			}
			return
		}
	case float64:
		if r, ok := a[1].(float64); ok {
			if r > l {
				st.Push(r)
			} else {
				st.Push(l)
			}
			return
		}
	}
	st.SetErr(StackUnsupported(a...))
}

// шаг на каждый элемент
func AvgStep(st vmi.State) {
	a := st.Args(3)
	if a == nil {
		st.SetErr(ErrStackUnderflow)
		return
	}

	sum, ok1 := util.ToFloat(a[0])
	cnt, ok2 := a[1].(int64)
	val, ok3 := util.ToFloat(a[2])
	if !ok1 || !ok2 || !ok3 {
		st.SetErr(StackUnsupported(a...))
		return
	}

	st.PushArgs(sum+val, cnt+1) // кладём обратно sum и cnt
}

func AvgFinal(st vmi.State) {
	a := st.Args(2)
	if a == nil {
		st.SetErr(ErrStackUnderflow)
		return
	}

	sum, ok1 := util.ToFloat(a[0])
	cnt, ok2 := a[1].(int64)
	if !ok1 || !ok2 || cnt == 0 {
		st.SetErr(StackUnsupported(a...))
		return
	}
	st.Push(sum / float64(cnt))
}
