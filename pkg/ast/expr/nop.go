package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Nop struct{}

func NewNop() asi.AST {
	return Nop{}
}

func (e Nop) Kind() asi.Kind {
	return asi.Nop
}

func (e Nop) P0(c asi.Emitter) error {
	return nil
}

func (e Nop) P1(c asi.Emitter) error {
	return nil
}

func (e Nop) P2(c asi.Emitter) error {
	return nil
}

func (e Nop) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"nop"}`)
}

func (e Nop) BuildString(sb *strings.Builder) {
	sb.WriteString("[nop]")
}
