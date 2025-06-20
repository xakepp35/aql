// Code generated by goyacc -o parser_gen.go -p aql aql.y. DO NOT EDIT.

//line aql.y:2
package ast

import __yyfmt__ "fmt"

//line aql.y:2

import (
	"github.com/xakepp35/aql/pkg/ast/asi"
	"github.com/xakepp35/aql/pkg/cvt"
	"github.com/xakepp35/aql/pkg/aql/op"
)

//line aql.y:12
type aqlSymType struct {
	yys int
	b   []byte /* сырые лексемы-строки от лексера */
	i   int64  /* счётчики (например, кол-во арг-ов) */

	n asi.AST   /* любой узел AST */
	a []asi.AST /* срез узлов (список арг-ов) */
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const TRUE = 57349
const FALSE = 57350
const NULL = 57351
const PLUS = 57352
const MINUS = 57353
const STAR = 57354
const SLASH = 57355
const PERCENT = 57356
const PIPE = 57357
const ANDAND = 57358
const OROR = 57359
const EQ = 57360
const NEQ = 57361
const LT = 57362
const LE = 57363
const GT = 57364
const GE = 57365
const DOT = 57366
const LBRACK = 57367
const RBRACK = 57368
const LPAREN = 57369
const RPAREN = 57370
const COLON = 57371
const COMMA = 57372
const OVER = 57373
const ARROW = 57374
const UMINUS = 57375

var aqlToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"STRING",
	"TRUE",
	"FALSE",
	"NULL",
	"PLUS",
	"MINUS",
	"STAR",
	"SLASH",
	"PERCENT",
	"PIPE",
	"ANDAND",
	"OROR",
	"EQ",
	"NEQ",
	"LT",
	"LE",
	"GT",
	"GE",
	"DOT",
	"LBRACK",
	"RBRACK",
	"LPAREN",
	"RPAREN",
	"COLON",
	"COMMA",
	"OVER",
	"ARROW",
	"UMINUS",
}

var aqlStatenames = [...]string{}

const aqlEofCode = 1
const aqlErrCode = 2
const aqlInitialStackSize = 16

//line aql.y:137

//line yacctab:1
var aqlExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const aqlPrivate = 57344

const aqlLast = 85

var aqlAct = [...]int8{
	2, 14, 15, 16, 17, 18, 19, 8, 11, 14,
	15, 16, 17, 18, 19, 58, 11, 7, 9, 72,
	63, 20, 41, 64, 21, 59, 62, 65, 12, 20,
	38, 39, 21, 40, 71, 66, 12, 67, 57, 51,
	52, 61, 6, 45, 46, 47, 48, 49, 50, 36,
	37, 5, 53, 54, 55, 25, 26, 27, 28, 29,
	30, 4, 23, 24, 22, 68, 69, 44, 70, 33,
	34, 35, 31, 32, 56, 43, 60, 13, 10, 3,
	1, 0, 0, 0, 42,
}

var aqlPact = [...]int16{
	5, -1000, -1000, 49, 45, 47, 37, 62, 57, -1000,
	25, 5, 5, -1000, 6, -1000, -1000, -1000, -1000, -1000,
	-1000, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 70, 5, -1000, -17,
	-3, -2, 45, 47, 37, 62, 62, 62, 62, 62,
	62, 57, 57, -1000, -1000, -1000, -1000, -6, 0, -1000,
	7, -1000, -1000, -1000, 5, 5, -1000, 5, 8, -9,
	-1000, -1000, -1000,
}

var aqlPgo = [...]int8{
	0, 80, 0, 79, 61, 51, 42, 17, 7, 18,
	78, 77, 76,
}

var aqlR1 = [...]int8{
	0, 1, 2, 3, 3, 4, 4, 5, 5, 6,
	6, 6, 6, 6, 6, 6, 7, 7, 7, 8,
	8, 8, 8, 9, 9, 9, 9, 10, 10, 10,
	10, 10, 10, 12, 12, 11, 11, 11, 11, 11,
	11, 11, 11,
}

var aqlR2 = [...]int8{
	0, 1, 1, 1, 3, 1, 3, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 1, 3, 3, 1,
	3, 3, 3, 1, 2, 2, 6, 1, 3, 4,
	6, 3, 4, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 3,
}

var aqlChk = [...]int16{
	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, 11, 31, -11, 4, 5, 6, 7, 8, 9,
	24, 27, 15, 17, 16, 18, 19, 20, 21, 22,
	23, 10, 11, 12, 13, 14, 24, 25, -9, -9,
	27, -2, -4, -5, -6, -7, -7, -7, -7, -7,
	-7, -8, -8, -9, -9, -9, 4, -2, 32, 28,
	-12, -2, 28, 26, 29, 27, 28, 30, -2, -2,
	-2, 26, 28,
}

var aqlDef = [...]int8{
	0, -2, 1, 2, 3, 5, 7, 9, 16, 19,
	23, 0, 0, 27, 35, 36, 37, 38, 39, 40,
	41, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 24, 25,
	0, 0, 4, 6, 8, 10, 11, 12, 13, 14,
	15, 17, 18, 20, 21, 22, 28, 0, 0, 31,
	0, 33, 42, 29, 0, 0, 32, 0, 0, 0,
	34, 30, 26,
}

var aqlTok1 = [...]int8{
	1,
}

var aqlTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33,
}

var aqlTok3 = [...]int8{
	0,
}

var aqlErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	aqlDebug        = 0
	aqlErrorVerbose = false
)

type aqlLexer interface {
	Lex(lval *aqlSymType) int
	Error(s string)
}

type aqlParser interface {
	Parse(aqlLexer) int
	Lookahead() int
}

type aqlParserImpl struct {
	lval  aqlSymType
	stack [aqlInitialStackSize]aqlSymType
	char  int
}

func (p *aqlParserImpl) Lookahead() int {
	return p.char
}

func aqlNewParser() aqlParser {
	return &aqlParserImpl{}
}

const aqlFlag = -1000

func aqlTokname(c int) string {
	if c >= 1 && c-1 < len(aqlToknames) {
		if aqlToknames[c-1] != "" {
			return aqlToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func aqlStatname(s int) string {
	if s >= 0 && s < len(aqlStatenames) {
		if aqlStatenames[s] != "" {
			return aqlStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func aqlErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !aqlErrorVerbose {
		return "syntax error"
	}

	for _, e := range aqlErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + aqlTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(aqlPact[state])
	for tok := TOKSTART; tok-1 < len(aqlToknames); tok++ {
		if n := base + tok; n >= 0 && n < aqlLast && int(aqlChk[int(aqlAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if aqlDef[state] == -2 {
		i := 0
		for aqlExca[i] != -1 || int(aqlExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; aqlExca[i] >= 0; i += 2 {
			tok := int(aqlExca[i])
			if tok < TOKSTART || aqlExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if aqlExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += aqlTokname(tok)
	}
	return res
}

func aqllex1(lex aqlLexer, lval *aqlSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(aqlTok1[0])
		goto out
	}
	if char < len(aqlTok1) {
		token = int(aqlTok1[char])
		goto out
	}
	if char >= aqlPrivate {
		if char < aqlPrivate+len(aqlTok2) {
			token = int(aqlTok2[char-aqlPrivate])
			goto out
		}
	}
	for i := 0; i < len(aqlTok3); i += 2 {
		token = int(aqlTok3[i+0])
		if token == char {
			token = int(aqlTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(aqlTok2[1]) /* unknown char */
	}
	if aqlDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", aqlTokname(token), uint(char))
	}
	return char, token
}

func aqlParse(aqllex aqlLexer) int {
	return aqlNewParser().Parse(aqllex)
}

func (aqlrcvr *aqlParserImpl) Parse(aqllex aqlLexer) int {
	var aqln int
	var aqlVAL aqlSymType
	var aqlDollar []aqlSymType
	_ = aqlDollar // silence set and not used
	aqlS := aqlrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	aqlstate := 0
	aqlrcvr.char = -1
	aqltoken := -1 // aqlrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		aqlstate = -1
		aqlrcvr.char = -1
		aqltoken = -1
	}()
	aqlp := -1
	goto aqlstack

ret0:
	return 0

ret1:
	return 1

aqlstack:
	/* put a state and value onto the stack */
	if aqlDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", aqlTokname(aqltoken), aqlStatname(aqlstate))
	}

	aqlp++
	if aqlp >= len(aqlS) {
		nyys := make([]aqlSymType, len(aqlS)*2)
		copy(nyys, aqlS)
		aqlS = nyys
	}
	aqlS[aqlp] = aqlVAL
	aqlS[aqlp].yys = aqlstate

aqlnewstate:
	aqln = int(aqlPact[aqlstate])
	if aqln <= aqlFlag {
		goto aqldefault /* simple state */
	}
	if aqlrcvr.char < 0 {
		aqlrcvr.char, aqltoken = aqllex1(aqllex, &aqlrcvr.lval)
	}
	aqln += aqltoken
	if aqln < 0 || aqln >= aqlLast {
		goto aqldefault
	}
	aqln = int(aqlAct[aqln])
	if int(aqlChk[aqln]) == aqltoken { /* valid shift */
		aqlrcvr.char = -1
		aqltoken = -1
		aqlVAL = aqlrcvr.lval
		aqlstate = aqln
		if Errflag > 0 {
			Errflag--
		}
		goto aqlstack
	}

aqldefault:
	/* default state action */
	aqln = int(aqlDef[aqlstate])
	if aqln == -2 {
		if aqlrcvr.char < 0 {
			aqlrcvr.char, aqltoken = aqllex1(aqllex, &aqlrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if aqlExca[xi+0] == -1 && int(aqlExca[xi+1]) == aqlstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			aqln = int(aqlExca[xi+0])
			if aqln < 0 || aqln == aqltoken {
				break
			}
		}
		aqln = int(aqlExca[xi+1])
		if aqln < 0 {
			goto ret0
		}
	}
	if aqln == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			aqllex.Error(aqlErrorMessage(aqlstate, aqltoken))
			Nerrs++
			if aqlDebug >= 1 {
				__yyfmt__.Printf("%s", aqlStatname(aqlstate))
				__yyfmt__.Printf(" saw %s\n", aqlTokname(aqltoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for aqlp >= 0 {
				aqln = int(aqlPact[aqlS[aqlp].yys]) + aqlErrCode
				if aqln >= 0 && aqln < aqlLast {
					aqlstate = int(aqlAct[aqln]) /* simulate a shift of "error" */
					if int(aqlChk[aqlstate]) == aqlErrCode {
						goto aqlstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if aqlDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", aqlS[aqlp].yys)
				}
				aqlp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if aqlDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", aqlTokname(aqltoken))
			}
			if aqltoken == aqlEofCode {
				goto ret1
			}
			aqlrcvr.char = -1
			aqltoken = -1
			goto aqlnewstate /* try again in the same state */
		}
	}

	/* reduction by production aqln */
	if aqlDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", aqln, aqlStatname(aqlstate))
	}

	aqlnt := aqln
	aqlpt := aqlp
	_ = aqlpt // guard against "declared and not used"

	aqlp -= int(aqlR2[aqln])
	// aqlp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if aqlp+1 >= len(aqlS) {
		nyys := make([]aqlSymType, len(aqlS)*2)
		copy(nyys, aqlS)
		aqlS = nyys
	}
	aqlVAL = aqlS[aqlp+1]

	/* consult goto table to find next state */
	aqln = int(aqlR1[aqln])
	aqlg := int(aqlPgo[aqln])
	aqlj := aqlg + aqlS[aqlp].yys + 1

	if aqlj >= aqlLast {
		aqlstate = int(aqlAct[aqlg])
	} else {
		aqlstate = int(aqlAct[aqlj])
		if int(aqlChk[aqlstate]) != -aqln {
			aqlstate = int(aqlAct[aqlg])
		}
	}
	// dummy call; replaced with literal code
	switch aqlnt {

	case 1:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:46
		{
			aqllex.(*bridge).result = aqlDollar[1].n
		}
	case 2:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:51
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 3:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:56
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 4:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:57
		{
			aqlVAL.n = aqllex.(*bridge).Pipe(aqlDollar[1].n, aqlDollar[3].n)
		}
	case 5:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:63
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 6:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:64
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Or)
		}
	case 7:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:69
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 8:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:70
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.And)
		}
	case 9:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:75
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 10:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:76
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Eq)
		}
	case 11:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:77
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Neq)
		}
	case 12:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:78
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Lt)
		}
	case 13:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:79
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Le)
		}
	case 14:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:80
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Gt)
		}
	case 15:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:81
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Ge)
		}
	case 16:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:86
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 17:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:87
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Add)
		}
	case 18:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:88
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Sub)
		}
	case 19:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:93
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 20:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:94
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Mul)
		}
	case 21:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:95
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Div)
		}
	case 22:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:96
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Mod)
		}
	case 23:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:101
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 24:
		aqlDollar = aqlS[aqlpt-2 : aqlpt+1]
//line aql.y:102
		{
			aqlVAL.n = aqllex.(*bridge).Unary(aqlDollar[2].n, op.Not)
		}
	case 25:
		aqlDollar = aqlS[aqlpt-2 : aqlpt+1]
//line aql.y:103
		{
			aqlVAL.n = aqllex.(*bridge).Over(aqlDollar[2].n, nil)
		}
	case 26:
		aqlDollar = aqlS[aqlpt-6 : aqlpt+1]
//line aql.y:104
		{
			aqlVAL.n = aqllex.(*bridge).Over(aqlDollar[2].n, aqlDollar[5].n)
		}
	case 27:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:109
		{
			aqlVAL.n = aqlDollar[1].n
		}
	case 28:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:110
		{
			aqlVAL.n = aqllex.(*bridge).Field(aqlDollar[1].n, aqlDollar[3].b)
		}
	case 29:
		aqlDollar = aqlS[aqlpt-4 : aqlpt+1]
//line aql.y:111
		{
			aqlVAL.n = aqllex.(*bridge).Binary(aqlDollar[1].n, aqlDollar[3].n, op.Index1)
		}
	case 30:
		aqlDollar = aqlS[aqlpt-6 : aqlpt+1]
//line aql.y:112
		{
			aqlVAL.n = aqllex.(*bridge).Ternary(aqlDollar[1].n, aqlDollar[3].n, aqlDollar[5].n, op.Index2)
		}
	case 31:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:113
		{
			aqlVAL.n = aqllex.(*bridge).Call(nil, aqlDollar[1].b)
		}
	case 32:
		aqlDollar = aqlS[aqlpt-4 : aqlpt+1]
//line aql.y:114
		{
			aqlVAL.n = aqllex.(*bridge).Call(aqlDollar[3].a, aqlDollar[1].b)
		}
	case 33:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:119
		{
			aqlVAL.a = []asi.AST{aqlDollar[1].n}
		}
	case 34:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:120
		{
			aqlVAL.a = append(aqlDollar[1].a, aqlDollar[3].n)
		}
	case 35:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:125
		{
			aqlVAL.n = aqllex.(*bridge).Ident(aqlDollar[1].b)
		}
	case 36:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:126
		{
			v, _ := cvt.ParseInt64(aqlDollar[1].b)
			aqlVAL.n = aqllex.(*bridge).Literal(v)
		}
	case 37:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:130
		{
			aqlVAL.n = aqllex.(*bridge).Literal(string(aqlDollar[1].b))
		}
	case 38:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:131
		{
			aqlVAL.n = aqllex.(*bridge).Literal(true)
		}
	case 39:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:132
		{
			aqlVAL.n = aqllex.(*bridge).Literal(false)
		}
	case 40:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:133
		{
			aqlVAL.n = aqllex.(*bridge).Literal(nil)
		}
	case 41:
		aqlDollar = aqlS[aqlpt-1 : aqlpt+1]
//line aql.y:134
		{
			aqlVAL.n = aqllex.(*bridge).Dup()
		}
	case 42:
		aqlDollar = aqlS[aqlpt-3 : aqlpt+1]
//line aql.y:135
		{
			aqlVAL.n = aqlDollar[2].n
		}
	}
	goto aqlstack /* stack new state and value */
}
