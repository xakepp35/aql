package main_test

// import (
// 	"testing"

// 	"github.com/itchyny/gojq"
// )

// func BenchmarkGojqRun(b *testing.B) {
// 	query, err := gojq.Parse("1 + 2 * 3") // эквивалент выражения: 1 + 2 * 3
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
// 		iter := code.Run(nil)
// 		_, _ = iter.Next()
// 	}
// }

// func BenchmarkGojqCompile(b *testing.B) {
// 	query, err := gojq.Parse("1 + 2 * 3") // эквивалент выражения: 1 + 2 * 3
// 	if err != nil {
// 		b.Fatalf("parse error: %v", err)
// 	}
// 	b.ReportAllocs()
// 	b.SetBytes(1)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_, _ = gojq.Compile(query)
// 	}
// }

// func BenchmarkGojqParse(b *testing.B) {
// 	b.ReportAllocs()
// 	b.SetBytes(1)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_, _ = gojq.Parse("1 + 2 * 3")
// 	}
// }

// func BenchmarkGojqFull(b *testing.B) {
// 	b.ReportAllocs()
// 	b.SetBytes(1)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		query, err := gojq.Parse("1 + 2 * 3") // эквивалент выражения: 1 + 2 * 3
// 		if err != nil {
// 			b.Fatalf("parse error: %v", err)
// 		}
// 		code, err := gojq.Compile(query)
// 		if err != nil {
// 			b.Fatalf("compile error: %v", err)
// 		}
// 		iter := code.Run(nil)
// 		_, _ = iter.Next()
// 	}
// }
