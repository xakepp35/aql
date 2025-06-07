package vm_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xakepp35/aql/pkg/vm"
	"github.com/xakepp35/aql/pkg/vm/op"
)

func TestProgramMarshalUnmarshal(t *testing.T) {
	t.Run("round trip", func(t *testing.T) {
		e := vm.NewCompiler()
		e.EmitInt(7)
		e.EmitString("hello")
		e.EmitBool(true)
		e.EmitNull()
		e.EmitOps(op.Add)
		e.EmitOps(op.Or)

		bin, err := e.MarshalBinary()
		require.NoError(t, err)
		require.True(t, len(bin) > 0)

		clone := vm.NewCompiler()
		err = clone.UnmarshalBinary(bin)
		require.NoError(t, err)

		s1 := vm.NewState()
		s2 := vm.NewState()
		e.Init(s1)
		clone.Init(s2)
		require.Equal(t, s1.Len(), s2.Len())
		for i := 0; i < s1.Len(); i++ {
			require.Equal(t, s1.Pop(), s2.Pop())
		}
		require.Equal(t, e.Program(), clone.Program())
	})
}

func TestProgramUnmarshalMarshal(t *testing.T) {
	t.Run("manual binary round trip", func(t *testing.T) {
		p := vm.NewCompiler()
		p.EmitString("world")
		p.EmitInt(99)
		p.EmitOps(op.Add)

		bin1, err := p.MarshalBinary()
		require.NoError(t, err)

		copy := vm.NewCompiler()
		err = copy.UnmarshalBinary(bin1)
		require.NoError(t, err)

		bin2, err := copy.MarshalBinary()
		require.NoError(t, err)
		require.True(t, bytes.Equal(bin1, bin2))
	})
}

func TestProgramMarshalErrors(t *testing.T) {
	t.Run("unmarshal too short", func(t *testing.T) {
		e := vm.NewCompiler()
		err := e.UnmarshalBinary([]byte{1, 2, 3})
		require.Error(t, err)
	})

	t.Run("unmarshal invalid size", func(t *testing.T) {
		e := vm.NewCompiler()
		buf := make([]byte, 32)
		buf[0] = 255 // total size too big
		err := e.UnmarshalBinary(buf)
		require.Error(t, err)
	})
}
