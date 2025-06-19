package expr

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Literal struct {
	X any
}

func (e Literal) Kind() asi.Kind {
	return asi.Literal
}

func (e Literal) P0(c asi.Emitter) error {
	return nil
}

func (e Literal) P1(c asi.Emitter) error {
	c.Any(e.X)
	return nil
}

func (e Literal) P2(c asi.Emitter) error {
	return nil
}

func (e Literal) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"literal","val":`)
	switch v := e.X.(type) {
	case string:
		b.WriteByte('"')
		b.WriteString(v)
		b.WriteByte('"')
	case int64:
		b.WriteString(strconv.FormatInt(v, 10))
	case bool:
		if v {
			b.WriteString("true")
		} else {
			b.WriteString("false")
		}
	case nil:
		b.WriteString("null")
	default:
		b.WriteString(`"?"`)
	}
	b.WriteByte('}')
}

func (e Literal) BuildString(b *strings.Builder) {
	switch v := e.X.(type) {
	case string:
		b.WriteString(`"`)
		b.WriteString(v)
		b.WriteString(`"`)
	case []byte:
		b.WriteString("0x")
		b.WriteString(hex.EncodeToString(v))
	case int64:
		b.WriteString(strconv.FormatInt(v, 10))
	case bool:
		if v {
			b.WriteString("true")
		} else {
			b.WriteString("false")
		}
	case nil:
		b.WriteString("null")
	default:
		b.WriteString("?")
	}
}
