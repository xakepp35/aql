package expr

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/xakepp35/aql/pkg/vmi"
)

type Literal struct {
	X any
}

func (e *Literal) Pre(c vmi.Compiler) error {
	return nil
}

func (e *Literal) Body(c vmi.Compiler) error {
	switch v := e.X.(type) {
	case int64:
		c.Int(v)
	case string:
		c.String(v)
	case bool:
		c.Bool(v)
	case nil:
		c.Null()
	default:
		return fmt.Errorf("unsupported literal type: %T", v)
	}
	return nil
}

func (e *Literal) Post(c vmi.Compiler) error {
	return nil
}

func (e *Literal) BuildJSON(b *bytes.Buffer) {
	b.WriteString(`{"expr":"literal","value":`)
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

func (e *Literal) BuildString(b *strings.Builder) {
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
