package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Swap struct{}

func NewSwap() asi.AST {
	return Swap{}
}

func (e Swap) Kind() asi.Kind {
	return asi.Swap
}

func (e Swap) P0(c asi.Emitter) error {
	return nil
}

func (e Swap) P1(c asi.Emitter) error {
	return nil
}

func (e Swap) P2(c asi.Emitter) error {
	return nil
}

func (e Swap) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"nop"}`)
}

func (e Swap) BuildString(sb *strings.Builder) {
	sb.WriteString("[nop]")
}
