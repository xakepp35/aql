%{
package ast
import "github.com/xakepp35/aql/pkg/vmi"
%}

%union {
	node  vmi.Node
	nodes []vmi.Node
	b     []byte
}

%token <b> IDENT NUMBER STRING
%token TRUE FALSE NULL
%token PLUS MINUS STAR SLASH PERCENT
%token PIPE
%token ANDAND OROR
%token EQ NEQ LT LE GT GE
%token DOT
%token LBRACK RBRACK
%token LPAREN RPAREN
%token COLON COMMA
%token OVER
%token ARROW

%type <node> expr pipe or and cmp add mul unary post atom
%type <nodes> arg_list

%start query

/* ---------- приоритеты ---------- */
%left  PIPE
%left  OROR
%left  ANDAND
%nonassoc EQ NEQ LT LE GT GE
%left  PLUS MINUS
%left  STAR SLASH PERCENT
%right UMINUS
%%

query :
	/* expr EOF      { aqllex.(*bridge).result = $1 } */
	expr           { aqllex.(*bridge).result = $1 }
	;

/* pipeline lowest */
expr   : pipe

pipe   : or
       | pipe PIPE or          { $$ = &PipeExpr{Left:$1, Right:$3} }
       ;

or     : and
       | or OROR and           { $$ = &LogicalExpr{Op:"||", Left:$1, Right:$3} }
       ;

and    : cmp
       | and ANDAND cmp        { $$ = &LogicalExpr{Op:"&&", Left:$1, Right:$3} }
       ;

cmp    : add
       | cmp EQ  add           { $$ = &CompareExpr{Op:"==", Left:$1, Right:$3} }
       | cmp NEQ add           { $$ = &CompareExpr{Op:"!=", Left:$1, Right:$3} }
       | cmp LT  add           { $$ = &CompareExpr{Op:"<",  Left:$1, Right:$3} }
       | cmp LE  add           { $$ = &CompareExpr{Op:"<=", Left:$1, Right:$3} }
       | cmp GT  add           { $$ = &CompareExpr{Op:">",  Left:$1, Right:$3} }
       | cmp GE  add           { $$ = &CompareExpr{Op:">=", Left:$1, Right:$3} }
       ;

add    : mul
       | add PLUS  mul         { $$ = &BinaryExpr{Op:"+", Left:$1, Right:$3} }
       | add MINUS mul         { $$ = &BinaryExpr{Op:"-", Left:$1, Right:$3} }
       ;

mul    : unary
       | mul STAR    unary     { $$ = &BinaryExpr{Op:"*", Left:$1, Right:$3} }
       | mul SLASH   unary     { $$ = &BinaryExpr{Op:"/", Left:$1, Right:$3} }
       | mul PERCENT unary     { $$ = &BinaryExpr{Op:"%", Left:$1, Right:$3} }
       ;

unary  : post
       | MINUS unary %prec UMINUS { $$ = &UnaryExpr{Op:"-", X:$2} }
       | OVER unary               { $$ = &OverExpr{Seq:$2} }
       | OVER unary ARROW LPAREN expr RPAREN
                                  { $$ = &OverExpr{Seq:$2, Scope:$5} }
       ;

/* постфиксная цепочка */
post   : atom
       | post DOT IDENT                { $$ = &FieldSel{X:$1, Name:$3} }
       | post LBRACK expr RBRACK       { $$ = &IndexExpr{X:$1, I:$3} }
       | post LBRACK expr COLON expr RBRACK { $$ = &IndexExpr{X:$1, I:$3, J:$5} }
       | post LPAREN arg_list RPAREN   { id := $1.(*Ident); $$ = &CallExpr{Fun:id.Name, Args:$3} }
       | post LPAREN RPAREN            { id := $1.(*Ident); $$ = &CallExpr{Fun:id.Name} }
       ;

arg_list :
         expr                     { $$ = []vmi.Node{$1} }
       | arg_list COMMA expr      { $$ = append($1, $3) }
       ;

atom   :
         IDENT                    { $$ = &Ident{Name:$1} }
       | NUMBER                   { $$ = &Number{Text:$1} }
       | STRING                   { $$ = &String{Text:$1} }
       | TRUE                     { $$ = &Bool{Val:true} }
       | FALSE                    { $$ = &Bool{Val:false} }
       | NULL                     { $$ = &Null{} }
       | LPAREN expr RPAREN       { $$ = $2 }
       ;
%%
