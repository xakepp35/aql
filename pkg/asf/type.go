// pkg/asf/type.go
package asf

// serial type
//go:generate stringer -type=Type -trimprefix=
type Type byte

const (
	Nop        Type = iota // do nothing
	Nil                    // nil
	U8                     // uint8
	U16                    // uint16
	U32                    // uint32
	U64                    // uint64
	U128                   // uint128 - custom implementation type U128 [16]byte
	U256                   // uint256 - custom implementation type U256 [32]byte
	I8                     // int8
	I16                    // int16
	I32                    // int32
	I64                    // int64
	I128                   // uint128 - custom implementation - type I128 [16]byte
	I256                   // uint256 - custom implementation - type I256 [32]byte
	F32                    // float32
	F64                    // float64
	Complex64              // complex64
	Complex128             // complex128
	PC                     // uint32, program counter,
	SP                     // uint32, stack pointer,
	Len                    // uint32, length of something
	Time                   // time.Time
	Dur                    // time.Duration, nanoseconds
	Bytes                  // []byte, bytes
	String                 // string
	Error                  // error, just like string, but indicates an error on the stack
	IBig                   // big.Int
	FBig                   // big.Float
	Type_Count             // overall counter of types
)
