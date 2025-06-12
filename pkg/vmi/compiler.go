package vmi

import (
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Programmer interface {
	Compiler
	atf.Emitter
	asi.AST
	BinarySerializer
}

type Compiler interface {
	Compile(src []byte) error
}

type Loader interface {
	PC() int64
	Program() ByteCode
}
