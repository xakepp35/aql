package vmo

import (
	"encoding/hex"
	"strconv"
	"time"
)

// Bytes пытается привести значение к []byte
func Bytes(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	switch v := a[0].(type) {
	case string:
		this.Push([]byte(v))
	case []byte:
		this.Push(v)
	case int64:
		this.Push([]byte(strconv.FormatInt(v, 10)))
	case float64:
		this.Push([]byte(strconv.FormatFloat(v, 'f', -1, 64)))
	case bool:
		if v {
			this.Push([]byte("true"))
		} else {
			this.Push([]byte("false"))
		}
	default:
		this.Fail(StackUnsupported(a...))
	}
}

// String пытается привести значение к string
func String(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	switch v := a[0].(type) {
	case string:
		this.Push(v)
	case []byte:
		this.Push(string(v))
	case int64:
		this.Push(strconv.FormatInt(v, 10))
	case float64:
		this.Push(strconv.FormatFloat(v, 'f', -1, 64))
	case bool:
		if v {
			this.Push("true")
		} else {
			this.Push("false")
		}
	default:
		this.Fail(StackUnsupported(a...))
	}
}

// Float приводит значение к float64
func Float(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	switch v := a[0].(type) {
	case float64:
		this.Push(v)
	case int64:
		this.Push(float64(v))
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			this.Fail(err)
			return
		}
		this.Push(f)
	case []byte:
		f, err := strconv.ParseFloat(string(v), 64)
		if err != nil {
			this.Fail(err)
			return
		}
		this.Push(f)
	default:
		this.Fail(StackUnsupported(a...))
	}
}

// Int приводит значение к int64
func Int(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	switch v := a[0].(type) {
	case int64:
		this.Push(v)
	case float64:
		this.Push(int64(v))
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			this.Fail(err)
			return
		}
		this.Push(i)
	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			this.Fail(err)
			return
		}
		this.Push(i)
	default:
		this.Fail(StackUnsupported(a...))
	}
}

// Bool приводит значение к bool
func Bool(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	switch v := a[0].(type) {
	case bool:
		this.Push(v)
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			this.Fail(err)
			return
		}
		this.Push(b)
	case []byte:
		b, err := strconv.ParseBool(string(v))
		if err != nil {
			this.Fail(err)
			return
		}
		this.Push(b)
	default:
		this.Fail(StackUnsupported(a...))
	}
}

// Time парсит строку/[]byte в time.Time (RFC3339)
func Time(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	var str string
	switch v := a[0].(type) {
	case time.Time:
		this.Push(v)
		return
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		this.Fail(StackUnsupported(a...))
		return
	}
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		this.Fail(err)
		return
	}
	this.Push(t)
}

// Hex преобразует значение в шестнадцатеричную строку.
func Hex(this *VM) {
	a := this.Pops(1)
	if a == nil {
		this.Fail(StackUnderflow(this.Dump()...))
		return
	}

	switch v := a[0].(type) {
	case int64:
		this.Push(strconv.FormatInt(v, 16)) // hex числа
	case []byte:
		this.Push(hex.EncodeToString(v))
	case string:
		this.Push(hex.EncodeToString([]byte(v)))
	default:
		this.Fail(StackUnsupported(a...))
	}
}
