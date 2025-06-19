package fparse3_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/ast/fparse3"
)

func BenchmarkLexer(b *testing.B) {
	var src = []byte("1 + 2 * 3")
	lexer := fparse3.NewLexer(src)
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lexer.Reset()
		lexer.Drain()
	}
}
