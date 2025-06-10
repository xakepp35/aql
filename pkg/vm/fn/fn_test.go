package fn_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/require"
	"github.com/xakepp35/aql/pkg/vm"
	"github.com/xakepp35/aql/pkg/vm/fn"
)

func TestFnAdd_Int64(t *testing.T) {
	s := vm.NewState()
	s.Push(int64(2))
	s.Push(int64(3))
	fn.Add(s)
	require.NoError(t, s.Err())
	require.Equal(t, int64(5), s.Pop().(int64))
}

func TestFnAdd_String(t *testing.T) {
	s := vm.NewState()
	s.Push("Hello ")
	s.Push("World")
	fn.Add(s)
	require.NoError(t, s.Err())
	require.Equal(t, "Hello World", s.Pop().(string))
}

func TestFnAdd_Slice(t *testing.T) {
	s := vm.NewState()
	s.Push([]any{1, 2})
	s.Push([]any{3, 4})
	fn.Add(s)
	require.NoError(t, s.Err())
	// require.ElementsMatch(t, []any{1, 2, 3, 4}, s.Pop().([]any))
}

func TestFnAdd_Bool(t *testing.T) {
	s := vm.NewState()
	s.Push(true)
	s.Push(false)
	fn.Add(s)
	require.NoError(t, s.Err())
	require.Equal(t, true, s.Pop().(bool))
}

func TestFnNot_Int64(t *testing.T) {
	s := vm.NewState()
	s.Push(int64(5))
	fn.Not(s)
	require.NoError(t, s.Err())
	require.Equal(t, int64(-5), s.Pop().(int64))
}

func TestFnNot_Float64(t *testing.T) {
	s := vm.NewState()
	s.Push(float64(1.5))
	fn.Not(s)
	require.NoError(t, s.Err())
	require.Equal(t, -1.5, s.Pop().(float64))
}

func TestFnSub_Int64(t *testing.T) {
	s := vm.NewState()
	s.Push(int64(10))
	s.Push(int64(3))
	fn.Sub(s)
	require.NoError(t, s.Err())
	require.Equal(t, int64(7), s.Pop().(int64))
}

func TestFnSub_Float64(t *testing.T) {
	s := vm.NewState()
	s.Push(float64(10.5))
	s.Push(float64(0.5))
	fn.Sub(s)
	require.NoError(t, s.Err())
	require.Equal(t, 10.0, s.Pop())
}

func TestFnMod(t *testing.T) {
	s := vm.NewState()
	s.Push(int64(10))
	s.Push(int64(3))
	fn.Mod(s)
	require.NoError(t, s.Err())
	require.Equal(t, int64(1), s.Pop().(int64))
}

func TestFnEq(t *testing.T) {
	s := vm.NewState()
	s.Push("x")
	s.Push("x")
	fn.Eq(s)
	require.NoError(t, s.Err())
	require.Equal(t, true, s.Pop())
}

func TestFnNot(t *testing.T) {
	s := vm.NewState()
	s.Push(true)
	fn.Not(s)
	require.NoError(t, s.Err())
	require.Equal(t, false, s.Pop())
}

func TestFnOr(t *testing.T) {
	s := vm.NewState()
	s.Push(false)
	s.Push(true)
	fn.Or(s)
	require.NoError(t, s.Err())
	require.Equal(t, true, s.Pop())
}

func TestFnAnd(t *testing.T) {
	s := vm.NewState()
	s.Push(true)
	s.Push(true)
	fn.And(s)
	require.NoError(t, s.Err())
	require.Equal(t, true, s.Pop())
}

func TestFnLt_String(t *testing.T) {
	s := vm.NewState()
	s.Push("a")
	s.Push("b")
	fn.Lt(s)
	require.NoError(t, s.Err())
	require.Equal(t, true, s.Pop())
}
