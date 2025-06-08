package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Dup struct{}

func (e *Dup) Pre(c vmi.Compiler) error {
	return nil
}

func (e *Dup) Body(c vmi.Compiler) error {
	c.Op(op.Dup)
	return nil
}

func (e *Dup) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"dup"}`)
}

func (e *Dup) BuildString(sb *strings.Builder) {
	sb.WriteString("[dup]")
}
