package main_test

import (
	"context"
	"testing"

	"github.com/xakepp35/aql/pkg/aql"
	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/ast"
	"github.com/xakepp35/aql/pkg/ast/expr"
	"github.com/xakepp35/aql/pkg/ast/fparse2"
	"github.com/xakepp35/aql/pkg/ast/fparse3"
)

const demoExprStr = "1+2*3"

var demoExpr = []byte(demoExprStr)

func BenchmarkFull(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = aql.Run(demoExprStr)
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
		_, _ = aql.Compile(demoExpr)
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
	m := aql.NewSrc(demoExpr)
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Run()
	}
}

func BenchmarkNop(b *testing.B) {
	m := aql.New(context.Background(), nil)
	m.Executor.WithEmit(make(asf.Emitter, 256))
	b.ReportAllocs()
	b.SetBytes(256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Run()
	}
}

func BenchmarkFParse2(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fparse2.Parse(demoExpr)
	}
}

func BenchmarkFParse3(b *testing.B) {
	a := expr.NewArena(8)
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Reset()
		op, _ := fparse3.ParseArena(demoExpr, a)
		_ = op
	}
}

func BenchmarkFCompile(b *testing.B) {
	c, _ := fparse2.Parse(demoExpr)
	e := make(asf.Emitter, 256)
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e = e[:0]
		c(&e)
	}
}
