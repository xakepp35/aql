package vmi

type Type = byte

const (
	TypeUint64 Type = iota // целое (и счётчик для OpCall)
	TypeString             // смещение в Program.Data
	TypeBool               // 0/1
	TypeNull               // всегда 0
)

type Node interface {
	Evaluater
}

type Evaluater interface {
	Eval(this State) error
}
