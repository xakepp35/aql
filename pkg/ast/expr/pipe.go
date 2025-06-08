package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vmi"
)

type Pipe struct {
	Left, Right vmi.AST
}

func (e *Pipe) Pre(c vmi.Compiler) error {
	if err := e.Left.Pre(c); err != nil {
		return err
	}
	if err := e.Right.Pre(c); err != nil {
		return err
	}
	return nil
}

func (e *Pipe) Body(c vmi.Compiler) error {
	if err := e.Left.Body(c); err != nil {
		return err
	}
	if err := e.Right.Body(c); err != nil {
		return err
	}
	// pipe just transfers stack to new section
	return nil
}

func (e *Pipe) Post(c vmi.Compiler) error {
	if err := e.Left.Post(c); err != nil {
		return err
	}
	if err := e.Right.Post(c); err != nil {
		return err
	}
	return nil
}

func (e *Pipe) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"pipe","left":`)
	e.Left.BuildJSON(b)
	b.WriteString(`,"right":`)
	e.Right.BuildJSON(b)
	b.WriteByte('}')
}

func (e *Pipe) BuildString(b *strings.Builder) {
	b.WriteString("[pipe ")
	e.Left.BuildString(b)
	b.WriteByte(' ')
	e.Right.BuildString(b)
	b.WriteByte(']')
}
