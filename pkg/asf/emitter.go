// pkg/asf/appender.go
package asf

import (
	"math"
	"unsafe"

	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/util"
)

type Emitter []Type

var _ atf.Emitter = (Emitter)(nil)

// go:inline
func NewEmitter(d []byte) Emitter {
	return *(*Emitter)(unsafe.Pointer(&d))
}

// go:inline
func (s Emitter) Data() []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// go:inline
func (s Emitter) Any(v any) atf.Emitter {
	switch vt := v.(type) {
	case nil:
		return s.Nil()
	case atf.U8:
		return s.U8(vt)
	case atf.U16:
		return s.U16(vt)
	case atf.U32:
		return s.U32(vt)
	case atf.U64:
		return s.U64(vt)
	case atf.U128:
		return s.U128(vt)
	case atf.U256:
		return s.U256(vt)
	case atf.I8:
		return s.I8(vt)
	case atf.I16:
		return s.I16(vt)
	case atf.I32:
		return s.I32(vt)
	case atf.I64:
		return s.I64(vt)
	case atf.I128:
		return s.I128(vt)
	case atf.I256:
		return s.I256(vt)
	case atf.F32:
		return s.F32(vt)
	case atf.F64:
		return s.F64(vt)
	case atf.Complex64:
		return s.Complex64(vt)
	case atf.Complex128:
		return s.Complex128(vt)
	case atf.PC:
		return s.PC(vt)
	case atf.SP:
		return s.SP(vt)
	case atf.Len:
		return s.Len(vt)
	case atf.Time:
		return s.Time(vt)
	case atf.Dur:
		return s.Dur(vt)
	case atf.Bytes:
		return s.Bytes(vt)
	case atf.String:
		return s.String(vt)
	case atf.Error:
		return s.Error(vt)
	case atf.IBig:
		return s.IBig(vt)
	case atf.FBig:
		return s.FBig(vt)
	default:
		panic("asf.Appender: unsupported type: " + util.TypeOf(v).String())
	}
}

// go:inline
func (s Emitter) Nop() atf.Emitter {
	return s
}

// go:inline
func (s Emitter) Nil() atf.Emitter {
	s = append(s, Nil)
	return s
}

// go:inline
func (s Emitter) U8(v atf.U8) atf.Emitter {
	s = append(s, U8, Type(v))
	return s
}

// go:inline
func (s Emitter) U16(v atf.U16) atf.Emitter {
	s = append(s, U16)
	return s.RawU16(v)
}

// go:inline
func (s Emitter) U32(v atf.U32) atf.Emitter {
	s = append(s, U32)
	return s.RawU32(v)
}

// go:inline
func (s Emitter) U64(v atf.U64) atf.Emitter {
	s = append(s, U64)
	return s.RawU64(v)
}

// go:inline
func (s Emitter) U128(v atf.U128) atf.Emitter {
	s = append(s, U128)
	return s.Raw(v[:]...)
}

// go:inline
func (s Emitter) U256(v atf.U256) atf.Emitter {
	s = append(s, U256)
	return s.Raw(v[:]...)
}

// go:inline
func (s Emitter) I8(v atf.I8) atf.Emitter {
	s = append(s, I8, Type(v))
	return s
}

// go:inline
func (s Emitter) I16(v atf.I16) atf.Emitter {
	s = append(s, I16)
	return s.RawU16(uint16(v))
}

// go:inline
func (s Emitter) I32(v atf.I32) atf.Emitter {
	s = append(s, I32)
	return s.RawU32(uint32(v))
}

// go:inline
func (s Emitter) I64(v atf.I64) atf.Emitter {
	s = append(s, I64)
	return s.RawU64(uint64(v))
}

// go:inline
func (s Emitter) I128(v atf.I128) atf.Emitter {
	s = append(s, I128)
	return s.Raw(v[:]...)
}

// go:inline
func (s Emitter) I256(v atf.I256) atf.Emitter {
	s = append(s, I256)
	return s.Raw(v[:]...)
}

// go:inline
func (s Emitter) F32(v atf.F32) atf.Emitter {
	s = append(s, F32)
	bits := math.Float32bits(v)
	return s.RawU32(bits)
}

// go:inline
func (s Emitter) F64(v atf.F64) atf.Emitter {
	s = append(s, F64)
	bits := math.Float64bits(v)
	return s.RawU64(bits)
}

// go:inline
func (s Emitter) Complex64(v atf.Complex64) atf.Emitter {
	s = append(s, Complex64)
	return s.RawU64(*(*uint64)(unsafe.Pointer(&v)))
}

// go:inline
func (s Emitter) Complex128(v atf.Complex128) atf.Emitter {
	s = append(s, Complex128)
	realBits := math.Float64bits(real(v))
	imagBits := math.Float64bits(imag(v))
	return s.
		RawU64(realBits).
		RawU64(imagBits)
}

// go:inline
func (s Emitter) PC(v atf.PC) atf.Emitter {
	s = append(s, PC)
	return s.RawU32(uint32(v))
}

// go:inline
func (s Emitter) SP(v atf.SP) atf.Emitter {
	s = append(s, SP)
	return s.RawU32(uint32(v))
}

// go:inline
func (s Emitter) Len(v atf.Len) atf.Emitter {
	s = append(s, Len)
	return s.RawU32(uint32(v))
}

// go:inline
func (s Emitter) Time(t atf.Time) atf.Emitter {
	s = append(s, Time)
	return s.RawU64(uint64(t.UnixNano()))
}

// go:inline
func (s Emitter) Dur(d atf.Dur) atf.Emitter {
	s = append(s, Dur)
	return s.RawU64(uint64(d))
}

// go:inline
func (s Emitter) Bytes(v atf.Bytes) atf.Emitter {
	s = append(s, Bytes)
	return s.
		RawLen(len(v)).
		Raw(v...)
}

// go:inline
func (s Emitter) String(v atf.String) atf.Emitter {
	s = append(s, String)
	return s.
		RawLen(len(v)).
		RawString(v)
}

// go:inline
func (s Emitter) Error(v atf.Error) atf.Emitter {
	s = append(s, Error)
	data := v.Error()
	return s.
		RawLen(len(data)).
		RawString(data)
}

// go:inline
func (s Emitter) IBig(v atf.IBig) atf.Emitter {
	s = append(s, IBig)
	var sign Type
	if v.Sign() < 0 {
		sign = 1
	}
	s = append(s, sign)
	data := v.Bytes()
	return s.
		RawLen(len(data)).
		Raw(data...)
}

// go:inline
func (s Emitter) FBig(v atf.FBig) atf.Emitter {
	s = append(s, FBig)
	data, _ := v.GobEncode()
	return s.
		RawLen(len(data)).
		Raw(data...)
}

// go:inline
func (s Emitter) RawLen(l int) atf.Emitter {
	return s.RawU32(uint32(l))
}

// go:inline
func (s Emitter) RawU16(v uint16) atf.Emitter {
	return append(s,
		Type(v),
		Type(v>>8),
	)
}

// go:inline
func (s Emitter) RawU32(v uint32) atf.Emitter {
	return append(s,
		Type(v),
		Type(v>>8),
		Type(v>>16),
		Type(v>>24),
	)
}

// go:inline
func (s Emitter) RawU64(v uint64) atf.Emitter {
	return append(s,
		Type(v),
		Type(v>>8),
		Type(v>>16),
		Type(v>>24),
		Type(v>>32),
		Type(v>>40),
		Type(v>>48),
		Type(v>>56),
	)
}

// go:inline
func (s Emitter) Raw(v ...byte) atf.Emitter {
	return s.RawBytes(v)
}

// go:inline
func (s Emitter) RawBytes(v []byte) atf.Emitter {
	for _, c := range v {
		s = append(s, Type(c))
	}
	return s
}

// go:inline
func (s Emitter) RawString(v string) atf.Emitter {
	for _, c := range v {
		s = append(s, Type(c))
	}
	return s
}
