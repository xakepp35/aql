package vm_test

import (
	"testing"

	"github.com/xakepp35/aql/pkg/aqc"
	"github.com/xakepp35/aql/pkg/vm"
)

var demoExpr = []byte("1 + 2 * 3")

func BenchmarkCompile(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := vm.NewCompiler()
		err := aqc.Compile(demoExpr, e)
		if err != nil {
			b.Fatalf("eval error: %v", err)
		}
	}
}

func BenchmarkCompileNoAlloc(b *testing.B) {
	prog := vm.NewCompiler()
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := aqc.Compile(demoExpr, prog)
		if err != nil {
			b.Fatalf("eval error: %v", err)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	e := vm.NewCompiler()
	err := aqc.Compile(demoExpr, e)
	if err != nil {
		b.Fatalf("parse error: %v", err)
	}
	prog := e.Program()
	this := vm.NewState()
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Init(this)
		this.Run(prog)
		if this.Err() != nil {
			b.Fatalf("eval error: %v", err)
		}
	}
}

func BenchmarkRunf(b *testing.B) {
	e := vm.NewCompiler()
	err := aqc.Compile(demoExpr, e)
	if err != nil {
		b.Fatalf("parse error: %v", err)
	}
	jit := e.JIT()
	this := vm.NewState()
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Init(this)
		this.Runf(jit)
		if this.Err() != nil {
			b.Fatalf("eval error: %v", err)
		}
	}
}
