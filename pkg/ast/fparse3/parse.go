// zero alloc ll parser
package fparse3

import (
	"errors"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/ast/asi"
	"github.com/xakepp35/aql/pkg/ast/expr"
)

/*──────────────────── PUBLIC ───────────────────*/
func ParseArena(src []byte, a *expr.Arena) (asi.AST, error) {
	p := parser{
		lex: lexer{s: src},
		a:   a,
	}
	p.next()
	return p.pipe()
}

func Parse(src []byte) (asi.AST, error) {
	p := parser{
		lex: lexer{s: src},
		a:   expr.NewArena(8),
	}
	p.next()
	return p.pipe()
}

/*──────────────────── PARSER ───────────────────*/

type parser struct {
	lex lexer
	a   *expr.Arena
}

//go:inline
func (p *parser) next() { p.lex.Next() }

//go:inline
func (p *parser) accept(k kind) bool {
	if p.lex.tok.k == k {
		p.next()
		return true
	}
	return false
}

//go:inline
func (p *parser) expect(k kind) {
	if !p.accept(k) {
		panic("syntax")
	}
}

/*──────── приоритетный каскад (LL) ────────*/

//go:inline
func (p *parser) pipe() (asi.AST, error) {
	l, _ := p.or()
	for p.accept('|') {
		r, _ := p.or()
		l = p.a.Binary(l, r, op.Pipe)
	}
	return l, nil
}

func (p *parser) or() (asi.AST, error) {
	l, _ := p.and()
	for p.lex.tok.k == tOrOr {
		p.next()
		r, _ := p.and()
		l = p.a.Binary(l, r, op.Or)
	}
	return l, nil
}

func (p *parser) and() (asi.AST, error) {
	l, _ := p.cmp()
	for p.lex.tok.k == tAndAnd {
		p.next()
		r, _ := p.cmp()
		l = p.a.Binary(l, r, op.And)
	}
	return l, nil
}

func (p *parser) cmp() (asi.AST, error) {
	l, _ := p.add()
	for {
		switch p.lex.tok.k {
		case tEq, tNeq, '<', tLe, '>', tGe:
			opc := [...]op.Code{tEq: op.Eq, tNeq: op.Neq, '<': op.Lt, tLe: op.Le, '>': op.Gt, tGe: op.Ge}[p.lex.tok.k]
			p.next()
			r, _ := p.add()
			l = p.a.Binary(l, r, opc)
		default:
			return l, nil
		}
	}
}

func (p *parser) add() (asi.AST, error) {
	l, _ := p.mul()
	for p.lex.tok.k == '+' || p.lex.tok.k == '-' {
		opc := op.Add
		if p.lex.tok.k == '-' {
			opc = op.Sub
		}
		p.next()
		r, _ := p.mul()
		l = p.a.Binary(l, r, opc)
	}
	return l, nil
}

func (p *parser) mul() (asi.AST, error) {
	l, _ := p.unary()
	for k := p.lex.tok.k; k == '*' || k == '/' || k == '%'; k = p.lex.tok.k {
		opc := [...]op.Code{'*': op.Mul, '/': op.Div, '%': op.Mod}[k]
		p.next()
		r, _ := p.unary()
		l = p.a.Binary(l, r, opc)
	}
	return l, nil
}

/*──── unary ────*/

func (p *parser) unary() (asi.AST, error) {
	switch p.lex.tok.k {
	case '-':
		p.next()
		r, _ := p.unary()
		return p.a.Unary(r, op.Not), nil
	case tOver:
		p.next()
		iter, _ := p.unary()
		if p.accept(tArrow) {
			p.expect('(')
			body, _ := p.pipe()
			p.expect(')')
			return p.a.Over(iter, body), nil
		}
		return p.a.Over(iter, nil), nil
	default:
		return p.post()
	}
}

/*──── postfix chain ────*/

func (p *parser) post() (asi.AST, error) {
	n, _ := p.atom()
	for {
		switch {
		case p.accept('.'):
			if p.lex.tok.k != tIdent {
				return nil, errors.New("ident expected")
			}
			name := p.name() // view — copy внутри emitter
			p.next()
			n = p.a.Field(n, name)

		case p.accept('['):
			idx1, _ := p.pipe()
			if p.accept(']') {
				n = p.a.Binary(n, idx1, op.Index1)
				continue
			}
			p.expect(':')
			idx2, _ := p.pipe()
			p.expect(']')
			n = p.a.Ternary(n, idx1, idx2, op.Index2)

		case p.lex.tok.k == tIdent && p.lex.s[p.lex.i] == '(':
			name := p.name()
			p.next() // идентификатор
			p.expect('(')
			var buf [4]asi.AST
			args := buf[:0]
			if !p.accept(')') {
				a, _ := p.pipe()
				args = append(args, a)
				for p.accept(',') {
					a, _ = p.pipe()
					args = append(args, a)
				}
				p.expect(')')
			}
			n = p.a.Call(append([]asi.AST(nil), args...), name)

		default:
			return n, nil
		}
	}
}

/*──── atom ────*/

func (p *parser) atom() (asi.AST, error) {
	switch p.lex.tok.k {
	case tIdent:
		name := p.name()
		p.next()
		return p.a.Ident(name), nil
	case tNum:
		v := int64(0)
		for _, d := range p.name() {
			v = v*10 + int64(d-'0')
		}
		p.next()
		return p.a.I64(v), nil
	case tStr:
		s := p.lex.s[p.lex.tok.off+1 : p.lex.tok.end-1]
		p.next()
		return p.a.StringBytes(s), nil
	case tTrue:
		p.next()
		return p.a.True(), nil
	case tFalse:
		p.next()
		return p.a.False(), nil
	case tNull:
		p.next()
		return p.a.Nil(), nil
	case '.':
		p.next()
		return p.a.Dup(), nil
	case '(':
		p.next()
		n, _ := p.pipe()
		p.expect(')')
		return n, nil
	default:
		return nil, errors.New("atom")
	}
}

//go:inline
func (p *parser) name() []byte {
	return p.lex.s[p.lex.tok.off:p.lex.tok.end]
}
