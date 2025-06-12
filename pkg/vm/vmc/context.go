package vmc

import "context"

type Context struct {
	context.Context
	context.CancelFunc
}

func (s *Context) Ctx() context.Context {
	return s.Context
}

func (s *Context) Cancel() {
	s.CancelFunc()
}

func (s *Context) SetCtx(ctx context.Context, cancel context.CancelFunc) {
	s.Context, s.CancelFunc = ctx, cancel
}
