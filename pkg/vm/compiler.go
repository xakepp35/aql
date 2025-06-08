package vm

import (
	"encoding/binary"

	"github.com/xakepp35/aql/pkg/aqc"
	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

// emit is for compiling program and using to init it
// только на фазе инициализации, используется чтобы инициализировать стек перед запуском Run
type Programmer struct {
	asf.Emitter
	vmi.AST
}

//go:inline
func NewProgrammer() vmi.Programmer {
	return &Programmer{}
}

//go:inline
func (e *Programmer) Compile(src []byte) error {
	return aqc.Compile(src, e)
}

//go:inline
func (c *Programmer) String(s string) int {
	pos := len(c.args)
	off := int64(len(c.data))
	c.data = AppendString(c.data, []byte(s))
	c.args = append(c.args, off)
	c.types = append(c.types, vmi.TypeString)
	return pos
}
