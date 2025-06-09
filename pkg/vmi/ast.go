package vmi

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/asf/atf"
)

// AST defines the interface for Abstract Syntax Tree nodes in AQL.
//
// Each node is split into 3 compilation phases:
// - Pre(): emit instructions *before* evaluating the main node
// - Body(): emit the core logic (the "meat")
// - Post(): emit anything that comes *after*, like field/index access or deferred logic
//
// This tri-phase layout gives us superpowers for code-gen:
//
// BuildJSON and BuildString let you serialize or debug the AST in clean formats.
//
// TL;DR: this interface is your portal to precise, phased bytecode emission.
//
// ⚡ It's like LEGO for compilers. But faster. And with zero GC. ⚡
type AST interface {
	Kind() ASTKind                           // Current node kind, or type
	Fold(acc any, fn func(any, AST) any) any // Accumulates a value over tree traversal - perfect for metrics, depth calc
	Walk(acc any, fn func(AST) error) error  // General-purpose traversal. Calls fn on each node in BFS order.
	BuildJSON(*bytes.Buffer)                 // Serializes the AST to JSON (great for introspection and tests)
	BuildString(*strings.Builder)            // Outputs AST in a readable format (Reverse Polish FTW)
	ASTExpr
}
