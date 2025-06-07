package fn

import (
	"time"

	"github.com/xakepp35/aql/pkg/vmi"
)

//go:inline
func PushNil(this vmi.State) {
	this.Push(nil)
}

//go:inline
func PushNow(this vmi.State) {
	this.Push(time.Now().UTC())
}

//go:inline
func Pop(this vmi.State) {
	_ = this.Pop()
}

//go:inline
func Dup(this vmi.State) {
	a := this.Pop()
	this.Push(a)
	this.Push(a)
}

//go:inline
func Swap(this vmi.State) {
	a := this.Pop()
	b := this.Pop()
	this.Push(a)
	this.Push(b)
}
