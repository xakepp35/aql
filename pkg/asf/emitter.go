// pkg/asf/appender.go
package asf

import (
	"encoding/hex"
	"math"
	"unsafe"

	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/util"
)

type Emitter []byte

var _ atf.Emitter = (*Emitter)(nil)

//go:inline
func NewEmitter(d []byte) *Emitter {
	return (*Emitter)(unsafe.Pointer(&d))
}

//go:inline
func NewEmitterCap(c int) atf.Emitter {
	res := make(Emitter, 0, c)
	return &res
}

//go:inline
func (s *Emitter) Data() []byte {
	return *s
}

//go:inline
func (s *Emitter) Len() atf.PC {
	return atf.PC(len(*s))
}

//go:inline
func (s *Emitter) Commit() {}

//go:inline
func (s *Emitter) AsHex() string {
	return hex.EncodeToString(*s)
}

//go:inline
func (s *Emitter) Any(v any) atf.Emitter {
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

//go:inline
func (s *Emitter) Nop() atf.Emitter {
	return s
}

//go:inline
func (s *Emitter) Nil() atf.Emitter {
	return s.
		RawU8(byte(Nil))

}

//go:inline
func (s *Emitter) U8(v atf.U8) atf.Emitter {
	return s.Raw(
		byte(U8),
		v,
	)

}

//go:inline
func (s *Emitter) U16(v atf.U16) atf.Emitter {
	return s.
		RawU8(byte(U16)).
		RawU16(v)
}

//go:inline
func (s *Emitter) U32(v atf.U32) atf.Emitter {
	return s.
		RawU8(byte(U32)).
		RawU32(v)
}

//go:inline
func (s *Emitter) U64(v atf.U64) atf.Emitter {
	return s.
		RawU8(byte(U64)).
		RawU64(v)
}

//go:inline
func (s *Emitter) U128(v atf.U128) atf.Emitter {
	return s.
		RawU8(byte(U128)).
		Raw(v[:]...)
}

//go:inline
func (s *Emitter) U256(v atf.U256) atf.Emitter {
	return s.
		RawU8(byte(U256)).
		Raw(v[:]...)
}

//go:inline
func (s *Emitter) I8(v atf.I8) atf.Emitter {
	return s.Raw(
		byte(I8),
		byte(v),
	)
}

//go:inline
func (s *Emitter) I16(v atf.I16) atf.Emitter {
	return s.
		RawU8(byte(I16)).
		RawU16(uint16(v))
}

//go:inline
func (s *Emitter) I32(v atf.I32) atf.Emitter {
	return s.
		RawU8(byte(I32)).
		RawU32(uint32(v))
}

//go:inline
func (s *Emitter) I64(v atf.I64) atf.Emitter {
	return s.RawU8(byte(I64)).
		RawU64(uint64(v))
}

//go:inline
func (s *Emitter) I128(v atf.I128) atf.Emitter {
	return s.
		RawU8(byte(I128)).
		Raw(v[:]...)
}

//go:inline
func (s *Emitter) I256(v atf.I256) atf.Emitter {
	return s.RawU8(byte(I256)).
		Raw(v[:]...)
}

//go:inline
func (s *Emitter) F32(v atf.F32) atf.Emitter {
	bits := math.Float32bits(v)
	return s.
		RawU8(byte(F32)).
		RawU32(bits)
}

//go:inline
func (s *Emitter) F64(v atf.F64) atf.Emitter {
	bits := math.Float64bits(v)
	return s.
		RawU8(byte(F64)).
		RawU64(bits)
}

//go:inline
func (s *Emitter) Complex64(v atf.Complex64) atf.Emitter {
	return s.
		RawU8(byte(Complex64)).
		RawU64(*(*uint64)(unsafe.Pointer(&v)))
}

//go:inline
func (s *Emitter) Complex128(v atf.Complex128) atf.Emitter {
	realBits := math.Float64bits(real(v))
	imagBits := math.Float64bits(imag(v))
	return s.
		RawU8(byte(Complex128)).
		RawU64(realBits).
		RawU64(imagBits)
}

//go:inline
func (s *Emitter) PC(v atf.PC) atf.Emitter {
	return s.
		RawU8(byte(PC)).
		RawU32(uint32(v))
}

//go:inline
func (s *Emitter) SP(v atf.SP) atf.Emitter {
	return s.
		RawU8(byte(SP)).
		RawU32(uint32(v))
}

//go:inline
func (s *Emitter) Time(t atf.Time) atf.Emitter {
	return s.
		RawU8(byte(Time)).
		RawU64(uint64(t.UnixNano()))
}

//go:inline
func (s *Emitter) Dur(d atf.Dur) atf.Emitter {
	return s.
		RawU8(byte(Dur)).
		RawU64(uint64(d))
}

//go:inline
func (s *Emitter) Bytes(v atf.Bytes) atf.Emitter {
	return s.
		RawU8(byte(Bytes)).
		RawLen(len(v)).
		Raw(v...)
}

//go:inline
func (s *Emitter) String(v atf.String) atf.Emitter {
	return s.
		RawU8(byte(String)).
		RawLen(len(v)).
		RawString(v)
}

//go:inline
func (s *Emitter) StringBytes(v atf.Bytes) atf.Emitter {
	return s.
		RawU8(byte(String)).
		RawLen(len(v)).
		Raw(v...)
}

//go:inline
func (s *Emitter) Error(v atf.Error) atf.Emitter {
	data := v.Error()
	return s.
		RawU8(byte(Error)).
		RawLen(len(data)).
		RawString(data)
}

func sign(v int) byte {
	if v < 0 {
		return 1
	}
	return 0
}

//go:inline
func (s *Emitter) IBig(v atf.IBig) atf.Emitter {
	data := v.Bytes()
	return s.
		RawU8(byte(IBig)).
		RawU8(sign(v.Sign())).
		RawLen(len(data)).
		Raw(data...)
}

//go:inline
func (s *Emitter) FBig(v atf.FBig) atf.Emitter {
	data, _ := v.GobEncode()
	return s.
		RawU8(byte(FBig)).
		RawLen(len(data)).
		Raw(data...)
}

//go:inline
func (s *Emitter) RawLen(l int) atf.Emitter {
	return s.
		RawU32(uint32(l))
}

//go:inline
func (s *Emitter) RawU8(v uint8) atf.Emitter {
	*s = append(*s, v)
	return s
}

//go:inline
func (s *Emitter) RawU16(v uint16) atf.Emitter {
	return s.Raw(
		byte(v),
		byte(v>>8),
	)
}

//go:inline
func (s *Emitter) RawU32(v uint32) atf.Emitter {
	return s.Raw(
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
	)
}

//go:inline
func (s *Emitter) RawU64(v uint64) atf.Emitter {
	return s.Raw(
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
		byte(v>>32),
		byte(v>>40),
		byte(v>>48),
		byte(v>>56),
	)
}

//go:inline
func (s *Emitter) Raw(v ...byte) atf.Emitter {
	*s = append(*s, v...)
	return s

}

//go:inline
func (s *Emitter) RawBytes(v []byte) atf.Emitter {
	*s = append(*s, v...)
	return s

}

//go:inline
func (s *Emitter) RawString(v string) atf.Emitter {
	*s = append(*s, v...)
	return s
}
