package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Field struct {
	X    vmi.AST
	Name []byte
}

func (e *Field) Pre(c vmi.Compiler) error {
	return e.X.Pre(c)
}

func (e *Field) Body(c vmi.Compiler) error {
	if err := e.X.Body(c); err != nil {
		return err
	}
	c.StringBytes(e.Name)
	c.Op(op.Field)
	return nil
}

func (e *Field) Post(c vmi.Compiler) error {
	return e.X.Post(c)
}

func (e *Field) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"field","x":`)
	e.X.BuildJSON(b)
	b.WriteString(`,"name":"`)
	b.Write(e.Name)
	b.WriteByte('"')
	b.WriteByte('}')
}

func (e *Field) BuildString(b *strings.Builder) {
	b.WriteString("[field ")
	e.X.BuildString(b)
	b.WriteByte(' ')
	b.Write(e.Name)
	b.WriteByte(']')
}
