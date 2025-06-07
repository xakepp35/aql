%{
package aqc

import "github.com/xakepp35/aql/pkg/vm/op"
%}

/* ───────────── union ───────────── */
%union {
  b []byte   /* литералы */
  i int64    /* счётчик аргументов */
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

/* ┌───────────── старт ─────────────┐ */
query:
    expr
  ;

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
  | MINUS unary %prec UMINUS { aqllex.(*Compiler).EmitOps(op.Not) }
  | OVER unary               { aqllex.(*Compiler).EmitNull(); aqllex.(*Compiler).EmitOps(op.Over) }
  | OVER unary ARROW LPAREN expr RPAREN { aqllex.(*Compiler).EmitOps(op.Over) }
  ;

/* ┌───────────── постфикс ─────────────┐ */
post:
    atom
  | post DOT IDENT { 
        c:=aqllex.(*Compiler); 
        c.EmitString(string($3)); 
        c.EmitOps(op.Field) 
  }
  | post LBRACK expr RBRACK           { aqllex.(*Compiler).EmitOps(op.Index1) }
  | post LBRACK expr COLON expr RBRACK { aqllex.(*Compiler).EmitOps(op.Index2) }
  | IDENT LPAREN RPAREN {
      c:=aqllex.(*Compiler)
      name := string($1)
      if builtin, ok := Builtins[name]; ok {
          c.EmitInt(0)
          c.EmitOps(builtin)
      } else {
          c.EmitString(name)
          c.EmitInt(0)
          c.EmitOps(op.Call)
      }
  }
  | IDENT LPAREN arg_list RPAREN {
      c := aqllex.(*Compiler)
      name := string($1)
      if builtin, ok := Builtins[name]; ok {
          c.EmitInt($3)  // arg count
          c.EmitOps(builtin)
      } else {
          c.EmitString(name)
          c.EmitInt($3)
          c.EmitOps(op.Call)
      }
  }
;

/* ┌───────── список аргументов ─────────┐ */
arg_list:
    expr                { $$ = 1 }
  | arg_list COMMA expr { $$ = $1 + 1 }
;

/* ┌───────────── атомы ─────────────┐ */
atom:
    IDENT   { c:=aqllex.(*Compiler); c.EmitString(string($1)); c.EmitOps(op.PushVar) }
  | NUMBER  { aqllex.(*Compiler).EmitInt(parseInt($1)) }
  | STRING  { aqllex.(*Compiler).EmitString(string($1)) }
  | TRUE    { aqllex.(*Compiler).EmitBool(true)  }
  | FALSE   { aqllex.(*Compiler).EmitBool(false) }
  | NULL    { aqllex.(*Compiler).EmitNull() }
  | LPAREN expr RPAREN { /* только группировка, ничего не эмитим */ }
  ;
%%
