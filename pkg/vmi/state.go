package vmi

import (
	"context"

	"github.com/xakepp35/aql/pkg/vm/op"
)

// OpFunc defines the signature for opcode handler functions.
type OpFunc func(this State)

type Program []op.Code

type JIT []OpFunc

// VM state hub
type State interface {
	Runner
	Executor
	Status
	Stack
	Variables
	Caller
	Context
	// Stream
}

type Runner interface {
	Run(Program)
	Runf(JIT)
	Next(Program) bool
}

type Executor interface {
	SetPC(uint)
	AddPC(int)
	PC() uint
}

type Status interface {
	SetErr(error)
	Err() error
}

// Caller execute functions from the vm code
type Caller interface {
	Call(fnName string)
	SetCall(fnName string, fn OpFunc)
}

// for io and math within the vm itself
type Stack interface {
	PushArgs(vals ...any)
	Args(n uint) []any // if stack is shorter than n - nil will be returned
	Push(v any)
	Pop() any
	Len() int
}

// Variables to access from within the VM
type Variables interface {
	SetVars(dict map[string]any)
	Vars() map[string]any
	Set(k string, v any)
	Get(k string) any
	Del(k string)
	Clear()
}

type NamedFuncs map[string]OpFunc

type Context interface {
	Context() context.Context
	SetContext(ctx context.Context)
}

// type Stream interface {
// 	Stream(v any)
// 	OpenStream() <-chan any
// }
