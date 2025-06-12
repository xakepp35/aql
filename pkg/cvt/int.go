package cvt

import (
	"errors"
)

var ErrInvalidInt = errors.New("invalid int64")

func ParseInt64(b []byte) (int64, error) {
	if len(b) == 0 {
		return 0, ErrInvalidInt
	}

	var neg bool
	var i int
	if b[0] == '-' {
		neg = true
		i++
	} else if b[0] == '+' {
		i++
	}

	var x int64
	for ; i < len(b); i++ {
		d := b[i]
		if d < '0' || d > '9' {
			return 0, ErrInvalidInt
		}
		x = x*10 + int64(d-'0')
	}

	if neg {
		x = -x
	}
	return x, nil
}
