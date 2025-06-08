package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Call struct {
	Args []vmi.AST
	Name []byte
}

func (e *Call) Pre(c vmi.Compiler) error {
	for _, a := range e.Args {
		if err := a.Pre(c); err != nil {
			return err
		}
	}
	return nil
}

func (e *Call) Body(c vmi.Compiler) error {
	for _, a := range e.Args {
		if err := a.Body(c); err != nil {
			return err
		}
	}
	c.StringBytes(e.Name)
	c.Int(int64(len(e.Args)))
	c.Op(op.Call)
	return nil
}

func (e *Call) Post(c vmi.Compiler) error {
	for _, a := range e.Args {
		if err := a.Post(c); err != nil {
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
