package asi

// Kind represents the expression type in AQL AST.
// It helps the compiler and debug tools identify the nature of each AST node.
// This enum is designed to reflect *how the expression behaves*, not how it looks.
// Order is grouped by arity and semantics (stack ops, control flow, accessors, etc.)

type Kind byte

const (
	// --- No-op / Dummy ---
	Nop Kind = iota
	// Nop: Represents a no-op AST node. Used as a stub, placeholder, or dummy.
	// Literally does nothing and emits nothing.

	// --- Stack manipulation / internals ---
	Dup
	// Dup: Duplicate the top element of the stack.
	// Example: `.` (dot) is syntactic sugar for top-of-stack (duplicator).
	// This enables things like piping into the same value multiple times.

	This
	// This: Refers to the current context object in chained access or inside OVER/PIPE.
	// Not the same as "self" in OOP, but rather a contextual anchor.

	// --- Literal values / identifiers ---
	Literal
	// Literal: Any literal value: numbers, strings, true/false/null, etc.
	// Might also resolve identifiers via variables or constants map.

	// --- Unary operator ---
	Unary
	// Unary: Single-operand operator (e.g. !ok, -num, not flag).
	// Operand is evaluated and pushed before applying the operation.

	// --- Binary operator ---
	Binary
	// Binary: Standard two-operand operator (e.g. a + b, a == b, a && b).
	// Operands are pushed in left-to-right order, op is applied after.

	// --- Ternary operator ---
	Ternary
	// Ternary: Triple-operand operator like slicing: list[start:end].
	// Used for constructs such as Index2 or custom logic involving three operands.

	// Call: Function or method invocation.
	// All arguments are compiled (evaluated and pushed) BEFORE the call.
	// The function reference itself must already be on the stack.
	// Arity (argument count) is handled by callee, not pushed, thus var args are not supported
	Call

	// Field: Field access, e.g.
	//   .user.name
	// Pops an object from stack, pushes the value of the field.
	Field

	// Pipe: Pipe operator like `1 | string`.
	// Semantically binary, but emits no op — just preserves stack state.
	// Often used in functional chains or when redirecting flow into new exprs.
	Pipe

	// Over: Streaming for-each operator, e.g.
	//   over .items => sum
	// Compiles as: load iterator, enter loop, evaluate body, repeat, then exit.
	// Requires matching Over/Jmp bytecode sequence and uses iterator.Next().
	Over

	// --- [Optional: add here if needed later] ---
	// Async      // for async exprs: @N(expr) → Forks a goroutine VM for evaluating expr, returns top N stack values.
	// Match      // pattern matching?
	// Block      // block scoping (e.g. { expr1; expr2; })
	// Assign     // assignment node (e.g. a = b)
	// Lambda     // inline lambdas, closures?
	// IfExpr     // ternary-like conditional: a ? b : c
)
