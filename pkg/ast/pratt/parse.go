package pratt

import (
	"errors"
	"strconv"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/ast/fexpr"
	"github.com/xakepp35/aql/pkg/util"
)

// Parse(src) → root-Compiler
func Parse(src []byte) (fexpr.Compiler, error) {
	lex := newLexer(src)
	p := parser{lex: lex}
	p.next() // первичный токен
	return p.expression(0)
}

type parser struct {
	lex *lexer
	cur token
}

func (p *parser) next() { p.cur = p.lex.nextToken() }

func (p *parser) accept(kind tokKind) bool {
	if p.cur.kind == kind {
		p.next()
		return true
	}
	return false
}

func (p *parser) err(msg string, args ...any) error { return util.EWrap(errors.New(msg), args...) }

/*──── Pratt table helpers ────*/

func bp(tok tokKind) int { // left binding power
	switch tok {
	case '|':
		return 5
	case tOrOr:
		return 10
	case tAndAnd:
		return 15
	case tEq, tNeq, '<', tLe, '>', tGe:
		return 20
	case '+', '-':
		return 30
	case '*', '/', '%':
		return 40
	case '.', '[', '(':
		return 50
	default:
		return 0
	}
}

/*──────────────  expression(rbp)  ───────────*/

func (p *parser) expression(rbp int) (fexpr.Compiler, error) {
	// nud ― prefix / literal / ident / "("
	left, err := p.nud()
	if err != nil {
		return nil, err
	}
	for rbp < bp(p.cur.kind) {
		led, err := p.led(left)
		if err != nil {
			return nil, err
		}
		left = led
	}
	return left, nil
}

/*───────────────  NUD  (prefix) ─────────────*/

func (p *parser) nud() (fexpr.Compiler, error) {
	switch p.cur.kind {
	case tNumber:
		v, _ := strconv.ParseInt(string(p.cur.lit), 10, 64)
		p.next()
		return fexpr.I64(v), nil

	case tString:
		s := p.cur.lit[1 : len(p.cur.lit)-1]
		p.next()
		return fexpr.StringBytes(s), nil

	case tTrue:
		p.next()
		return fexpr.True, nil

	case tFalse:
		return fexpr.False, nil
	case tNull:
		return fexpr.Nil, nil

	case '-': // prefix minus = op.Not
		p.next()
		right, _ := p.expression(90)
		return fexpr.Unary(right, op.Not), nil

	case '(': // group
		p.next()
		n, err := p.expression(0)
		if err != nil {
			return nil, err
		}
		if !p.accept(tokKind(')')) {
			return nil, p.err("')' expected")
		}
		return n, nil

	case tIdent:
		id := p.cur.lit
		p.next()
		return fexpr.Ident(id), nil // начальный идентификатор (можно .field / call)
	case '.':
		p.next()
		return fexpr.Dup, nil
	default:
		return nil, p.err("unexpected token in nud")
	}
}

/*───────────────  LED  (infix / postfix) ─────────────*/

func (p *parser) led(left fexpr.Compiler) (fexpr.Compiler, error) {
	switch p.cur.kind {

	/*──── бинарные ────*/
	case '+', '-', '*', '/', '%',
		'|', tAndAnd, tOrOr,
		tEq, tNeq, '<', tLe, '>', tGe:
		opcode := map[tokKind]op.Code{
			'+': op.Add, '-': op.Sub, '*': op.Mul, '/': op.Div, '%': op.Mod,
			'|': op.Pipe, tAndAnd: op.And, tOrOr: op.Or,
			tEq: op.Eq, tNeq: op.Neq, '<': op.Lt, tLe: op.Le, '>': op.Gt, tGe: op.Ge,
		}[p.cur.kind]
		bpower := bp(p.cur.kind)
		p.next()
		right, err := p.expression(bpower)
		if err != nil {
			return nil, err
		}
		return fexpr.Binary(left, right, opcode), nil

	/*──── field  left . ident ────*/
	case '.':
		p.next()
		if p.cur.kind != tIdent {
			return nil, p.err("ident expected")
		}
		name := p.cur.lit
		p.next()
		return fexpr.Field(left, name), nil

	/*──── call  ident(…) ────*/
	case tokKind('('):
		// left _must_ be identCompiler
		name := p.cur.lit
		p.next() // '('
		var args []fexpr.Compiler
		if !p.accept(tokKind(')')) {
			arg, _ := p.expression(0)
			args = append(args, arg)
			for p.accept(tokKind(',')) {
				arg, _ = p.expression(0)
				args = append(args, arg)
			}
			if !p.accept(tokKind(')')) {
				return nil, p.err("')' expected")
			}
		}
		return fexpr.Call(args, name), nil

	/*──── index / slice ────*/
	case tokKind('['):
		p.next()
		a, _ := p.expression(0)
		if p.accept(tokKind(']')) { // single index
			return fexpr.Binary(left, a, op.Index1), nil
		}
		if !p.accept(tokKind(':')) {
			return nil, p.err("':' exp")
		}
		b, _ := p.expression(0)
		if !p.accept(tokKind(']')) {
			return nil, p.err("']'")
		}
		return fexpr.Ternary(left, a, b, op.Index2), nil
	}
	return nil, p.err("led: no match")
}
