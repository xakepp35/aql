package op

import "github.com/xakepp35/aql/pkg/asf"

//go:generate stringer -type=Code -trimprefix=
type Code asf.Type

const (
	// flow
	Halt Code = iota + 28
	Call
	Over
	Loop
	Break

	// stack

	Nil
	Pop
	Dup
	Swap
	Id

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
