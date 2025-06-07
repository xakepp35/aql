package fn

import (
	"errors"

	"github.com/xakepp35/aql/pkg/util"
)

var (
	ErrStackUnderflow   = errors.New("stack underflow")
	ErrStackUnsupported = errors.New("stack unsupported")
	ErrStackMissingPC   = errors.New("stack missing pc")
	ErrStackUniterable  = errors.New("stack uniterable")
	ErrDivisionByZero   = errors.New("division by zero")
	ErrModuloByZero     = errors.New("modulo by zero")
)

//go:inline
func StackUnsupported(args ...any) error {
	return util.EWrap(ErrStackUnsupported, util.ToTypes(args...))
}
