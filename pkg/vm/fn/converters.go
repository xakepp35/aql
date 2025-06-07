package fn

import (
	"strconv"
	"time"

	"github.com/xakepp35/aql/pkg/vmi"
)

// ToString пытается привести значение к string
func ToString(this vmi.State) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
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
		this.SetErr(StackUnsupported(a...))
	}
}

// ToFloat приводит значение к float64
func ToFloat(this vmi.State) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
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
			this.SetErr(err)
			return
		}
		this.Push(f)
	case []byte:
		f, err := strconv.ParseFloat(string(v), 64)
		if err != nil {
			this.SetErr(err)
			return
		}
		this.Push(f)
	default:
		this.SetErr(StackUnsupported(a...))
	}
}

// ToInt приводит значение к int64
func ToInt(this vmi.State) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
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
			this.SetErr(err)
			return
		}
		this.Push(i)
	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			this.SetErr(err)
			return
		}
		this.Push(i)
	default:
		this.SetErr(StackUnsupported(a...))
	}
}

// ToBool приводит значение к bool
func ToBool(this vmi.State) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
		return
	}

	switch v := a[0].(type) {
	case bool:
		this.Push(v)
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			this.SetErr(err)
			return
		}
		this.Push(b)
	case []byte:
		b, err := strconv.ParseBool(string(v))
		if err != nil {
			this.SetErr(err)
			return
		}
		this.Push(b)
	default:
		this.SetErr(StackUnsupported(a...))
	}
}

// ToTime парсит строку/[]byte в time.Time (RFC3339)
func ToTime(this vmi.State) {
	a := this.Args(1)
	if a == nil {
		this.SetErr(ErrStackUnderflow)
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
		this.SetErr(StackUnsupported(a...))
		return
	}
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		this.SetErr(err)
		return
	}
	this.Push(t)
}
