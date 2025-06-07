package vm_test

// import (
// 	"testing"

// 	"github.com/itchyny/gojq"
// )

// func BenchmarkGojqRun(b *testing.B) {
// 	query, err := gojq.Parse(". + 2 * 3") // эквивалент выражения: . + 2 * 3
// 	if err != nil {
// 		b.Fatalf("parse error: %v", err)
// 	}

// 	code, err := gojq.Compile(query)
// 	if err != nil {
// 		b.Fatalf("compile error: %v", err)
// 	}
// 	b.ReportAllocs()
// 	b.SetBytes(1)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		iter := code.Run(1)
// 		_, _ = iter.Next()
// 		// v, ok := iter.Next()
// 		// if !ok {
// 		// 	b.Fatalf("no result")
// 		// }
// 		// if err, ok := v.(error); ok {
// 		// 	b.Fatalf("runtime error: %v", err)
// 		// }
// 		// if v != 7 {
// 		// 	b.Fatalf("wrong result: got %v", v)
// 		// }
// 	}
// }
