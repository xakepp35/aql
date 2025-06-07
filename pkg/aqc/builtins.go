package aqc

import "github.com/xakepp35/aql/pkg/vm/op"

var Builtins = map[string]op.Code{
	"sum":   op.Add,
	"avg":   op.Avg,
	"count": op.Count,
	"min":   op.Min,
	"max":   op.Max,
	"topk":  op.Topk,
	// converters
	"to_int":    op.ToInt,
	"to_float":  op.ToFloat,
	"to_string": op.ToString,
	"to_bool":   op.ToBool,
	"to_time":   op.ToTime,
	// crypto
	"fnv1a": op.FNV1a,
	// formats
	"pack_kv":     op.PackKV,
	"unpack_kv":   op.UnpackKV,
	"pack_json":   op.PackJSON,
	"unpack_json": op.UnpackJSON,
}
