package vmo

import (
	"crypto/sha256"
)

// FNV1a возвращает хеш uint64 по алгоритму FNV-1a (с анроллингом)
func FNV1a(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	var data []byte
	switch v := a[0].(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		this.Fail(StackUnsupported(a...))
		return
	}

	const (
		offset64 = 14695981039346656037
		prime64  = 1099511628211
	)

	var h uint64 = offset64
	n := len(data)
	i := 0

	// Loop unrolling: по 8 байт
	for ; i+8 <= n; i += 8 {
		h ^= uint64(data[i+0])
		h *= prime64
		h ^= uint64(data[i+1])
		h *= prime64
		h ^= uint64(data[i+2])
		h *= prime64
		h ^= uint64(data[i+3])
		h *= prime64
		h ^= uint64(data[i+4])
		h *= prime64
		h ^= uint64(data[i+5])
		h *= prime64
		h ^= uint64(data[i+6])
		h *= prime64
		h ^= uint64(data[i+7])
		h *= prime64
	}

	// Остаток
	for ; i < n; i++ {
		h ^= uint64(data[i])
		h *= prime64
	}

	this.Push(int64(h)) // в рамках стека используем int64
}

// SHA256 возвращает хеш от строки или байт в виде []byte (32 байта)
func SHA256(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	var data []byte
	switch v := a[0].(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		this.Fail(StackUnsupported(a...))
		return
	}

	hash := sha256.Sum256(data)
	this.Push(hash[:]) // срез от массива [32]byte
}
