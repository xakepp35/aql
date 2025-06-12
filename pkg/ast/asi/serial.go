package asi

import (
	"bytes"
	"strings"
)

// JSONBuilder serializes the AST to JSON format, useful for introspection and testing.
type JSONBuilder interface {
	BuildJSON(*bytes.Buffer) // Serializes the AST to JSON (great for introspection and tests)
}

type StringBuilder interface {
	BuildString(*strings.Builder) // Outputs AST in a readable format (Reverse Polish FTW)
}
