package vmo

var Builtins = NewBuiltins()

func NewBuiltins() Functions {
	return Functions{
		// ---- converters ----
		"int":    Int,
		"float":  Float,
		"bytes":  Bytes,
		"string": String,
		"bool":   Bool,
		"time":   Time,
		"hex":    Hex,
		// ---- stats ----
		"sum":   Add,
		"avg":   Avg,
		"count": Count,
		"min":   Min,
		"max":   Max,
		// "topk":  TopK,
		// ---- crypto ----
		"fnv1a":  FNV1a,
		"sha256": SHA256,
		// ---- formats ----
		// "pack_kv":     PackKV,
		// "unpack_kv":   UnpackKV,
		// "pack_json":   PackJSON,
		// "unpack_json": UnpackJSON,
		// ---- iterators ----
		// Iter      // convert stack object to Iterator
		// HashProbe //	map + key → hashSingleIter (O(1) точка)
		// RBSeek    // root + lo + hi → rbSeekIter
		// Limit     // + n → limitIter
		// Batch     //iterator + n → batchIter (для групповго commit)
		// //
		// From
		// Store
		// Upsert
		// Update
		// Delete
		// Txn
	}
}
