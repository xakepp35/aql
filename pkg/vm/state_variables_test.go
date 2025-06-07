package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xakepp35/aql/pkg/vm"
)

func TestVariablesBasic(t *testing.T) {
	t.Run("Set and Get", func(t *testing.T) {
		vars := vm.NewVariables()
		vars.Set("x", 123)
		require.Equal(t, 123, vars.Get("x"))
	})

	t.Run("Overwrite value", func(t *testing.T) {
		vars := vm.NewVariables()
		vars.Set("x", 1)
		vars.Set("x", 2)
		require.Equal(t, 2, vars.Get("x"))
	})

	t.Run("Delete key", func(t *testing.T) {
		vars := vm.NewVariables()
		vars.Set("x", 1)
		vars.Del("x")
		require.Nil(t, vars.Get("x"))
	})

	t.Run("Clear map", func(t *testing.T) {
		vars := vm.NewVariables()
		vars.Set("a", 1)
		vars.Set("b", 2)
		vars.Clear()
		vars.Set("c", 3) // ensure we can still set after clear
		require.Nil(t, vars.Get("a"))
		require.Nil(t, vars.Get("b"))
		require.Equal(t, 3, vars.Get("c"))
	})

	t.Run("Empty get returns nil", func(t *testing.T) {
		vars := vm.NewVariables()
		require.Nil(t, vars.Get("nothing"))
	})
}
