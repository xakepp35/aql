// zero alloc ll parser
package fparse2

import (
	"errors"

	"github.com/xakepp35/aql/pkg/aql/op"
	"github.com/xakepp35/aql/pkg/ast/fexpr"
)

/*──────────────────── PUBLIC ───────────────────*/

func Parse(src []byte) (fexpr.Compiler, error) {
	p := parser{lex: lexer{s: src}}
	p.next()
	return p.pipe()
}

/*──────────────────── PARSER ───────────────────*/

type parser struct {
	lex lexer
}

//go:inline
func (p *parser) next() { p.lex.next() }

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
func (p *parser) pipe() (fexpr.Compiler, error) {
	l, _ := p.or()
	for p.accept('|') {
		r, _ := p.or()
		l = fexpr.Binary(l, r, op.Pipe)
	}
	return l, nil
}

//go:inline
func (p *parser) or() (fexpr.Compiler, error) {
	l, _ := p.and()
	for p.lex.tok.k == tOrOr {
		p.next()
		r, _ := p.and()
		l = fexpr.Binary(l, r, op.Or)
	}
	return l, nil
}

//go:inline
func (p *parser) and() (fexpr.Compiler, error) {
	l, _ := p.cmp()
	for p.lex.tok.k == tAndAnd {
		p.next()
		r, _ := p.cmp()
		l = fexpr.Binary(l, r, op.And)
	}
	return l, nil
}

//go:inline
func (p *parser) cmp() (fexpr.Compiler, error) {
	l, _ := p.add()
	for {
		switch p.lex.tok.k {
		case tEq, tNeq, '<', tLe, '>', tGe:
			opc := [...]op.Code{tEq: op.Eq, tNeq: op.Neq, '<': op.Lt, tLe: op.Le, '>': op.Gt, tGe: op.Ge}[p.lex.tok.k]
			p.next()
			r, _ := p.add()
			l = fexpr.Binary(l, r, opc)
		default:
			return l, nil
		}
	}
}

//go:inline
func (p *parser) add() (fexpr.Compiler, error) {
	l, _ := p.mul()
	for p.lex.tok.k == '+' || p.lex.tok.k == '-' {
		opc := op.Add
		if p.lex.tok.k == '-' {
			opc = op.Sub
		}
		p.next()
		r, _ := p.mul()
		l = fexpr.Binary(l, r, opc)
	}
	return l, nil
}

//go:inline
func (p *parser) mul() (fexpr.Compiler, error) {
	l, _ := p.unary()
	for k := p.lex.tok.k; k == '*' || k == '/' || k == '%'; k = p.lex.tok.k {
		opc := [...]op.Code{'*': op.Mul, '/': op.Div, '%': op.Mod}[k]
		p.next()
		r, _ := p.unary()
		l = fexpr.Binary(l, r, opc)
	}
	return l, nil
}

/*──── unary ────*/
//go:inline
func (p *parser) unary() (fexpr.Compiler, error) {
	switch p.lex.tok.k {
	case '-':
		p.next()
		r, _ := p.unary()
		return fexpr.Unary(r, op.Not), nil
	case tOver:
		p.next()
		iter, _ := p.unary()
		if p.accept(tArrow) {
			p.expect('(')
			body, _ := p.pipe()
			p.expect(')')
			return fexpr.Over(iter, body), nil
		}
		return fexpr.Over(iter, fexpr.Nop), nil
	default:
		return p.post()
	}
}

//go:inline
func (p *parser) name() []byte {
	return p.lex.s[p.lex.tok.off:p.lex.tok.end]
}

/*──── postfix chain ────*/
//go:inline
func (p *parser) post() (fexpr.Compiler, error) {
	n, _ := p.atom()
	for {
		switch {
		case p.accept('.'):
			if p.lex.tok.k != tIdent {
				return nil, errors.New("ident expected")
			}
			name := p.name() // view — copy внутри emitter
			p.next()
			n = fexpr.Field(n, name)

		case p.accept('['):
			idx1, _ := p.pipe()
			if p.accept(']') {
				n = fexpr.Binary(n, idx1, op.Index1)
				continue
			}
			p.expect(':')
			idx2, _ := p.pipe()
			p.expect(']')
			n = fexpr.Ternary(n, idx1, idx2, op.Index2)

		case p.lex.tok.k == tIdent && p.lex.s[p.lex.i] == '(':
			name := p.name()
			p.next() // идентификатор
			p.expect('(')
			var buf [4]fexpr.Compiler
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
			n = fexpr.Call(append([]fexpr.Compiler(nil), args...), name)

		default:
			return n, nil
		}
	}
}

/*──── atom ────*/
//go:inline
func (p *parser) atom() (fexpr.Compiler, error) {
	switch p.lex.tok.k {
	case tIdent:
		name := p.name()
		p.next()
		return fexpr.Ident(name), nil
	case tNum:
		v := int64(0)
		for _, d := range p.name() {
			v = v*10 + int64(d-'0')
		}
		p.next()
		return fexpr.I64(v), nil
	case tStr:
		s := p.lex.s[p.lex.tok.off+1 : p.lex.tok.end-1]
		p.next()
		return fexpr.StringBytes(s), nil
	case tTrue:
		p.next()
		return fexpr.True, nil
	case tFalse:
		p.next()
		return fexpr.False, nil
	case tNull:
		p.next()
		return fexpr.Nil, nil
	case '.':
		p.next()
		return fexpr.Dup, nil
	case '(':
		p.next()
		n, _ := p.pipe()
		p.expect(')')
		return n, nil
	default:
		return nil, errors.New("atom")
	}
}
