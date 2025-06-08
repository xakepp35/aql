package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Binary struct {
	Left  vmi.AST
	Right vmi.AST
	Op    op.Code
}

func (e *Binary) Pre(c vmi.Compiler) error {
	if err := e.Left.Pre(c); err != nil {
		return err
	}
	if err := e.Right.Pre(c); err != nil {
		return err
	}
	return nil
}

func (e *Binary) Body(c vmi.Builder) error {
	if err := e.Left.Body(c); err != nil {
		return err
	}
	if err := e.Right.Body(c); err != nil {
		return err
	}
	c.Op(e.Op)
	return nil
}

func (e *Binary) Post(c vmi.Compiler) error {
	if err := e.Left.Post(c); err != nil {
		return err
	}
	if err := e.Right.Post(c); err != nil {
		return err
	}
	return nil
}

func (e *Binary) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"binary","op":"`)
	b.WriteString(e.Op.String())
	b.WriteString(`","left":`)
	e.Left.BuildJSON(b)
	b.WriteString(`,"right":`)
	e.Right.BuildJSON(b)
	b.WriteByte('}')
}

func (e *Binary) BuildString(b *strings.Builder) {
	b.WriteString("[binary ")
	e.Left.BuildString(b)
	b.WriteByte(' ')
	e.Right.BuildString(b)
	b.WriteByte(' ')
	b.WriteString(e.Op.String())
	b.WriteByte(']')
}
