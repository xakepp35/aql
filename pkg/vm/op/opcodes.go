package op

//go:generate stringer -type=Code -trimprefix=
type Code byte

const (
	// flow

	Nop Code = iota
	Call
	Over
	Loop
	Break
	Halt

	// stack

	PushNil
	PushNow
	PushVar
	Pop
	Dup
	Swap

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

	// stats
	Avg
	Count
	Min
	Max
	Topk

	// converters

	ToInt
	ToFloat
	ToString
	ToBool
	ToTime

	// crypto

	FNV1a

	// data formats

	PackKV
	PackJSON
	UnpackJSON
	UnpackKV

	// data

	Pipe
	Index1
	Index2
	Field

	// iterators

	Iter      // convert stack object to Iterator
	HashProbe //	map + key → hashSingleIter (O(1) точка)
	RBSeek    // root + lo + hi → rbSeekIter
	Limit     // + n → limitIter
	Batch     //iterator + n → batchIter (для групповго commit)

	//
	From
	Store
	Upsert
	Update
	Delete
	Txn
)
