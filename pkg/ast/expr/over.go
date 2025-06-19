package expr

import (
	"bytes"
	"errors"
	"strings"

	"github.com/xakepp35/aql/pkg/ast/asi"
)

type Over struct {
	Iter asi.AST // итератор
	Expr asi.AST // выражение
}

func (e Over) Kind() asi.Kind {
	return asi.Over
}

func (e Over) P0(c asi.Emitter) error {
	if err := e.Iter.P0(c); err != nil {
		return err
	}
	if err := e.Expr.P0(c); err != nil {
		return err
	}
	return nil
}

func (e Over) P1(c asi.Emitter) error {
	// // Кладём итератор
	// if err := e.Iter.P1(c); err != nil {
	// 	return err
	// }

	// end := c.I64(-1) // Резервируем место под адрес выхода (запишем позже)

	// // Кладём op.Over + место выхода (пока 0)
	// c.PC(c.Len()) // Начало цикла
	// c.RawU8(op.Over)

	// // Тело выражения, выполняемое на каждой итерации
	// if err := e.Expr.P1(c); err != nil {
	// 	return err
	// }

	// c.Set(end, c.PC()) // заменяем зарезервированное значение на текущий адрес
	// c.Op(end + 2)      // назад к op.Over, который был до тела

	// // Метка: сюда надо будет прыгнуть после окончания

	return errors.New("Over unimplemented")
}

func (e Over) P2(c asi.Emitter) error {
	if err := e.Iter.P2(c); err != nil {
		return err
	}
	if err := e.Expr.P2(c); err != nil {
		return err
	}
	return nil
}

func (e Over) BuildJSON(buf *bytes.Buffer) {
	buf.WriteString(`{"expr":"over","iter":`)
	e.Expr.BuildJSON(buf)
	buf.WriteString(`,"expr":`)
	buf.WriteByte('}')
}

func (e Over) BuildString(buf *strings.Builder) {
	buf.WriteString(`[over `)
	e.Iter.BuildString(buf)
	buf.WriteByte(' ')
	e.Expr.BuildString(buf)
	buf.WriteByte(']')
}
