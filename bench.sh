#!/bin/bash
# go test -bench ^BenchmarkEmit$ -run ^$ -cpuprofile cpu.prof github.com/xakepp35/aql/pkg/ast/expr
go test -benchmem -run=^$ -bench ^BenchmarkMain$ -cpuprofile cpu.prof github.com/xakepp35/aql/cmd/aq
go tool pprof -http=localhost:8080 ./pkg/ast/expr.test cpu.prof