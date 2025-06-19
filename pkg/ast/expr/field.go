package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Field struct {
	Arg  asi.AST
	Name []byte
}

func (e Field) Kind() asi.Kind {
	return asi.Field
}

func (e Field) P0(c asi.Emitter) error {
	return e.Arg.P0(c)
}

func (e Field) P1(c asi.Emitter) error {
	if err := e.Arg.P1(c); err != nil {
		return err
	}
	c.StringBytes(e.Name)
	c.RawU8(atf.U8(op.Field))
	return nil
}

func (e Field) P2(c asi.Emitter) error {
	return e.Arg.P2(c)
}

func (e Field) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"field","arg":`)
	e.Arg.BuildJSON(b)
	b.WriteString(`,"name":"`)
	b.Write(e.Name)
	b.WriteByte('"')
	b.WriteByte('}')
}

func (e Field) BuildString(b *strings.Builder) {
	b.WriteString("[field ")
	e.Arg.BuildString(b)
	b.WriteByte(' ')
	b.Write(e.Name)
	b.WriteByte(']')
}
