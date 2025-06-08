//go:generate goyacc -o parser_gen.go -p aql aql.y
package aqc

import (
	"strconv"

	"github.com/xakepp35/aql/pkg/asf/atf"
	"github.com/xakepp35/aql/pkg/lexer"
	"github.com/xakepp35/aql/pkg/util"
	"github.com/xakepp35/aql/pkg/vmi"
)

// Compile AnyQueryLanguage bytecode compiler (see vm.Program for emitter)
func Compile(src []byte, e atf.Emitter) error {
	// aqlDebug = 4
	// aqlErrorVerbose = true
	c := &Compiler{lx: lexer.New(src), Emitter: e}
	if aqlParse(c) != 0 || len(c.errs) > 0 {
		return util.EWrap(vmi.ErrCompile, util.List(c.errs...)...)
	}
	return nil
}

type Compiler struct {
	atf.Emitter
	lx   *lexer.Lexer
	errs []string
}

func (b *Compiler) Lex(lval *aqlSymType) int {
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

func (c *Compiler) Error(s string) { c.errs = append(c.errs, s) }

func parseInt(b []byte) int64 {
	res, _ := strconv.ParseInt(string(b), 10, 64)
	return res
}
