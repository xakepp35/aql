package lexer_test

import (
	"bytes"
	"testing"

	"github.com/xakepp35/aql/pkg/lexer"
)

// helper to collect kinds and literals for easier assertions
type pair struct {
	k   lexer.TokenKind
	lit string
}

func TestWholeTokenSet(t *testing.T) {
	src := ` true false null 123 0x2a 3.14 1e-9 "str" _id foo123 + - * / % . , : ; ( ) [ ] { } | || && == != < <= > >= => `

	// expected sequence (order mirrors src string above)
	want := []pair{
		{lexer.TTrue, "true"},
		{lexer.TFalse, "false"},
		{lexer.TNull, "null"},
		{lexer.TNumber, "123"},
		{lexer.TNumber, "0x2a"},
		{lexer.TFloat, "3.14"},
		{lexer.TFloat, "1e-9"},
		{lexer.TString, "\"str\""},
		{lexer.TIdentifier, "_id"},
		{lexer.TIdentifier, "foo123"},
		{lexer.TPlus, "+"},
		{lexer.TMinus, "-"},
		{lexer.TStar, "*"},
		{lexer.TSlash, "/"},
		{lexer.TPercent, "%"},
		{lexer.TDot, "."},
		{lexer.TComma, ","},
		{lexer.TColon, ":"},
		{lexer.TSemi, ";"},
		{lexer.TLParen, "("},
		{lexer.TRParen, ")"},
		{lexer.TLBracket, "["},
		{lexer.TRBracket, "]"},
		{lexer.TLBrace, "{"},
		{lexer.TRBrace, "}"},
		{lexer.TPipe, "|"},
		{lexer.TOr, "||"},
		{lexer.TAnd, "&&"},
		{lexer.TEq, "=="},
		{lexer.TNeq, "!="},
		{lexer.TLt, "<"},
		{lexer.TLte, "<="},
		{lexer.TGt, ">"},
		{lexer.TGte, ">="},
		{lexer.TArrow, "=>"},
	}

	lx := lexer.New([]byte(src))
	for i, exp := range want {
		tok := lx.Next()
		if tok.Kind != exp.k {
			t.Fatalf("idx %d: want kind %v got %v", i, exp.k, tok.Kind)
		}
		if !bytes.Equal(tok.Lit, []byte(exp.lit)) {
			t.Fatalf("idx %d: want lit %q got %q", i, exp.lit, string(tok.Lit))
		}
	}
	if tok := lx.Next(); tok.Kind != lexer.TEOF {
		t.Fatalf("expected EOF, got %v", tok.Kind)
	}
}

func BenchmarkLexerThroughput(b *testing.B) {
	data := []byte(`{"a":1, "b":2, "c":[1,2,3]} `)
	lx := lexer.New(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// restart each iter
		lx = lexer.New(data)
		for tok := lx.Next(); tok.Kind != lexer.TEOF; tok = lx.Next() {
		}
	}
}
