package vmi

import "github.com/xakepp35/aql/pkg/vm/op"

type Compiler interface {
	Compile(src []byte) error
	Emitter
	Patcher
	Loader
	BinarySerializer
}

type Emitter interface {
	EmitNull() int
	EmitInt(int64) int
	EmitBool(bool) int
	EmitString(string) int
	EmitOps(...op.Code) int
}

type Patcher interface {
	PatchInt(pos int, v int)
}

type Loader interface {
	Init(State)
	Program() Program
	JIT() JIT
}

type BinarySerializer interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}
