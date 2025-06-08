package vmi

import "errors"

var (
	ErrCompile = errors.New("compile")
)

// run errors
var (
	ErrRuntime       = errors.New("runtime")
	ErrHalted        = errors.New("halted")
	ErrFinished      = errors.New("finished")
	ErrUnknownType   = errors.New("unknown type")
	ErrUnimplemented = errors.New("unimplemented")
)

// stack errors
var (
	ErrStackUnderflow   = errors.New("stack underflow")
	ErrStackUnsupported = errors.New("stack unsupported")
	ErrStackWrongArg    = errors.New("stack wrong arg")
	ErrStackUniterable  = errors.New("stack uniterable")
)

// functional errors
var (
	ErrDivisionByZero    = errors.New("division by zero")
	ErrModuloByZero      = errors.New("modulo by zero")
	ErrFunctionUndefined = errors.New("function undefined")
	ErrVariableUndefined = errors.New("variable undefined")
)
