package ast_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/ast"
)

func BenchmarkParseSimpleArithmetic(b *testing.B) {
	expr := []byte("1 + 2 * 3")
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ast.Parse(expr)
		if err != nil {
			b.Fatalf("eval error: %v", err)
		}
	}
}
