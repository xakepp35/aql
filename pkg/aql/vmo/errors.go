package vmo

import (
	"github.com/xakepp35/aql/pkg/util"
	"github.com/xakepp35/aql/pkg/vmi"
)

//go:inline
func StackUnderflow(args ...any) error {
	return util.EWrap(vmi.ErrStackUnderflow, args...)
}

//go:inline
func StackUnsupported(args ...any) error {
	return util.EWrap(vmi.ErrStackUnsupported, util.ToTypes(args...))
}

//go:inline
func IdentifierUndefined(name string) error {
	return util.EWrap(vmi.ErrIdentifierUndefined, name)
}
