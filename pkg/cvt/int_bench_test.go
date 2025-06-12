package cvt_test

import (
	"strconv"
	"testing"

	"github.com/xakepp35/aql/pkg/cvt"
)

var benchInput = []byte("-1234567890")
var result1 int64
var result2 int64

func BenchmarkParseInt64(b *testing.B) {
	var r int64
	for i := 0; i < b.N; i++ {
		r, _ = cvt.ParseInt64(benchInput)
	}
	result1 = r
}

func BenchmarkStrconvParseInt(b *testing.B) {
	bi := string(benchInput)
	var r int64
	for i := 0; i < b.N; i++ {
		r, _ = strconv.ParseInt(bi, 10, 64)
	}
	result2 = r
}
