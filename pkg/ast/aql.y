%{
package ast

import (
	"github.com/xakepp35/aql/pkg/ast/asi"
    "github.com/xakepp35/aql/pkg/cvt"
	"github.com/xakepp35/aql/pkg/vm/op"
)
%}

/* ───────────── union ───────────── */
%union {
	b []byte          /* сырые лексемы-строки от лексера */
	i int64           /* счётчики (например, кол-во арг-ов) */

	n asi.AST        /* любой узел AST */
	a []asi.AST      /* срез узлов (список арг-ов) */
}

/* ───────────── лексемы ───────────── */
%token <b> IDENT NUMBER STRING
%token TRUE FALSE NULL
%token PLUS MINUS STAR SLASH PERCENT
%token PIPE ANDAND OROR
%token EQ NEQ LT LE GT GE
%token DOT LBRACK RBRACK LPAREN RPAREN COLON COMMA
%token OVER ARROW

/* ───────────── sem-types ───────────── */
%type <n> query expr pipe or and cmp add mul unary post atom
%type <a> arg_list

/* ───────────── приоритеты ───────────── */
%left  PIPE
%left  OROR
%left  ANDAND
%nonassoc EQ NEQ LT LE GT GE
%left  PLUS MINUS
%left  STAR SLASH PERCENT
%right UMINUS

%%

/* ┌───────────── старт ─────────────┐ */
query:
	  expr          { aqllex.(*bridge).result = $1 }
	;

/* ┌───────────── выражения ─────────────┐ */
expr:
	  pipe          { $$ = $1 }
	;

/* ── оператор | (pipe) ── */
pipe:
	  or                   { $$ = $1 }
	| pipe PIPE or         { $$ = aqllex.(*bridge).Pipe($1, $3) }
	/* «semantically binary, but emits no op»  → отдельный узел Pipe */
	;

/* ── || ── */
or:
	  and                  { $$ = $1 }
	| or OROR and          { $$ = aqllex.(*bridge).Binary($1, $3, op.Or) }
	;

/* ── && ── */
and:
	  cmp                  { $$ = $1 }
	| and ANDAND cmp       { $$ = aqllex.(*bridge).Binary($1, $3, op.And) }
	;

/* ── ==, !=, <, … ── */
cmp:
	  add                       { $$ = $1 }
	| cmp EQ  add               { $$ = aqllex.(*bridge).Binary($1,$3, op.Eq) }
	| cmp NEQ add               { $$ = aqllex.(*bridge).Binary($1,$3, op.Neq) }
	| cmp LT  add               { $$ = aqllex.(*bridge).Binary($1,$3, op.Lt) }
	| cmp LE  add               { $$ = aqllex.(*bridge).Binary($1,$3, op.Le) }
	| cmp GT  add               { $$ = aqllex.(*bridge).Binary($1,$3, op.Gt) }
	| cmp GE  add               { $$ = aqllex.(*bridge).Binary($1,$3, op.Ge) }
	;

/* ── +, – ── */
add:
	  mul                       { $$ = $1 }
	| add PLUS  mul             { $$ = aqllex.(*bridge).Binary($1,$3, op.Add) }
	| add MINUS mul             { $$ = aqllex.(*bridge).Binary($1,$3, op.Sub) }
	;

/* ── *, /, % ── */
mul:
	  unary                     { $$ = $1 }
	| mul STAR    unary         { $$ = aqllex.(*bridge).Binary($1,$3, op.Mul) }
	| mul SLASH   unary         { $$ = aqllex.(*bridge).Binary($1,$3, op.Div) }
	| mul PERCENT unary         { $$ = aqllex.(*bridge).Binary($1,$3, op.Mod) }
	;

/* ┌───────────── унарный ─────────────┐ */
unary:
	  post                                      { $$ = $1 }
	| MINUS unary %prec UMINUS                 { $$ = aqllex.(*bridge).Unary($2, op.Not) }
	| OVER unary                               { $$ = aqllex.(*bridge).Over($2, nil) }
	| OVER unary ARROW LPAREN expr RPAREN      { $$ = aqllex.(*bridge).Over($2, $5) }
	;

/* ┌───────────── постфикс ─────────────┐ */
post:
	  atom                                                         { $$ = $1 }
	| post DOT IDENT                                               { $$ = aqllex.(*bridge).Field($1, $3) }
	| post LBRACK expr RBRACK                                      { $$ = aqllex.(*bridge).Binary($1,$3, op.Index1) }
	| post LBRACK expr COLON expr RBRACK                           { $$ = aqllex.(*bridge).Ternary($1,$3,$5, op.Index2) }
	| IDENT LPAREN RPAREN                                          { $$ = aqllex.(*bridge).Call(nil, $1) }
	| IDENT LPAREN arg_list RPAREN                                 { $$ = aqllex.(*bridge).Call($3, $1) }
	;

/* ┌───────── список аргументов ─────────┐ */
arg_list:
	  expr                           { $$ = []asi.AST{$1} }
	| arg_list COMMA expr            { $$ = append($1, $3) }
	;

/* ┌───────────── атомы ─────────────┐ */
atom:
	  IDENT   { $$ = aqllex.(*bridge).Ident($1) }
	| NUMBER  {
		v, _ := cvt.ParseInt64($1)
		$$ = aqllex.(*bridge).Literal(v)
	  }
	| STRING  { $$ = aqllex.(*bridge).Literal(string($1)) }
	| TRUE    { $$ = aqllex.(*bridge).Literal(true) }
	| FALSE   { $$ = aqllex.(*bridge).Literal(false) }
	| NULL    { $$ = aqllex.(*bridge).Literal(nil) }
	| DOT     { $$ = aqllex.(*bridge).Dup() }              /* "." - is an alias for top-of-stack */
	| LPAREN expr RPAREN { $$ = $2 }            /* просто группировка */
	;
%%
