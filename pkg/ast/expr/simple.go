package expr

import "github.com/xakepp35/aql/pkg/ast/asi"

func (s *Arena) I64(v int64) asi.AST {
	return s.Literal(v)
}

func (s *Arena) F64(v float64) asi.AST {
	return s.Literal(v)
}

func (s *Arena) True() asi.AST {
	return s.Literal(true)
}

func (s *Arena) False() asi.AST {
	return s.Literal(false)
}

func (s *Arena) Nil() asi.AST {
	return s.Literal(nil)
}

func (s *Arena) StringBytes(v []byte) asi.AST {
	return s.Literal(string(v))
}
