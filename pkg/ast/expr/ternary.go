package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Ternary struct {
	Args [3]vmi.AST
	Op   op.Code
}

func (e *Ternary) P0(c vmi.Compiler) error {
	if err := e.Args[0].P0(c); err != nil {
		return err
	}
	if err := e.Args[1].P0(c); err != nil {
		return err
	}
	if err := e.Args[2].P0(c); err != nil {
		return err
	}
	return nil
}

func (e *Ternary) P1(c vmi.Compiler) error {
	if err := e.Args[0].P1(c); err != nil {
		return err
	}
	if err := e.Args[1].P1(c); err != nil {
		return err
	}
	if err := e.Args[2].P1(c); err != nil {
		return err
	}
	c.Op(e.Op)
	return nil
}

func (e *Ternary) P2(c vmi.Compiler) error {
	if err := e.Args[0].P2(c); err != nil {
		return err
	}
	if err := e.Args[1].P2(c); err != nil {
		return err
	}
	if err := e.Args[2].P2(c); err != nil {
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
