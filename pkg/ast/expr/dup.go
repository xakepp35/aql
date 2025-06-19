package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Dup struct{}

func NewDup() asi.AST {
	return Dup{}
}

func (e Dup) Kind() asi.Kind {
	return asi.Dup
}

func (e Dup) P0(c asi.Emitter) error {
	return nil
}

func (e Dup) P1(c asi.Emitter) error {
	c.RawU8(atf.U8(op.Dup))
	return nil
}

func (e Dup) P2(c asi.Emitter) error {
	return nil
}

func (e Dup) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"dup"}`)
}

func (e Dup) BuildString(sb *strings.Builder) {
	sb.WriteString("[dup]")
}
