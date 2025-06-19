package asi

type AST interface {
	Expr
	JSONBuilder
	StringBuilder
}
