package ast_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/ast"
	"github.com/xakepp35/aql/pkg/vm"
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

func BenchmarkEvalSimpleArithmetic(b *testing.B) {
	expr := []byte("1 + 2 * 3")

	node, err := ast.Parse(expr)
	if err != nil {
		b.Fatalf("parse error: %v", err)
	}

	this := vm.NewState()
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := node.Eval(this)
		if err != nil {
			b.Fatalf("eval error: %v", err)
		}
	}
}
