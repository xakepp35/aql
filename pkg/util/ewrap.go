package util

import "strings"

// Функция создания обёртки для ошибки
func EWrap(err error, args ...any) error {
	return &wrappedErr{err: err, args: args}
}

type wrappedErr struct {
	err  error
	args []any
}

func (e *wrappedErr) Error() string {
	var b strings.Builder
	b.WriteString(e.err.Error())
	if len(e.args) > 0 {
		b.WriteString(": ")
		for i := range e.args {
			b.WriteString(ToString(e.args[i]))
			b.WriteString(", ")
		}
	}
	return b.String()
}

func (e *wrappedErr) Unwrap() error { return e.err }
