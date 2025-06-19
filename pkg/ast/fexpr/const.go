package fexpr

import "github.com/xakepp35/aql/pkg/asf"

// go:inline
func Nil(e *asf.Emitter) error { e.Nil(); return nil }

// go:inline
func False(e *asf.Emitter) error { e.I64(0); return nil }

// go:inline
func True(e *asf.Emitter) error { e.I64(1); return nil }

func StringBytes(s []byte) Compiler {
	return func(e *asf.Emitter) error {
		e.StringBytes(s)
		return nil
	}
}

// go:inline
func I64(v int64) Compiler {
	return nil
	// return func(e *asf.Emitter) error {
	// 	e.I64(v)
	// 	return nil
	// }
}

// go:inline
func F64(v float64) Compiler {
	return func(e *asf.Emitter) error {
		e.F64(v)
		return nil
	}
}
