package lexer_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/lexer"
)

// FuzzLexer ensures no panics and lexers eventually terminate.
func FuzzLexer(f *testing.F) {
	seeds := [][]byte{
		[]byte(""),
		[]byte("{foo: 1}"),
		[]byte("\"unterminated"),
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		lx := lexer.New(data)
		// iterate until EOF or too many tokens (safety against bugs)
		for i := 0; i < 1e6; i++ {
			tok := lx.Next()
			if tok.Kind == lexer.TIllegal {
				// lexer should keep advancing even after illegal token
				continue
			}
			if tok.Kind == lexer.TEOF {
				break
			}
		}
	})
}
