// pkg/asf/context.go
package asf

import (
	"encoding/binary"
	"errors"
	"math"
	"math/big"
	"time"

	"github.com/xakepp35/aql/pkg/asf/atf"
)

type Program struct {
	Emit Emitter
	Head atf.PC
}

//go:inline
func NewProgram(data Emitter, head atf.PC) *Program {
	return &Program{
		Emit: data,
		Head: head,
	}
}

func (p *Program) Load(t Type) (any, error) {
	if t > Type_Count { // VM-opcode
		return nil, ErrInvalid
	}
	switch Type(t) {
	case Nil:
		return nil, nil
	case U8:
		return p.U8(), nil
	case U16:
		return p.U16(), nil
	case U32:
		return p.U32(), nil
	case U64:
		return p.U64(), nil
	case U128:
		return p.U128(), nil
	case U256:
		return p.U256(), nil
	case I8:
		return p.I8(), nil
	case I16:
		return p.I16(), nil
	case I32:
		return p.I32(), nil
	case I64:
		return p.I64(), nil
	case I128:
		return p.I128(), nil
	case I256:
		return p.I256(), nil
	case F32:
		return p.F32(), nil
	case F64:
		return p.F64(), nil
	case Complex64:
		return p.Complex64(), nil
	case Complex128:
		return p.Complex128(), nil
	case PC:
		return p.PC(), nil
	case SP:
		return p.SP(), nil
	case Len:
		return p.Len(), nil
	case Time:
		return p.Time(), nil
	case Dur:
		return p.Dur(), nil
	case Bytes:
		return p.Bytes(), nil
	case String:
		return p.String(), nil
	case Error:
		return p.Error(), nil
	case IBig:
		return p.IBig(), nil
	case FBig:
		return p.FBig(), nil
	default:
		return nil, ErrInvalid
	}
}

//go:inline
func (c *Program) Type() Type {
	t := Type(c.Emit[c.Head])
	c.Head++
	return t
}

//go:inline
func (c *Program) Nil() any {
	return nil
}

//go:inline
func (c *Program) U8() atf.U8 {
	v := atf.U8(c.Emit[c.Head])
	c.Head++
	return v
}

//go:inline
func (c *Program) U16() atf.U16 {
	v := atf.U16(binary.LittleEndian.Uint16(c.hd()))
	c.Head += 2
	return v
}

//go:inline
func (c *Program) hd() []byte {
	return c.Emit.Data()[c.Head:]
}

//go:inline
func (c *Program) U32() atf.U32 {
	v := atf.U32(binary.LittleEndian.Uint32(c.hd()))
	c.Head += 4
	return v
}

//go:inline
func (c *Program) U64() atf.U64 {
	v := atf.U64(binary.LittleEndian.Uint64(c.hd()))
	c.Head += 8
	return v
}

//go:inline
func (c *Program) U128() atf.U128 {
	var v atf.U128
	copy(v[:], c.hd())
	c.Head += 16
	return v
}

//go:inline
func (c *Program) U256() atf.U256 {
	var v atf.U256
	copy(v[:], c.hd())
	c.Head += 32
	return v
}

//go:inline
func (c *Program) I8() atf.I8 {
	v := int8(c.Emit[c.Head])
	c.Head++
	return v
}

//go:inline
func (c *Program) I16() atf.I16 {
	v := int16(binary.LittleEndian.Uint16(c.hd()))
	c.Head += 2
	return v
}

//go:inline
func (c *Program) I32() atf.I32 {
	v := int32(binary.LittleEndian.Uint32(c.hd()))
	c.Head += 4
	return v
}

//go:inline
func (c *Program) I64() atf.I64 {
	v := int64(binary.LittleEndian.Uint64(c.hd()))
	c.Head += 8
	return v
}

//go:inline
func (c *Program) I128() atf.I128 {
	var v atf.I128
	copy(v[:], c.hd())
	c.Head += 16
	return v
}

//go:inline
func (c *Program) I256() atf.I256 {
	var v atf.I256
	copy(v[:], c.hd())
	c.Head += 32
	return v
}

//go:inline
func (c *Program) F32() atf.F32 {
	bits := c.U32()
	v := atf.F32(math.Float32frombits(bits))
	return v
}

//go:inline
func (c *Program) F64() atf.F64 {
	bits := c.U64()
	v := atf.F64(math.Float64frombits(bits))
	return v
}

//go:inline
func (c *Program) Complex64() atf.Complex64 {
	u := c.U64()
	// safety: layout of complex64 is two float32
	realBits := uint32(u)
	imagBits := uint32(u >> 32)
	v := atf.Complex64(complex(
		math.Float32frombits(realBits),
		math.Float32frombits(imagBits),
	))
	return v
}

//go:inline
func (c *Program) Complex128() atf.Complex128 {
	rbits := c.U64()
	ibits := c.U64()
	v := atf.Complex128(complex(
		math.Float64frombits(rbits),
		math.Float64frombits(ibits),
	))
	return v
}

//go:inline
func (c *Program) PC() atf.PC {
	v := atf.PC(binary.LittleEndian.Uint32(c.hd()))
	c.Head += 4
	return v
}

//go:inline
func (c *Program) SP() atf.SP {
	v := atf.SP(binary.LittleEndian.Uint32(c.hd()))
	c.Head += 4
	return v
}

//go:inline
func (c *Program) Len() int {
	n := int(binary.LittleEndian.Uint32(c.hd()))
	c.Head += 4
	return n
}

//go:inline
func (c *Program) Bytes() atf.Bytes {
	n := c.Len()
	v := c.Emit[c.Head : c.Head+atf.PC(n)]
	c.Head += atf.PC(n)
	return atf.Bytes(v.Data())
}

//go:inline
func (c *Program) String() atf.String {
	n := c.Len()
	v := string(c.Emit[c.Head : c.Head+atf.PC(n)])
	c.Head += atf.PC(n)
	return atf.String(v)
}

//go:inline
func (c *Program) Error() atf.Error {
	n := c.Len()
	v := string(c.Emit[c.Head : c.Head+atf.PC(n)])
	c.Head += atf.PC(n)
	return errors.New(v)
}

//go:inline
func (c *Program) IBig() atf.IBig {
	// first byte: sign
	sign := c.Emit[c.Head]
	c.Head++
	n := c.Len()
	data := c.Emit[c.Head : c.Head+atf.PC(n)]
	c.Head += atf.PC(n)
	var bi big.Int
	bi.SetBytes(data.Data())
	if sign == 1 {
		bi.Neg(&bi)
	}
	return bi
}

//go:inline
func (c *Program) FBig() atf.FBig {
	n := c.Len()
	data := c.Emit[c.Head : c.Head+atf.PC(n)]
	c.Head += atf.PC(n)
	var bf big.Float
	bf.GobDecode(data.Data())
	return bf
}

//go:inline
func (c *Program) Time() atf.Time {
	ns := binary.LittleEndian.Uint64(c.hd())
	c.Head += 8
	return atf.Time(time.Unix(0, int64(ns)))
}

//go:inline
func (c *Program) Dur() atf.Dur {
	ns := binary.LittleEndian.Uint64(c.hd())
	c.Head += 8
	return atf.Dur(time.Duration(ns))
}

func NewEmitterCap(c int) atf.Emitter {
	return make(Emitter, 0, c)
}

func (c *Program) MarshalBinary() []byte {
	return NewEmitterCap(8 + len(c.Emit)).
		RawString("ASF1").
		RawU32(uint32(c.Head)).
		RawBytes(c.Emit.Data()).
		Data()
}

func (c *Program) UnmarshalBinary(src []byte) error {
	if len(src) < 8 {
		return ErrInvalid
	}
	c.Emit = NewEmitter(src[8:])
	c.Head = 4
	c.Head = atf.PC(c.U64())
	return nil
}
