package pratt

// import (
// 	"slices"

// 	"github.com/xakepp35/aql/pkg/ast/asi"
// 	"github.com/xakepp35/aql/pkg/cvt"
// )

// /* ---------- NU D ---------- */

// type NumNud struct{ b Builder }

// func (n NumNud) Parse(pr *Parser) (asi.AST, error) {
// 	v, _ := cvt.ParseInt64(pr.cur.lit)
// 	pr.next()
// 	return n.b.Literal(v), nil
// }

// type PrefixNud struct {
// 	b  Builder
// 	op string
// }

// func (n PrefixNud) Parse(pr *Parser) (asi.AST, error) {
// 	pr.next()                 // съели оператор
// 	right, err := pr.expr(90) // высокий rbp
// 	if err != nil {
// 		return nil, err
// 	}
// 	return n.b.Unary(right, n.op), nil
// }

/* ---------- LE D ---------- */

// type BinaryLed struct {
// 	b   Builder
// 	op  string
// 	lbp int
// }

// func (l BinaryLed) bp() int { return l.lbp }
// func (l BinaryLed) Parse(pr *Parser, left asi.AST) (asi.AST, error) {
// 	pr.next()
// 	right, err := pr.expr(l.lbp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return l.b.Binary(left, right, l.op), nil
// }

// type FieldLed struct{ b Builder }

// func (l FieldLed) bp() int { return 200 } // высокий
// func (l FieldLed) Parse(pr *Parser, base asi.AST) (asi.AST, error) {
// 	pr.next()
// 	if pr.cur.kind != pr.id("IDENT") {
// 		return nil, pr.err("ident exp")
// 	}
// 	name := slices.Clone(pr.cur.lit)
// 	pr.next()
// 	return l.b.Field(base, name), nil
// }
