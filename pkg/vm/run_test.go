package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/xakepp35/aql/pkg/vm"
)

func TestRunPrograms(t *testing.T) {
	t.Run("int + int", runCase("int + int", `1 + 2`, nil, nil, int64(3)))
	// t.Run("float + float", runCase("float + float", `1.5 + 2.5`, nil, nil, 4.0))
	// t.Run("string + string", runCase("string + string", `"foo" + "bar"`, nil, nil, "foobar"))
	// t.Run("array concat", runCase("array concat", `[1,2] + [3]`, nil, nil, []any{int64(1), int64(2), int64(3)}))
	// t.Run("map merge", runCase("map merge", `{a: 1} + {b: 2}`, nil, nil, map[string]any{"a": int64(1), "b": int64(2)}))
	// t.Run("bool or", runCase("bool or", `true + false`, nil, nil, true))
	// t.Run("negation", runCase("negation", `-42`, nil, nil, int64(-42)))
	// t.Run("subtraction", runCase("subtraction", `100 - 58`, nil, nil, int64(42)))
	// t.Run("multiplication", runCase("multiplication", `6 * 7`, nil, nil, int64(42)))
	// t.Run("division", runCase("division", `84 / 2`, nil, nil, int64(42)))
	// t.Run("modulo", runCase("modulo", `85 % 2`, nil, nil, int64(1)))
	// t.Run("equality", runCase("equality", `1 == 1`, nil, nil, true))
	// t.Run("inequality", runCase("inequality", `2 != 3`, nil, nil, true))
	// t.Run("lt int", runCase("lt int", `3 < 4`, nil, nil, true))
	// t.Run("lt string", runCase("lt string", `'a' < 'b'`, nil, nil, true))
	// t.Run("le string", runCase("le string", `'a' <= 'a'`, nil, nil, true))
	// t.Run("gt float", runCase("gt float", `3.14 > 2.71`, nil, nil, true))
	// t.Run("ge int", runCase("ge int", `5 >= 5`, nil, nil, true))
	// t.Run("literal true", runCase("literal true", `true`, nil, nil, true))
	// t.Run("null push", runCase("null push", `null`, nil, nil, nil))
	// t.Run("not false", runCase("not false", `!false`, nil, nil, true))
	// t.Run("pipe chain", runCase("pipe chain", `1 | 2 | 3`, nil, nil, int64(3)))
	// t.Run("input passthrough", runCase("input passthrough", `.`, 42, nil, int64(42)))
	// t.Run("input used in expr", runCase("input used in expr", `. + 1`, 41, nil, int64(42)))
	// t.Run("var access", runCase("var access", `$x`, nil, map[string]any{"x": 42}, int64(42)))
	// t.Run("var math", runCase("var math", `$x + 8`, nil, map[string]any{"x": 34}, int64(42)))
	// t.Run("var as string", runCase("var as string", `$name + '!'`, nil, map[string]any{"name": "Bob"}, "Bob!"))
	// t.Run("call min", runCase("call min", `min(3,1,2)`, nil, nil, int64(1)))
	// t.Run("call max", runCase("call max", `max(3,1,2)`, nil, nil, int64(3)))
	// t.Run("field access", runCase("field access", `{a: 42}.a`, nil, nil, int64(42)))
	// t.Run("index array", runCase("index array", `[10,20,30][1]`, nil, nil, int64(20)))
	// t.Run("index map", runCase("index map", `{x: [1,2,3]}["x"]`, nil, nil, []any{int64(1), int64(2), int64(3)}))
	// t.Run("range index", runCase("range index", `[10,20,30,40][1:3]`, nil, nil, []any{int64(20), int64(30)}))
	// t.Run("over apply", runCase("over apply", `42 over (.+1)`, nil, nil, int64(43)))
	// t.Run("complex pipe", runCase("complex pipe", `1 | (. + 1) | (. * 2)`, nil, nil, int64(4)))
	// t.Run("nested field/index", runCase("nested field/index", `{a: {b: [5,6,7]}}.a.b[1]`, nil, nil, int64(6)))
	// t.Run("pipe with var", runCase("pipe with var", `$val | (. * 2)`, nil, map[string]any{"val": 21}, int64(42)))
	// t.Run("null eq", runCase("null eq", `null == null`, nil, nil, true))
	// t.Run("null neq", runCase("null neq", `null != null`, nil, nil, false))
	// t.Run("var not exist", runCase("var not exist", `$notfound == null`, nil, nil, true))
	// t.Run("or logic", runCase("or logic", `true || false`, nil, nil, true))
	// t.Run("and logic", runCase("and logic", `true && false`, nil, nil, false))
	// t.Run("nested math", runCase("nested math", `(1 + 2) * (3 + 4)`, nil, nil, int64(21)))
	// t.Run("expr in over", runCase("expr in over", `over (.+1) 41`, nil, nil, int64(42)))
	// t.Run("var bool logic", runCase("var bool logic", `$a && $b`, nil, map[string]any{"a": true, "b": false}, false))
	// t.Run("var object merge", runCase("var object merge", `$a + $b`, nil, map[string]any{"a": map[string]any{"x": 1}, "b": map[string]any{"y": 2}}, map[string]any{"x": int64(1), "y": int64(2)}))
	// t.Run("string compare", runCase("string compare", `'foo' > 'bar'`, nil, nil, true))
	// t.Run("string <= compare", runCase("string <= compare", `'bar' <= 'baz'`, nil, nil, true))
	// t.Run("deep equality", runCase("deep equality", `{x: [1,2]} == {x: [1,2]}`, nil, nil, true))
	// t.Run("deep inequality", runCase("deep inequality", `{x: [1,2]} != {x: [2,1]}`, nil, nil, true))
	// t.Run("field with string key", runCase("field with string key", `{"a": 42}.a`, nil, nil, int64(42)))
}

func runCase(name, src string, input any, vars map[string]any, expect any) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		t.Parallel()
		state := vm.NewState()
		state.SetVars(vars)
		out, err := vm.Run([]byte(src), input)
		require.NoError(t, err, "program %q failed", src)
		require.Equal(t, expect, out)
	}
}
