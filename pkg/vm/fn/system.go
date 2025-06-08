package fn

import (
	"time"

	"github.com/xakepp35/aql/pkg/vmi"
)

//go:inline
func Nil(this vmi.VM) {
	this.Push(nil)
}

//go:inline
func Now(this vmi.VM) {
	this.Push(time.Now().UTC())
}
