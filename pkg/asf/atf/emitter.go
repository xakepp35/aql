// pkg/asf/atf/appender.go
package atf

// Emitter is a bytes roster for ASF TLV-format
type Emitter interface {
	Any(any) Emitter               // switch router, for manual use - prefer not to!
	Nop() Emitter                  // no operation - returns himself
	Nil() Emitter                  // appends nil interface
	U8(U8) Emitter                 // uint8
	U16(U16) Emitter               // uint16
	U32(U32) Emitter               // uint32
	U64(U64) Emitter               // uint64
	U128(U128) Emitter             // Uint128 [16]byte
	U256(U256) Emitter             // Uint256 [32]byte
	I8(I8) Emitter                 // int8
	I16(I16) Emitter               // int16
	I32(I32) Emitter               // int32
	I64(I64) Emitter               // int64
	I128(I128) Emitter             // Int128 [16]byte
	I256(I256) Emitter             // Int256 [32]byte
	F32(F32) Emitter               // float32
	F64(F64) Emitter               // float64
	Complex64(Complex64) Emitter   // complex64
	Complex128(Complex128) Emitter // complex128
	PC(PC) Emitter                 // uint32, program counter,
	SP(SP) Emitter                 // uint32, stack pointer,
	Time(Time) Emitter             // time.Time
	Dur(Dur) Emitter               // time.Duration, nanoseconds
	Bytes(Bytes) Emitter           // []byte, bytes
	String(String) Emitter         // string
	Error(Error) Emitter           // error, just like string, but indicates an error on the stack
	IBig(IBig) Emitter             // big.Int
	FBig(FBig) Emitter             // big.Float

	// raw append functions

	RawLen(l int) Emitter
	RawU16(v uint16) Emitter
	RawU32(v uint32) Emitter
	RawU64(v uint64) Emitter
	Raw(v ...byte) Emitter
	RawBytes(v []byte) Emitter
	RawString(v string) Emitter

	// byte slice fetcher, zero-copy
	Data() []byte
}
