package vm_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/require"
	"github.com/xakepp35/aql/pkg/vm"
	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

func TestProgramInit(t *testing.T) {
	t.Run("init stack from args", func(t *testing.T) {
		e := vm.NewCompiler()
		e.EmitInt(42)
		e.EmitBool(true)
		e.EmitBool(false)
		e.EmitNull()
		e.EmitString("hello")

		s := vm.NewState()
		e.Init(s)

		require.Equal(t, "hello", s.Pop())
		require.Nil(t, s.Pop())
		require.Equal(t, false, s.Pop())
		require.Equal(t, true, s.Pop())
		require.Equal(t, int64(42), s.Pop().(int64))
	})

	t.Run("empty init", func(t *testing.T) {
		e := vm.NewCompiler()
		this := vm.NewState()
		e.Init(this)
		require.Equal(t, 0, this.Len())
	})
}

func TestEmitterRunVariants(t *testing.T) {
	testCases := []struct {
		name   string
		init   func() vmi.Compiler
		expect any
	}{
		{
			name: "int64 addition",
			init: func() vmi.Compiler {
				e := vm.NewCompiler()
				e.EmitInt(1)
				e.EmitInt(2)
				e.EmitOps(op.Add)
				return e
			},
			expect: int64(3),
		},
		{
			name: "string concat",
			init: func() vmi.Compiler {
				e := vm.NewCompiler()
				e.EmitString("foo")
				e.EmitString("bar")
				e.EmitOps(op.Add)
				return e
			},
			expect: "foobar",
		},
		{
			name: "bool or",
			init: func() vmi.Compiler {
				e := vm.NewCompiler()
				e.EmitBool(true)
				e.EmitBool(false)
				e.EmitOps(op.Or)
				return e
			},
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := tc.init()
			this := vm.NewState()
			e.Init(this)
			for this.Next(e.Program()) {
			}
			require.NoError(t, this.Err())
			v := this.Pop()

			require.Equal(t, tc.expect, v)
		})
	}
}
