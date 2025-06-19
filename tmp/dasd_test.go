package tmp

import (
	"fmt"
	ast2 "github.com/xakepp35/aql/pkg/ast"
	"github.com/xakepp35/aql/pkg/require"
	"github.com/xakepp35/aql/tmp/ast"
	"github.com/xakepp35/aql/tmp/lexer"
	"github.com/xakepp35/aql/tmp/parser"
	"testing"
)

func TestName(t *testing.T) {
	l := lexer.NewLexer([]byte("22 + 11 + 10"))

	parse, err := parser.NewParser().Parse(l)
	require.NoError(t, err)

	fmt.Println(parse.(ast.Expr).Eval())

}

func BenchmarkParseSimpleArithmeticGocc(b *testing.B) {
	expr := []byte("1 + 2 + 3")
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	p := parser.NewParser()
	for i := 0; i < b.N; i++ {
		_, err := p.Parse(lexer.NewLexer(expr))
		if err != nil {
			b.Fatalf("eval error: %v", err)
		}
	}
}

func BenchmarkParseSimpleArithmeticGoYacc(b *testing.B) {
	expr := []byte("1 + 2 + 3")
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ast2.Parse(expr)
		if err != nil {
			b.Fatalf("eval error: %v", err)
		}
	}
}
