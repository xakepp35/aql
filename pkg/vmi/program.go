package vmi

import (
	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
)

type Program interface {
	Op() op.Code
	Nil() any
	U8() atf.U8
	U16() atf.U16
	U32() atf.U32
	U64() atf.U64
	U128() atf.U128
	U256() atf.U256
	I8() atf.I8
	I16() atf.I16
	I32() atf.I32
	I64() atf.I64
	I128() atf.I128
	I256() atf.I256
	F32() atf.F32
	F64() atf.F64
	Complex64() atf.Complex64
	Complex128() atf.Complex128
	PC() atf.PC
	SP() atf.SP
	Len() int
	Bytes() atf.Bytes
	String() atf.String
	Error() atf.Error
	IBig() atf.IBig
	FBig() atf.FBig
	Time() atf.Time
	Dur() atf.Dur
}

type BinarySerializer interface {
	MarshalBinary() []byte
	UnmarshalBinary([]byte) error
}
