package vmi

import (
	"github.com/xakepp35/aql/pkg/asf/atf"
)

type Programmer interface {
	Compiler
	atf.Emitter
	AST
	BinarySerializer
}

type Compiler interface {
	Compile(src []byte) error
}

type Loader interface {
	PC() int64
	Program() ByteCode
}
