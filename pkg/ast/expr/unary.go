package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Unary[N int] struct {
	Op    op.Code
	Child [N]vmi.AST
}

func (e *Unary) P0(c vmi.Compiler) error {
	return e.Right.P0(c)
}

func (e *Unary) P1(c vmi.Compiler) error {
	if err := e.Right.P1(c); err != nil {
		return err
	}
	c.Op(e.Op)
	return nil
}

func (e *Unary) P2(c vmi.Compiler) error {
	return e.Right.P2(c)
}

func (e *Unary) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"unary","op":"`)
	b.WriteString(e.Op.String())
	b.WriteString(`","right":`)
	e.Right.BuildJSON(b)
	b.WriteByte('}')
}

func (e *Unary) BuildString(b *strings.Builder) {
	b.WriteString("[unary ")
	e.Right.BuildString(b)
	b.WriteByte(' ')
	b.WriteString(e.Op.String())
	b.WriteByte(']')
}
