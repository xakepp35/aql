package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Pipe struct {
	Args [2]asi.AST
}

func (e Pipe) Kind() asi.Kind {
	return asi.Pipe
}

func (e Pipe) P0(c asi.Emitter) error {
	if err := e.Args[0].P0(c); err != nil {
		return err
	}
	if err := e.Args[1].P0(c); err != nil {
		return err
	}
	return nil
}

func (e Pipe) P1(c asi.Emitter) error {
	if err := e.Args[0].P1(c); err != nil {
		return err
	}
	if err := e.Args[1].P1(c); err != nil {
		return err
	}
	// pipe just transfers stack to new section
	return nil
}

func (e Pipe) P2(c asi.Emitter) error {
	if err := e.Args[0].P2(c); err != nil {
		return err
	}
	if err := e.Args[1].P2(c); err != nil {
		return err
	}
	return nil
}

func (e Pipe) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"pipe","left":`)
	e.Args[0].BuildJSON(b)
	b.WriteString(`,"right":`)
	e.Args[1].BuildJSON(b)
	b.WriteByte('}')
}

func (e Pipe) BuildString(b *strings.Builder) {
	b.WriteString("[pipe ")
	e.Args[0].BuildString(b)
	b.WriteByte(' ')
	e.Args[1].BuildString(b)
	b.WriteByte(']')
}
