package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/ast/asi"
)

type This struct{}

func NewThis() asi.AST {
	return This{}
}

func (e This) Kind() asi.Kind {
	return asi.This
}

func (e This) P0(c asi.Emitter) error {
	return nil
}

func (e This) P1(c asi.Emitter) error {
	c.RawU8(atf.U8(op.This))
	return nil
}

func (e This) P2(c asi.Emitter) error {
	return nil
}

func (e This) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"this"}`)
}

func (e This) BuildString(sb *strings.Builder) {
	sb.WriteString("[this]")
}
