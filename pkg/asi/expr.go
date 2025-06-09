package asi

import "github.com/xakepp35/aql/pkg/asf/atf"

// ASTExpr is a blazing-fast 3-phase visitor interface for AST nodes,
// designed for zero-alloc, zero-copy, hot-path bytecode emission.
//
// Each node emits bytecode via `atf.Emitter`, following a strict DFS (Depth-First Search)
// traversal order, split into three phases:
//
//	Phase 0 (P0): Pre-walk
//	  - Setup phase. Used to prepare subtrees, preload values, reserve memory,
//	    or evaluate any static context before generating actual instructions.
//	  - Runs strictly BEFORE phase P1 of the same node.
//
//	Phase 1 (P1): Main body
//	  - Emits main instructions. Examples:
//	      - c.Op(op.Add)
//	      - c.PushLiteral(42)
//	  - Ensures the correct stack layout for execution.
//
//	Phase 2 (P2): Post-walk
//	  - Cleanup, chaining, accessors (like `.x.y.z()`), or stream unrolling.
//	    Use it for delayed or chained operations that logically happen *after*
//	    the main expression (e.g., field access after a call).
//	  - Runs AFTER P1 of the same node and AFTER P2 of all child nodes.
//
// Expression: 1 + user.getSalary().amount
// ---------------------------------------
//
//	AST tree:
//	  Binary(
//	    Left: Literal(1),
//	    Right:
//	      Field(
//	        Base: Call(
//	          Fn: Field(
//	            Base: Ident("user"),
//	            Name: "getSalary"
//	          ),
//	          Args: []
//	        ),
//	        Name: "amount"
//	      ),
//	    Op: Add
//	  )
//
// Traversal & Emission:
// ----------------------
//
// PHASE P0 (Preparation):
// - Literal(1)       → c.I64(1)                   // push 1
// - Ident("user")    → c.Id("user")               // push user object
//
// PHASE P1 (Main logic):
// - Field("getSalary")  → c.Op(Field)             // access user.getSalary
// - Call()              → c.Op(Call)              // invoke getSalary()
// - Field("amount")     → c.Op(Field)             // access .amount on result of call
// - Binary Add          → c.Op(Add)               // add 1 + user.getSalary().amount
//
// PHASE P2 (Post-processing):
// - Ident("user")       → user.P2()               // optional: apply access tracking, .this adjustment
// - Field("getSalary")  → maybe attach debug info
// - Field("amount")     → check final access chain is valid
//
// Final Bytecode (schematic):
//
//	I64(1)
//	Id("user")
//	Op(Field)     // .getSalary
//	Op(Call)      // ()
//	Op(Field)     // .amount
//	Op(Add)
//
// All nodes must conform to this order to ensure consistent and predictable bytecode layout.
type Expr interface {
	Kind() Kind           // Returns the specific type (e.g., Binary, Call)
	P0(atf.Emitter) error // Pre-phase: setup, preload, DFS walk
	P1(atf.Emitter) error // Body-phase: emit main logic
	P2(atf.Emitter) error // Post-phase: chaining, cleanup
}

// type Kind byte

// const (
// 	Nop Kind = 0 // NO-OP, does nothing. Dummy stub thing
// 	// Unary operator
// 	//   !ok
// 	Unary
// 	// Binary operator
// 	//   a+b
// 	Binary
// 	Ternary // .list[1:2]
// 	Call    // sum(1,2,3) - calls function sum, expressions inside brackets should be calculated in advance and pushed on stack. argc is NOT pushed on stack! varargs are not supported
// 	This    // this - i saw it in zq but dont know how does it differ from .
// 	Dup     // . means stack duplicator
// 	Field   // .x
// 	Literal // a - tries to search unary operators map, then tries to search variables map
// 	Pipe    // 1 | string - pipe is no op
// 	Over
// )
