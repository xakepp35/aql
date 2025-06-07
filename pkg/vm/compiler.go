package vm

import (
	"encoding/binary"

	"github.com/xakepp35/aql/pkg/aqc"
	"github.com/xakepp35/aql/pkg/vm/fn"
	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

// emit is for compiling program and using to init it
// только на фазе инициализации, используется чтобы инициализировать стек перед запуском Run
type Compiler struct {
	prog  vmi.Program
	args  []int64    // аргументы (индексы, числа, смещения и т.п.)
	types []vmi.Type // type of each arg, detected at compile time
	data  []byte     // строковые/числовые литералы (индексируемые)
}

//go:inline
func NewCompiler() vmi.Compiler {
	return &Compiler{
		prog:  make(vmi.Program, 0),
		args:  make([]int64, 0),
		types: make([]vmi.Type, 0),
		data:  make([]byte, 0),
	}
}

//go:inline
func (e *Compiler) Compile(src []byte) error {
	return aqc.Compile(src, e)
}

//go:inline
func (e *Compiler) Init(this vmi.State) {
	for i, t := range e.types {
		arg := e.args[i]
		switch t {
		case vmi.TypeUint64:
			this.Push(int64(arg))
		case vmi.TypeBool:
			this.Push(arg != 0)
		case vmi.TypeString:
			s := ReadString(e.data, arg)
			this.Push(s)
		case vmi.TypeNull:
			this.Push(nil)
		}
	}
}

//go:inline
func (e *Compiler) Program() vmi.Program {
	return e.prog
}

//go:inline
func (e *Compiler) JIT() vmi.JIT {
	return vmi.JIT(fn.NewJIT(e.prog...))
}

//go:inline
func (c *Compiler) EmitInt(n int64) int {
	pos := len(c.args)
	c.args = append(c.args, (int64)(n))
	c.types = append(c.types, vmi.TypeUint64)
	return pos
}

//go:inline
func (c *Compiler) EmitBool(b bool) int {
	pos := len(c.args)
	var n int16
	if b {
		n = 1
	}
	c.args = append(c.args, (int64)(n))
	c.types = append(c.types, vmi.TypeBool)
	return pos
}

//go:inline
func (c *Compiler) EmitNull() int {
	pos := len(c.args)
	c.args = append(c.args, 0)
	c.types = append(c.types, vmi.TypeNull)
	return pos
}

//go:inline
func (c *Compiler) EmitString(s string) int {
	pos := len(c.args)
	length := uint32(len(s))
	off := int64(len(c.data))
	// пишем длину в 4 байтах (LittleEndian)
	var lenbuf [4]byte
	binary.LittleEndian.PutUint32(lenbuf[:], length)
	c.data = append(c.data, lenbuf[:]...)
	c.data = append(c.data, s...)
	c.args = append(c.args, off)
	c.types = append(c.types, vmi.TypeString)
	return pos
}

//go:inline
func (c *Compiler) EmitOps(ops ...op.Code) int {
	c.prog = append(c.prog, ops...)
	return len(c.prog)
}

//go:inline
func (c *Compiler) PatchInt(pos int, v int) {
	c.args[pos] = int64(v)
}

//go:inline
func ReadString(data []byte, off int64) string {
	if int(off)+4 > len(data) {
		return "" // защита от выхода за границы
	}
	length := binary.LittleEndian.Uint32(data[off : off+4])
	start := int(off + 4)
	end := start + int(length)
	if end > len(data) {
		return "" // недопустимый диапазон
	}
	return string(data[start:end])
}
