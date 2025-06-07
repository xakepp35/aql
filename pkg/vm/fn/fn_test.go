package fn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xakepp35/aql/pkg/vm"
	"github.com/xakepp35/aql/pkg/vm/fn"
)

func TestFnAdd_Int64(t *testing.T) {
	s := vm.NewState()
	s.Push(int64(2))
	s.Push(int64(3))
	fn.Add(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, int64(5), s.Pop())
}

func TestFnAdd_String(t *testing.T) {
	s := vm.NewState()
	s.Push("Hello ")
	s.Push("World")
	fn.Add(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, "Hello World", s.Pop())
}

func TestFnAdd_Slice(t *testing.T) {
	s := vm.NewState()
	s.Push([]any{1, 2})
	s.Push([]any{3, 4})
	fn.Add(s)
	assert.NoError(t, s.Err())
	assert.ElementsMatch(t, []any{1, 2, 3, 4}, s.Pop())
}

func TestFnAdd_Map(t *testing.T) {
	s := vm.NewState()
	s.Push(map[string]any{"a": 1})
	s.Push(map[string]any{"b": 2})
	fn.Add(s)
	assert.NoError(t, s.Err())
	out := s.Pop().(map[string]any)
	assert.Equal(t, 2, len(out))
	assert.Equal(t, 1, out["a"])
	assert.Equal(t, 2, out["b"])
}

func TestFnAdd_Bool(t *testing.T) {
	s := vm.NewState()
	s.Push(true)
	s.Push(false)
	fn.Add(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, true, s.Pop())
}

func TestFnNot_Int64(t *testing.T) {
	s := vm.NewState()
	s.Push(int64(5))
	fn.Not(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, int64(-5), s.Pop())
}

func TestFnNot_Float64(t *testing.T) {
	s := vm.NewState()
	s.Push(float64(1.5))
	fn.Not(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, -1.5, s.Pop())
}

func TestFnSub_Int64(t *testing.T) {
	s := vm.NewState()
	s.Push(int64(10))
	s.Push(int64(3))
	fn.Sub(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, int64(7), s.Pop())
}

func TestFnSub_Float64(t *testing.T) {
	s := vm.NewState()
	s.Push(float64(10.5))
	s.Push(float64(0.5))
	fn.Sub(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, 10.0, s.Pop())
}

func TestFnMod(t *testing.T) {
	s := vm.NewState()
	s.Push(int64(10))
	s.Push(int64(3))
	fn.Mod(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, int64(1), s.Pop())
}

func TestFnEq(t *testing.T) {
	s := vm.NewState()
	s.Push("x")
	s.Push("x")
	fn.Eq(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, true, s.Pop())
}

func TestFnNot(t *testing.T) {
	s := vm.NewState()
	s.Push(true)
	fn.Not(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, false, s.Pop())
}

func TestFnOr(t *testing.T) {
	s := vm.NewState()
	s.Push(false)
	s.Push(true)
	fn.Or(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, true, s.Pop())
}

func TestFnAnd(t *testing.T) {
	s := vm.NewState()
	s.Push(true)
	s.Push(true)
	fn.And(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, true, s.Pop())
}

func TestFnLt_String(t *testing.T) {
	s := vm.NewState()
	s.Push("a")
	s.Push("b")
	fn.Lt(s)
	assert.NoError(t, s.Err())
	assert.Equal(t, true, s.Pop())
}
