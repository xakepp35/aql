package op

//go:generate stringer -type=Code -trimprefix=
type Code byte

const (
	// flow
	Halt Code = iota + 28
	Call
	Over
	Loop
	Break

	// stack

	Pop
	Dup
	Swap
	Ident

	// logic & math

	Not
	And
	Or
	Xor
	Shl
	Shr
	Add
	Sub
	Mul
	Div
	Mod

	// comparison

	Eq
	Neq
	Lt
	Le
	Gt
	Ge

	// data

	Pipe
	Index1
	Index2
	Field
)
