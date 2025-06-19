// Package vmi defines the core interface for our blazing fast,
// zero-copy / zero-alloc virtual machine. This ain't your grandma's VM.
package vmi

import (
	"context"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/asf/atf"
)

// Func defines a VM opcode or builtin function handler.
// Think of it like an instruction callback: when the VM hits this op, this runs.
type Func func(this VM)

// ByteCode is our compiled program — a glorious stream of opcodes ready to be executed.
type ByteCode []op.Code

// PC is a Program Counter. It points to the current instruction in the bytecode.
// Like a bookmark in the book of high-speed execution.
type PC uint64

// VM is the main interface that glues everything together.
// It combines multiple behaviors (Runner, Executor, Stack, etc.)
// into one ultra-powerful execution engine.
type VM interface {
	Runner           // controls the main run loop
	Executor         // controls program counter and error state
	Statuser         // gets current error status of execution
	Stacker          // performs math, logic, IO and magic
	Variabler        // gives the VM its working memory
	Functioner       // lets you plug in your own awesome ops
	Contexter        // handles cancellation and lifecycle
	Streamer         // channels for async & reactive data flow
	BinarySerializer // serializes memory and stack
}

// Runner controls the main VM execution loop.
// Call Run() to start and Next() to tick manually, instruction by instruction.
type Runner interface {
	Run(VM)       // fire up the engine — let’s gooooo
	Next(VM) bool // advance by one instruction; returns false if done or broken
}

// Executor gives you low-level control over the program flow and error state.
// Great for implementing conditional jumps, loops, and controlled crashes.
type Executor interface {
	Jmp(pc atf.PC) // jump to the given program counter (without looking back)
	PC() atf.PC    // get the current PC (handy for debugging)
}

type Statuser interface {
	Fail(err error) // set an error and immediately halt execution (sad face)
	Status() error  // get the current error (if any)
}

// Functioner lets you inject or override named functions.
// These are the high-level function calls like sum(), avg(), print(), etc.
type Functioner interface {
	SetCalls(map[string]Func)  // bulk-register your arsenal of functions
	SetCall(string, Func)      // register or override a single function
	GetCall(string) Func       // get a function by name (or nil if not found)
	GetCalls() map[string]Func // get all
}

// Stacker provides classic stack-machine operations.
// Push, pop, and get args — the bread and butter of zero-alloc computing.
type Stacker interface {
	StackManipulator
	StackReader
}

// StackManipulator api
type StackManipulator interface {
	Pushs(...any) Stacker  // push a bunch of args at once
	Push(v any) Stacker    // push a single value to the stack
	Pops(count uint) []any // retrieve the last n args (non-destructively). nil if not enough
	Pop() any              // pop the top value from the stack
}

// StackReader api
type StackReader interface {
	Depth() int  // how deep is the stack?
	Dump() []any // peek at the whole stack (dev/debug friendly)
}

// Variabler is the mutable working memory of your VM.
// Think of it like a local symbol table — define, get, and clear named values.
type Variabler interface {
	SetVars(dict map[string]any) // inject a full map of vars (atomic overwrite)
	Vars() map[string]any        // snapshot of all variables
	Set(k string, v any)         // set var k = v
	Get(k string) (any, bool)    // get var by name, returns (nil, false) if not found
	Del(k string)                // delete var by name
	Clear()                      // nuke all variables
}

// Contexter manages cancellation & deadline propagation.
// Useful for killing long-running jobs or integrating into real-world systems.
type Contexter interface {
	Ctx() context.Context                                  // get current context (can be used inside functions)
	Cancel()                                               // signal cancelation — next op will fail
	SetCtx(ctx context.Context, cancel context.CancelFunc) // swap in new context+cancel pair
}

// Streamer is for reactive use-cases — think async channels.
// Send() pushes output values, Recv() pulls incoming ones.
// Ideal for pipelines and stream-processing logic.
type Streamer interface {
	Send(any)  // channel to send values out of the VM
	Recv() any // channel to receive input values
}
