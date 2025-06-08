package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Unary struct {
	Op    op.Code
	Right vmi.AST
}

func (e *Unary) Pre(c vmi.Compiler) error {
	return e.Right.Pre(c)
}

func (e *Unary) Body(c vmi.Compiler) error {
	if err := e.Right.Body(c); err != nil {
		return err
	}
	c.Ops(e.Op)
	return nil
}

func (e *Unary) Post(c vmi.Compiler) error {
	return e.Right.Post(c)
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
