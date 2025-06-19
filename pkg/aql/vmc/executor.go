package vmc

import (
	"context"
	"time"

	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/asf/atf"
)

type Breakpoint = func(*Executor) bool

type Err = error

type Gas struct {
	Spent uint64
	Limit uint64
}
type Executor struct {
	asf.Program
	Stack
	Gas
	Err
	context.Context
	context.CancelFunc
	Breakpoint
}

func New(ctx context.Context, cancel context.CancelFunc) Executor {
	if ctx == nil {
		ctx = context.Background()
	}
	if cancel == nil {
		ctx, cancel = context.WithCancel(ctx)
	}
	return Executor{
		Stack: make(Stack, 0, 8),
		Program: asf.Program{
			Emit: asf.Emitter{},
			Head: 0,
		},
		Err:        nil,
		Context:    ctx,
		CancelFunc: cancel,
	}
}

// Reset state
//
//go:inline
func (s *Executor) Reset() {
	s.Head = 0
	s.Stack = s.Stack[:0]
	s.Err = nil
}

//go:inline
func (s *Executor) active() bool {
	return (s.Gas.Limit == 0 || s.Gas.Spent < s.Gas.Limit) && s.Head < s.Emit.Len() && s.Err == nil
}

//go:inline
func (s *Executor) Active() bool {
	if !s.active() {
		return false
	}
	if s.Gas.Spent&0xf != 0 {
		return true
	}
	select {
	case <-s.Context.Done():
		return false
	default:
		return true
	}
}

//go:inline
// func (s *Executor) Active() bool {
// 	if s.Gas.Spent&0xf != 0 {
// 		return s.active()
// 	}
// 	select {
// 	case <-s.Context.Done():
// 		return false
// 	default:
// 		return s.active()
// 	}
// }

//go:inline
func (s *Executor) Op() byte {
	o := s.Emit[s.Head]
	s.Head++
	s.Gas.Spent++
	return o
}

//go:inline
func (s *Executor) Jmp(pc atf.PC) {
	s.Head = pc
}

//go:inline
func (s *Executor) PC() atf.PC {
	return s.Head
}

//go:inline
func (s *Executor) Fail(err error) {
	s.Err = err
}

//go:inline
func (s *Executor) Status() error {
	return s.Err
}

//go:inline
func (s *Executor) Ctx() context.Context {
	return s.Context
}

//go:inline
func (s *Executor) Cancel() {
	s.CancelFunc()
}

//go:inline
func (s *Executor) WithEmit(emit asf.Emitter) {
	s.Emit = emit
}

//go:inline
func (s *Executor) WithCtx(ctx context.Context, cancel context.CancelFunc) {
	s.Context, s.CancelFunc = ctx, cancel
}

//go:inline
func (s *Executor) WithTimeout(timeout time.Duration) {
	s.Context, s.CancelFunc = context.WithTimeout(s.Context, timeout)
}

//go:inline
func (s *Executor) WithDeadline(deadline time.Time) {
	s.Context, s.CancelFunc = context.WithDeadline(s.Context, deadline)
}

//go:inline
func (s *Executor) WithGas(spent, limit uint64) {
	s.Gas.Spent = spent
	s.Gas.Limit = limit
}
