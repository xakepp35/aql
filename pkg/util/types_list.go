package util

import (
	"reflect"
)

func TypeOf(v any) reflect.Type {
	return reflect.TypeOf(v)
}

func ToTypes(args ...any) []any {
	res := make([]any, len(args))
	for i := range args {
		res[i] = TypeOf(args[i])
	}
	return res
}
