package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Unary struct {
	Arg asi.AST
	Op  op.Code
}

func (e Unary) Kind() asi.Kind {
	return asi.Unary
}

func (e Unary) P0(c asi.Emitter) error {
	return e.Arg.P0(c)
}

func (e Unary) P1(c asi.Emitter) error {
	if err := e.Arg.P1(c); err != nil {
		return err
	}
	c.U8(atf.U8(e.Op))
	return nil
}

func (e Unary) P2(c asi.Emitter) error {
	return e.Arg.P2(c)
}

func (e Unary) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"unary","op":"`)
	b.WriteString(e.Op.String())
	b.WriteString(`","right":`)
	e.Arg.BuildJSON(b)
	b.WriteByte('}')
}

func (e Unary) BuildString(b *strings.Builder) {
	b.WriteString("[unary ")
	e.Arg.BuildString(b)
	b.WriteByte(' ')
	b.WriteString(e.Op.String())
	b.WriteByte(']')
}
