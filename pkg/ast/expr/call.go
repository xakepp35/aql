package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Call struct {
	Args []asi.AST
	Name []byte
}

func (e *Call) Kind() asi.Kind {
	return asi.Call
}

func (e *Call) P0(c asi.Emitter) error {
	for _, a := range e.Args {
		if err := a.P0(c); err != nil {
			return err
		}
	}
	return nil
}

func (e *Call) P1(c asi.Emitter) error {
	for _, a := range e.Args {
		if err := a.P1(c); err != nil {
			return err
		}
	}
	c.
		StringBytes(e.Name).
		RawU8(atf.U8(op.Call)).
		Commit()
	return nil
}

func (e *Call) P2(c asi.Emitter) error {
	for _, a := range e.Args {
		if err := a.P2(c); err != nil {
			return err
		}
	}
	return nil
}

func (e *Call) BuildJSON(b *bytes.Buffer) {
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

func (e *Call) BuildString(sb *strings.Builder) {
	sb.WriteString("[call ")
	for _, a := range e.Args {
		a.BuildString(sb)
		sb.WriteByte(' ')
	}
	sb.Write(e.Name)
	sb.WriteByte(']')
}
