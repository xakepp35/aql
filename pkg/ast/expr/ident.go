package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Ident struct {
	Args []asi.AST
	Name []byte
}

func (e *Ident) Kind() asi.Kind {
	return asi.Ident
}

func (e *Ident) P0(c asi.Emitter) error {
	for _, a := range e.Args {
		if err := a.P0(c); err != nil {
			return err
		}
	}
	return nil
}

func (e *Ident) P1(c asi.Emitter) error {
	for _, a := range e.Args {
		if err := a.P1(c); err != nil {
			return err
		}
	}
	c.
		StringBytes(e.Name).
		RawU8(atf.U8(op.Ident)).
		Commit()
	return nil
}

func (e *Ident) P2(c asi.Emitter) error {
	for _, a := range e.Args {
		if err := a.P2(c); err != nil {
			return err
		}
	}
	return nil
}

func (e *Ident) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"call","name":"`)
	b.Write(e.Name)
	b.WriteString(`","args":[`)
	for i, a := range e.Args {
		if i > 0 {
			b.WriteByte(',')
		}
		a.BuildJSON(b)
	}
	b.WriteByte(']')
	b.WriteByte('}')
}

func (e *Ident) BuildString(sb *strings.Builder) {
	sb.WriteString("[call ")
	for _, a := range e.Args {
		a.BuildString(sb)
		sb.WriteByte(' ')
	}
	sb.Write(e.Name)
	sb.WriteByte(']')
}
