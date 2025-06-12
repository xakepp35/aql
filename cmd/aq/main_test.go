package main_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/ast"
	"github.com/xakepp35/aql/pkg/vm"
)

const demoExprStr = "1+2*3"

var demoExpr = []byte(demoExprStr)

func BenchmarkFull(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = vm.Run(demoExprStr)
	}
}

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ast.Parse(demoExpr)
	}
}

func BenchmarkLex(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ast.Lex(demoExpr)
	}
}

func BenchmarkCompile(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = vm.Compile(demoExpr)
	}
}

func BenchmarkEmit(b *testing.B) {
	at, _ := ast.Parse(demoExpr)
	em := make(asf.Emitter, 0, 256)
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		em = em[:0]
		at.P0(&em)
		at.P1(&em)
		at.P2(&em)
	}
}

func BenchmarkMath(b *testing.B) {
	m := vm.NewSrc(demoExpr)
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Run()
		// if x := m.Pop(); x != int64(7) {
		// 	b.Fatalf("expected 7, got %v", x)
		// }
	}
}

func BenchmarkNop(b *testing.B) {
	m := vm.New()
	m.Emit = make(asf.Emitter, 256)
	b.ReportAllocs()
	b.SetBytes(256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Run()
	}
}
