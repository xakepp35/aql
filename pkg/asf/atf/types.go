// pkg/asf/atf/type.go
package atf

import (
	"math/big"
	"time"
)

type (
	Nil        = any // nil type
	U8         = uint8
	U16        = uint16
	U32        = uint32
	U64        = uint64
	U128       [16]byte
	U256       [32]byte
	I8         = int8
	I16        = int16
	I32        = int32
	I64        = int64
	I128       [16]byte
	I256       [32]byte
	F32        = float32
	F64        = float64
	Complex64  = complex64
	Complex128 = complex128
	PC         uint32 // program counter,
	SP         uint32 // stack pointer,
	Len        uint32 // length of something
	Time       = time.Time
	Dur        = time.Duration // nanoseconds
	Bytes      = []byte
	String     = string
	Error      = error // just like string, but indicates an error on the stack
	IBig       = big.Int
	FBig       = big.Float
)
