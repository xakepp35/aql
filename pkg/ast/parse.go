//go:generate goyacc -o parser_gen.go -p aql aql.y
package ast

import (
	"fmt"

	"github.com/xakepp35/aql/pkg/ast/asi"
	"github.com/xakepp35/aql/pkg/ast/expr"
	"github.com/xakepp35/aql/pkg/lexer"
)

var Arn = expr.NewArena(8)

func Lex(src []byte) error {
	lx := lexer.New(src)
	var t lexer.Token
	for t.Kind != lexer.TEOF {
		t = lx.Next()
	}
	return nil
}

// Parse разбирает выражение и возвращает AST.
func Parse(src []byte) (asi.AST, error) {
	// aqlDebug = 4
	// aqlErrorVerbose = true
	Arn.Reset()
	b := bridge{lx: lexer.New(src), Arena: Arn}

	p := aqlNewParser()

	if p.Parse(&b) != 0 || len(b.errs) > 0 {
		return nil, fmt.Errorf("%v", b.errs)
	}
	return b.result, nil
}

type bridge struct {
	lx *lexer.Lexer
	*expr.Arena
	result asi.AST
	errs   []string
}

func (b *bridge) Lex(lval *aqlSymType) int {
	t := b.lx.Next()
	switch t.Kind {
	// case lexer.TEOF:
	// 	return EOF
	case lexer.TIdentifier:
		lval.b = t.Lit
		return IDENT
	case lexer.TNumber:
		lval.b = t.Lit
		return NUMBER
	case lexer.TString:
		lval.b = t.Lit
		return STRING
	case lexer.TTrue:
		return TRUE
	case lexer.TFalse:
		return FALSE
	case lexer.TNull:
		return NULL
	case lexer.TPlus:
		return PLUS
	case lexer.TMinus:
		return MINUS
	case lexer.TStar:
		return STAR
	case lexer.TSlash:
		return SLASH
	case lexer.TPercent:
		return PERCENT
	case lexer.TPipe:
		return PIPE
	case lexer.TAnd:
		return ANDAND
	case lexer.TOr:
		return OROR
	case lexer.TEq:
		return EQ
	case lexer.TNeq:
		return NEQ
	case lexer.TLt:
		return LT
	case lexer.TLte:
		return LE
	case lexer.TGt:
		return GT
	case lexer.TGte:
		return GE

	case lexer.TDot:
		return DOT
	case lexer.TLBracket:
		return LBRACK
	case lexer.TRBracket:
		return RBRACK
	case lexer.TLParen:
		return LPAREN
	case lexer.TRParen:
		return RPAREN
	case lexer.TColon:
		return COLON
	case lexer.TComma:
		return COMMA
	case lexer.TOver:
		return OVER
	case lexer.TArrow:
		return ARROW
	default:
		return 0
	}
}

func (b *bridge) Error(s string) { b.errs = append(b.errs, s) }
