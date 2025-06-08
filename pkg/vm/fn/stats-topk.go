package fn

import (
	"container/heap"
	"sort"

	"github.com/xakepp35/aql/pkg/vmi"
)

/* ---------- Top-K ----------------------------------------------- */

type kv struct {
	key any
	n   int64
}

type minHeap []kv

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].n < h[j].n } // min-heap
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(kv)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type topK struct {
	K     int
	Table map[any]*kv
	H     minHeap
}

/* step-opcode, генерируется curry-функцией */

func NewTopKOp(k int) func(vmi.VM) {
	return func(this vmi.VM) {
		a := this.Args(2)
		if a == nil {
			this.SetErr(ErrStackUnderflow)
			return
		}

		acc, ok := a[0].(*topK)
		if !ok {
			this.SetErr(StackUnsupported(a[0]))
			return
		}
		key := a[1]
		e := acc.Table[key]
		if e == nil {
			e = &kv{key: key}
			acc.Table[key] = e
		}
		e.n++

		if len(acc.H) < acc.K {
			heap.Push(&acc.H, *e) // alloc  — объективно нужно
		} else if acc.H[0].n < e.n {
			acc.H[0] = *e
			heap.Fix(&acc.H, 0)
		}
		this.Push(acc) // вернуть обновлённый аккумулятор
	}
}

// финальный op
func TopKFinal(st vmi.VM) {
	a := st.Args(1)
	if a == nil {
		st.SetErr(ErrStackUnderflow)
		return
	}

	if acc, ok := a[0].(*topK); ok {
		out := make([]kv, len(acc.H))
		copy(out, acc.H)
		sort.Slice(out, func(i, j int) bool { return out[i].n > out[j].n })
		st.Push(out)
		return
	}
	st.SetErr(StackUnsupported(a[0]))
}
