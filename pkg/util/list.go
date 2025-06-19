package util

import (
	"reflect"
)

//go:inline
func List[T any](args ...T) []any {
	a := make([]any, len(args))
	for i := range args {
		a[i] = args[i]
	}
	return a
}

//go:inline
func ToTypes(args ...any) []any {
	res := make([]any, len(args))
	for i := range args {
		res[i] = TypeOf(args[i])
	}
	return res
}

//go:inline
func TypeOf(v any) reflect.Type {
	return reflect.TypeOf(v)
}
