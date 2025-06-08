package expr

import (
	"bytes"
	"strings"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

type Over struct {
	Iter vmi.AST // итератор
	Expr vmi.AST // выражение
}

func (e *Over) Pre(c vmi.Compiler) error {
	if err := e.Iter.Pre(c); err != nil {
		return err
	}
	if err := e.Expr.Pre(c); err != nil {
		return err
	}
	return nil
}

func (e *Over) Body(c vmi.Compiler) error {
	// Кладём итератор
	if err := e.Iter.Body(c); err != nil {
		return err
	}

	end := c.Int(-1) // Резервируем место под адрес выхода (запишем позже)

	// Кладём op.Over + место выхода (пока 0)
	c.Int(c.PC()) // Начало цикла
	c.Op(op.Over)

	// Тело выражения, выполняемое на каждой итерации
	if err := e.Expr.Body(c); err != nil {
		return err
	}

	c.Set(end, c.PC()) // заменяем зарезервированное значение на текущий адрес
	c.Op(end + 2)      // назад к op.Over, который был до тела

	// Метка: сюда надо будет прыгнуть после окончания

	return nil
}

func (e *Over) Post(c vmi.Compiler) error {
	if err := e.Iter.Post(c); err != nil {
		return err
	}
	if err := e.Expr.Post(c); err != nil {
		return err
	}
	return nil
}

func (e *Over) BuildJSON(buf *bytes.Buffer) {
	buf.WriteString(`{"expr":"over","iter":`)
	e.Expr.BuildJSON(buf)
	buf.WriteString(`,"expr":`)
	buf.WriteByte('}')
}

func (e *Over) BuildString(buf *strings.Builder) {
	buf.WriteString(`[over `)
	e.Iter.BuildString(buf)
	buf.WriteByte(' ')
	e.Expr.BuildString(buf)
	buf.WriteByte(']')
}
