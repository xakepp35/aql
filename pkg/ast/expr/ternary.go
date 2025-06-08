package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Ternary struct {
	Left   vmi.AST
	Right1 vmi.AST
	Right2 vmi.AST
	Op     op.Code
}

func (e *Ternary) Pre(c vmi.Compiler) error {
	if err := e.Left.Pre(c); err != nil {
		return err
	}
	if err := e.Right1.Pre(c); err != nil {
		return err
	}
	if err := e.Right2.Pre(c); err != nil {
		return err
	}
	return nil
}

func (e *Ternary) Body(c vmi.Compiler) error {
	if err := e.Left.Body(c); err != nil {
		return err
	}
	if err := e.Right1.Body(c); err != nil {
		return err
	}
	if err := e.Right2.Body(c); err != nil {
		return err
	}
	c.Op(e.Op)
	return nil
}

func (e *Ternary) Post(c vmi.Compiler) error {
	if err := e.Left.Post(c); err != nil {
		return err
	}
	if err := e.Right1.Post(c); err != nil {
		return err
	}
	if err := e.Right2.Post(c); err != nil {
		return err
	}
	return nil
}

func (e *Ternary) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"ternary","op":"`)
	b.WriteString(e.Op.String())
	b.WriteString(`","left":`)
	e.Left.BuildJSON(b)
	b.WriteString(`,"right1":`)
	e.Right1.BuildJSON(b)
	b.WriteString(`,"right2":`)
	e.Right2.BuildJSON(b)
	b.WriteByte('}')
}

func (e *Ternary) BuildString(b *strings.Builder) {
	b.WriteString("[ternary ")
	e.Left.BuildString(b)
	b.WriteByte(' ')
	e.Right1.BuildString(b)
	b.WriteByte(' ')
	e.Right2.BuildString(b)
	b.WriteByte(' ')
	b.WriteString(e.Op.String())
	b.WriteByte(']')
}
