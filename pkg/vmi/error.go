package vmi

import "errors"

var (
	ErrCompile       = errors.New("compile")
	ErrHalted        = errors.New("halted")
	ErrUnimplemented = errors.New("unimplemented")
)
