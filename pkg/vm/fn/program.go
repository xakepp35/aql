package fn

import (
	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type JIT vmi.JIT

//go:inline
func NewJIT(prog ...op.Code) JIT {
	res := make(JIT, len(prog))
	for i := range prog {
		res[i] = Func[prog[i]]
	}
	return res
}

//go:inline
func (f JIT) Run(this vmi.VM) {
	this.Runf(vmi.JIT(f))
}

//go:inline
type Program vmi.ByteCode

//go:inline
func (p Program) Run(this vmi.VM) {
	this.Run(vmi.ByteCode(p))
}
