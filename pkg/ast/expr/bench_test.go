package expr_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/ast/expr"
	"github.com/xakepp35/aql/pkg/vm/op"
)

func BenchmarkAlloc(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := expr.NewBinary(expr.NewDup(), expr.NewDup(), op.Add)
		_ = e
	}
}

func BenchmarkEmit(b *testing.B) {
	ex := expr.NewBinary(expr.NewDup(), expr.NewDup(), op.Add)
	e := make(asf.Emitter, 0, 128)
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e = e[:0] // reset the emitter
		ex.P0(&e)
		ex.P1(&e)
		ex.P2(&e)
	}
}
