package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Ternary struct {
	Args [3]asi.AST
	Op   op.Code
}

func (e Ternary) Kind() asi.Kind {
	return asi.Ternary
}

func (e Ternary) P0(c asi.Emitter) error {
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

func (e Ternary) P1(c asi.Emitter) error {
	if err := e.Args[0].P1(c); err != nil {
		return err
	}
	if err := e.Args[1].P1(c); err != nil {
		return err
	}
	if err := e.Args[2].P1(c); err != nil {
		return err
	}
	c.U8(atf.U8(e.Op))
	return nil
}

func (e Ternary) P2(c asi.Emitter) error {
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

func (e Ternary) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"ternary","op":"`)
	b.WriteString(e.Op.String())
	b.WriteString(`","args":[`)
	e.Args[0].BuildJSON(b)
	b.WriteByte(',')
	e.Args[1].BuildJSON(b)
	b.WriteByte(',')
	e.Args[2].BuildJSON(b)
	b.WriteString(`]}`)
}

func (e *Ternary) BuildString(b *strings.Builder) {
	b.WriteString("[ternary ")
	e.Args[0].BuildString(b)
	b.WriteByte(' ')
	e.Args[1].BuildString(b)
	b.WriteByte(' ')
	e.Args[2].BuildString(b)
	b.WriteByte(' ')
	b.WriteString(e.Op.String())
	b.WriteByte(']')
}
