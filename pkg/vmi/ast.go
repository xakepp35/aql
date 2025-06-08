package vmi

import (
	"bytes"
	"strings"
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
	ASTCompiler
}

// ASTCompiler recursive BFS 3-phase visitor
type ASTCompiler interface {
	Pre(Compiler) error  // Pre-compilation phase (setup, left side of ops, etc.)
	Body(Compiler) error // Main body compilation (emit the actual instruction logic)
	Post(Compiler) error // Post-compilation (access chains, deferred ops, etc.)
}

type ASTKind int32

const (
	AST_Binary  ASTKind = 0 // a+b
	AST_Call                // sum(1,2,3)
	AST_Dup                 // .
	AST_Field               // .x
	AST_Literal             // a
	AST_Pipe                // 1 | string
	AST_Over                // over .salaries => sum(.)
	AST_Ternary             // .list[1:2]
	AST_Unary               // !ok
)
