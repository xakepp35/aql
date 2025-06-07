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
  | pipe PIPE or   { aqllex.(*Compiler).EmitOps(op.Pipe) }
  ;

or:
    and
  | or OROR and    { aqllex.(*Compiler).EmitOps(op.Or) }
  ;

and:
    cmp
  | and ANDAND cmp { aqllex.(*Compiler).EmitOps(op.And) }
  ;

cmp:
    add
  | cmp EQ  add    { aqllex.(*Compiler).EmitOps(op.Eq) }
  | cmp NEQ add    { aqllex.(*Compiler).EmitOps(op.Neq) }
  | cmp LT  add    { aqllex.(*Compiler).EmitOps(op.Lt) }
  | cmp LE  add    { aqllex.(*Compiler).EmitOps(op.Le) }
  | cmp GT  add    { aqllex.(*Compiler).EmitOps(op.Gt) }
  | cmp GE  add    { aqllex.(*Compiler).EmitOps(op.Ge) }
  ;

add:
    mul
  | add PLUS  mul  { aqllex.(*Compiler).EmitOps(op.Add) }
  | add MINUS mul  { aqllex.(*Compiler).EmitOps(op.Sub) }
  ;

mul:
    unary
  | mul STAR    unary  { aqllex.(*Compiler).EmitOps(op.Mul) }
  | mul SLASH   unary  { aqllex.(*Compiler).EmitOps(op.Div) }
  | mul PERCENT unary  { aqllex.(*Compiler).EmitOps(op.Mod) }
  ;

/* ┌───────────── унарный ─────────────┐ */
unary:
    post
  | MINUS unary %prec UMINUS
      {
        /* 0 - X */
        aqllex.(*Compiler).EmitOps(op.Not)
      }
  | OVER unary
      { aqllex.(*Compiler).EmitOps(op.PushNull); aqllex.(*Compiler).EmitOps(op.Over) }
  | OVER unary ARROW LPAREN expr RPAREN
      { aqllex.(*Compiler).EmitOps(op.Over) }
  ;

/* ┌───────────── постфикс ─────────────┐ */
post:
    atom
  | post DOT IDENT
      { aqllex.(*Compiler).EmitOps(op.Field); aqllex.(*Compiler).EmitArgStr(string($3)) }
  | post LBRACK expr RBRACK
      { aqllex.(*Compiler).EmitOps(op.Index1) }
  | post LBRACK expr COLON expr RBRACK
      { aqllex.(*Compiler).EmitOps(op.Index2) }
  | IDENT LPAREN RPAREN
      {
        aqllex.(*Compiler).EmitOps(op.Call)
        aqllex.(*Compiler).EmitArgStr(string($1))
        aqllex.(*Compiler).EmitArg(0)
      }
  | IDENT LPAREN arg_list RPAREN
      {
        aqllex.(*Compiler).EmitOps(op.Call)
        aqllex.(*Compiler).EmitArgStr(string($1))
        aqllex.(*Compiler).EmitArg($3)
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
    IDENT           { aqllex.(*Compiler).EmitOps(op.PushVar);  aqllex.(*Compiler).EmitArgStr(string($1)) }
  | NUMBER          { aqllex.(*Compiler).EmitOps(op.PushNum);  aqllex.(*Compiler).EmitArgStr(string($1)) }
  | STRING          { aqllex.(*Compiler).EmitOps(op.PushStr);  aqllex.(*Compiler).EmitArgStr(string($1)) }
  | TRUE            { aqllex.(*Compiler).EmitOps(op.PushBool); aqllex.(*Compiler).EmitArg(1) }
  | FALSE           { aqllex.(*Compiler).EmitOps(op.PushBool); aqllex.(*Compiler).EmitArg(0) }
  | NULL            { aqllex.(*Compiler).EmitOps(op.PushNull) }
  | LPAREN expr RPAREN
                    { /* только группировка, ничего не эмитим */ }
  ;
%%
