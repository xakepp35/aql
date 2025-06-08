package fn

import (
	"errors"

	"github.com/xakepp35/aql/pkg/util"
)

func StackUnderflow(args ...any) error {
	return util.EWrap(ErrStackUnderflow, args...)
}

//go:inline
func StackUnsupported(args ...any) error {
	return util.EWrap(ErrStackUnsupported, util.ToTypes(args...))
}

func FunctionUndefined(name string) error {
	return util.EWrap(ErrFunctionUndefined, name)
}

func VariableUndefined(name string) error {
	return util.EWrap(ErrVariableUndefined, name)
}
