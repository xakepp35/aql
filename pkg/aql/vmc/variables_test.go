package vmc_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/aql/vmc"
	"github.com/xakepp35/aql/pkg/require"
)

func TestVariablesBasic(t *testing.T) {
	t.Run("Set and Get", func(t *testing.T) {
		vars := vmc.NewVariables()
		vars.Set("x", 123)
		v, ok := vars.Get("x")
		require.Equal(t, 123, v)
		require.True(t, ok)
	})

	t.Run("Overwrite value", func(t *testing.T) {
		vars := vmc.NewVariables()
		vars.Set("x", 1)
		vars.Set("x", 2)
		v, ok := vars.Get("x")
		require.Equal(t, 2, v)
		require.True(t, ok)
	})

	t.Run("Delete key", func(t *testing.T) {
		vars := vmc.NewVariables()
		vars.Set("x", 1)
		vars.Del("x")
		v, ok := vars.Get("x")
		require.Nil(t, v)
		require.False(t, ok)
	})

	t.Run("Clear map", func(t *testing.T) {
		vars := vmc.NewVariables()
		vars.Set("a", 1)
		vars.Set("b", 2)
		vars.Clear()
		vars.Set("c", 3) // ensure we can still set after clear
		a, ok := vars.Get("a")
		require.Nil(t, a)
		require.False(t, ok)
		b, ok := vars.Get("b")
		require.Nil(t, b)
		require.False(t, ok)
		c, ok := vars.Get("c")
		require.Equal(t, 3, c)
		require.True(t, ok)
	})

	t.Run("Empty get returns nil", func(t *testing.T) {
		vars := vmc.NewVariables()
		v, ok := vars.Get("x")
		require.Nil(t, v)
		require.False(t, ok)
	})
}
