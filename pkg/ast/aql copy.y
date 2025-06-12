%{
package ast

import (
	"github.com/xakepp35/aql/pkg/ast/asi"
    "github.com/xakepp35/aql/pkg/ast/expr"
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
	| pipe PIPE or         { $$ = &expr.Pipe{Args: [2]asi.AST{$1, $3}} }
	/* «semantically binary, but emits no op»  → отдельный узел Pipe */
	;

/* ── || ── */
or:
	  and                  { $$ = $1 }
	| or OROR and          { $$ = &expr.Binary{Args: [2]asi.AST{$1, $3}, Op: op.Or} }
	;

/* ── && ── */
and:
	  cmp                  { $$ = $1 }
	| and ANDAND cmp       { $$ = &expr.Binary{Args: [2]asi.AST{$1, $3}, Op: op.And} }
	;

/* ── ==, !=, <, … ── */
cmp:
	  add                       { $$ = $1 }
	| cmp EQ  add               { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Eq} }
	| cmp NEQ add               { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Neq} }
	| cmp LT  add               { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Lt} }
	| cmp LE  add               { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Le} }
	| cmp GT  add               { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Gt} }
	| cmp GE  add               { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Ge} }
	;

/* ── +, – ── */
add:
	  mul                       { $$ = $1 }
	| add PLUS  mul             { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Add} }
	| add MINUS mul             { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Sub} }
	;

/* ── *, /, % ── */
mul:
	  unary                     { $$ = $1 }
	| mul STAR    unary         { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Mul} }
	| mul SLASH   unary         { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Div} }
	| mul PERCENT unary         { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Mod} }
	;

/* ┌───────────── унарный ─────────────┐ */
unary:
	  post                                      { $$ = $1 }
	| MINUS unary %prec UMINUS                 { $$ = &expr.Unary{Arg: $2, Op: op.Not} }
	| OVER unary                               { $$ = &expr.Over{Iter: $2, Expr: nil} }
	| OVER unary ARROW LPAREN expr RPAREN      { $$ = &expr.Over{Iter: $2, Expr: $5} }
	;

/* ┌───────────── постфикс ─────────────┐ */
post:
	  atom                                                         { $$ = $1 }
	| post DOT IDENT                                               { $$ = &expr.Field{Arg: $1, Name: $3} }
	| post LBRACK expr RBRACK                                      { $$ = &expr.Binary{Args: [2]asi.AST{$1,$3}, Op: op.Index1} }
	| post LBRACK expr COLON expr RBRACK                           { $$ = &expr.Ternary{Args: [3]asi.AST{$1,$3,$5}, Op: op.Index2} }
	| IDENT LPAREN RPAREN                                          { $$ = &expr.Call{Args: nil, Name: $1} }
	| IDENT LPAREN arg_list RPAREN                                 { $$ = &expr.Call{Args: $3, Name: $1} }
	;

/* ┌───────── список аргументов ─────────┐ */
arg_list:
	  expr                           { $$ = []asi.AST{$1} }
	| arg_list COMMA expr            { $$ = append($1, $3) }
	;

/* ┌───────────── атомы ─────────────┐ */
atom:
	  IDENT   { $$ = &expr.Ident{Name: $1} }
	| NUMBER  {
		v, _ := cvt.ParseInt64($1)
		$$ = &expr.Literal{X: v}
	  }
	| STRING  { $$ = &expr.Literal{X: string($1)} }
	| TRUE    { $$ = &expr.Literal{X: true} }
	| FALSE   { $$ = &expr.Literal{X: false} }
	| NULL    { $$ = &expr.Literal{X: nil} }
	| DOT     { $$ = &expr.Dup{} }              /* "." - is an alias for top-of-stack */
	| LPAREN expr RPAREN { $$ = $2 }            /* просто группировка */
	;
%%
