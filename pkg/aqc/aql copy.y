%{
package compiler

import "github.com/xakepp35/aql/pkg/vmi"
%}

/* ───────────── union ───────────── */
%union {
  b []byte    /* литералы */
  i uint64    /* счётчик аргументов */
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
%type <i> arg_list

/* ───────────── приоритеты ───────────── */
%left  PIPE
%left  OROR
%left  ANDAND
%nonassoc EQ NEQ LT LE GT GE
%left  PLUS MINUS
%left  STAR SLASH PERCENT
%right UMINUS

%%

/* ┌───────────── выражения ─────────────┐ */
expr:
    pipe
  ;

pipe:
    or
  | pipe PIPE or   { aqllex.(*Compiler).Ops(op.Pipe) }
  ;

or:
    and
  | or OROR and    { aqllex.(*Compiler).Ops(op.Or) }
  ;

and:
    cmp
  | and ANDAND cmp { aqllex.(*Compiler).Ops(op.And) }
  ;

cmp:
    add
  | cmp EQ  add    { aqllex.(*Compiler).Ops(op.Eq) }
  | cmp NEQ add    { aqllex.(*Compiler).Ops(op.Neq) }
  | cmp LT  add    { aqllex.(*Compiler).Ops(op.Lt) }
  | cmp LE  add    { aqllex.(*Compiler).Ops(op.Le) }
  | cmp GT  add    { aqllex.(*Compiler).Ops(op.Gt) }
  | cmp GE  add    { aqllex.(*Compiler).Ops(op.Ge) }
  ;

add:
    mul
  | add PLUS  mul  { aqllex.(*Compiler).Ops(op.Add) }
  | add MINUS mul  { aqllex.(*Compiler).Ops(op.Sub) }
  ;

mul:
    unary
  | mul STAR    unary  { aqllex.(*Compiler).Ops(op.Mul) }
  | mul SLASH   unary  { aqllex.(*Compiler).Ops(op.Div) }
  | mul PERCENT unary  { aqllex.(*Compiler).Ops(op.Mod) }
  ;

/* ┌───────────── унарный ─────────────┐ */
unary:
    post
  | MINUS unary %prec UMINUS
      {
        /* 0 - X */
        aqllex.(*Compiler).Ops(op.Not)
      }
  | OVER unary
      { aqllex.(*Compiler).Ops(op.PushNull); aqllex.(*Compiler).Ops(op.Over) }
  | OVER unary ARROW LPAREN expr RPAREN
      { aqllex.(*Compiler).Ops(op.Over) }
  ;

/* ┌───────────── постфикс ─────────────┐ */
post:
    atom
  | post DOT IDENT
      { aqllex.(*Compiler).Ops(op.Field); aqllex.(*Compiler).ArgStr(string($3)) }
  | post LBRACK expr RBRACK
      { aqllex.(*Compiler).Ops(op.Index1) }
  | post LBRACK expr COLON expr RBRACK
      { aqllex.(*Compiler).Ops(op.Index2) }
  | IDENT LPAREN RPAREN
      {
        aqllex.(*Compiler).Ops(op.Call)
        aqllex.(*Compiler).ArgStr(string($1))
        aqllex.(*Compiler).Arg(0)
      }
  | IDENT LPAREN arg_list RPAREN
      {
        aqllex.(*Compiler).Ops(op.Call)
        aqllex.(*Compiler).ArgStr(string($1))
        aqllex.(*Compiler).Arg($3)
      }
  ;

/* ┌───────── список аргументов ─────────┐ */
arg_list:
    expr            { $$ = 1 }
  | arg_list COMMA expr
                    { $$ = $1 + 1 }
  ;

/* ┌───────────── атомы ─────────────┐ */
atom:
    IDENT           { aqllex.(*Compiler).Ops(op.Var);  aqllex.(*Compiler).ArgStr(string($1)) }
  | NUMBER          { aqllex.(*Compiler).Ops(op.PushNum);  aqllex.(*Compiler).ArgStr(string($1)) }
  | STRING          { aqllex.(*Compiler).Ops(op.PushStr);  aqllex.(*Compiler).ArgStr(string($1)) }
  | TRUE            { aqllex.(*Compiler).Ops(op.PushBool); aqllex.(*Compiler).Arg(1) }
  | FALSE           { aqllex.(*Compiler).Ops(op.PushBool); aqllex.(*Compiler).Arg(0) }
  | NULL            { aqllex.(*Compiler).Ops(op.PushNull) }
  | LPAREN expr RPAREN
                    { /* только группировка, ничего не эмитим */ }
  ;
%%
