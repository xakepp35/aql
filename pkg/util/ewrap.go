package util

import "strings"

// Функция создания обёртки для ошибки
//go:inline
func EWrap(err error, args ...any) error {
	return &wrappedErr{err: err, args: args}
}

type wrappedErr struct {
	err  error
	args []any
}

//go:inline
func (e *wrappedErr) Error() string {
	var b strings.Builder
	b.WriteString(e.err.Error())
	b.WriteString(": [")
	if len(e.args) > 0 {
		for i := range e.args {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(ToString(e.args[i]))
		}
	}
	b.WriteByte(']')
	return b.String()
}

//go:inline
func (e *wrappedErr) Unwrap() error { return e.err }
