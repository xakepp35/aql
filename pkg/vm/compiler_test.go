package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xakepp35/aql/pkg/vm"
	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

func TestProgramInit(t *testing.T) {
	t.Run("init stack from args", func(t *testing.T) {
		e := vm.NewProgrammer()
		e.Int(42)
		e.Bool(true)
		e.Bool(false)
		e.Null()
		e.String("hello")

		s := vm.NewState()
		e.Init(s)

		require.Equal(t, "hello", s.Pop())
		require.Nil(t, s.Pop())
		require.Equal(t, false, s.Pop())
		require.Equal(t, true, s.Pop())
		require.Equal(t, int64(42), s.Pop())
	})

	t.Run("empty init", func(t *testing.T) {
		e := vm.NewProgrammer()
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
				e := vm.NewProgrammer()
				e.Int(1)
				e.Int(2)
				e.Ops(op.Add)
				return e
			},
			expect: int64(3),
		},
		{
			name: "string concat",
			init: func() vmi.Compiler {
				e := vm.NewProgrammer()
				e.String("foo")
				e.String("bar")
				e.Ops(op.Add)
				return e
			},
			expect: "foobar",
		},
		{
			name: "bool or",
			init: func() vmi.Compiler {
				e := vm.NewProgrammer()
				e.Bool(true)
				e.Bool(false)
				e.Ops(op.Or)
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
