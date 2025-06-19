package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Binary struct {
	Args [2]asi.AST
	Op   op.Code
}

func NewBinary(a, b asi.AST, opCode op.Code) asi.AST {
	return &Binary{
		Args: [2]asi.AST{a, b},
		Op:   opCode,
	}
}

func (e *Binary) Kind() asi.Kind {
	return asi.Binary
}

func (e *Binary) P0(c asi.Emitter) error {
	if err := e.Args[0].P0(c); err != nil {
		return err
	}
	if err := e.Args[1].P0(c); err != nil {
		return err
	}
	return nil
}

func (e *Binary) P1(c asi.Emitter) error {
	if err := e.Args[0].P1(c); err != nil {
		return err
	}
	if err := e.Args[1].P1(c); err != nil {
		return err
	}
	c.RawU8(atf.U8(e.Op))
	return nil
}

func (e *Binary) P2(c asi.Emitter) error {
	if err := e.Args[0].P2(c); err != nil {
		return err
	}
	if err := e.Args[1].P2(c); err != nil {
		return err
	}
	return nil
}

func (e Binary) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"binary","op":"`)
	b.WriteString(e.Op.String())
	b.WriteString(`","args":[`)
	e.Args[0].BuildJSON(b)
	b.WriteByte(',')
	e.Args[1].BuildJSON(b)
	b.WriteString(`]}`)
}

func (e Binary) BuildString(b *strings.Builder) {
	b.WriteString("[binary ")
	e.Args[0].BuildString(b)
	b.WriteByte(' ')
	e.Args[1].BuildString(b)
	b.WriteByte(' ')
	b.WriteString(e.Op.String())
	b.WriteByte(']')
}
