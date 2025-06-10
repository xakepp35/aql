// require/require.go
package require

import (
	"strings"
	"testing"
	"unsafe"
)

func Nil(t *testing.T, v any, msg ...string) {
	if v != nil {
		t.Helper()
		t.Fatalf("expected nil: %v %s", v, join(msg))
	}
}

func NotNil(t *testing.T, v any, msg ...string) {
	if v == nil {
		t.Helper()
		t.Fatalf("unexpected nil: %v %s", v, join(msg))
	}
}

func ErrorContains(t *testing.T, err error, str string, msg ...string) {
	if err == nil || !strings.Contains(err.Error(), str) {
		t.Helper()
		t.Fatalf("unexpected %v: %s %s", err, str, join(msg))
	}
}

// NoError fails the test if err is not nil.
func NoError(t *testing.T, err error, msg ...string) {
	if err != nil {
		t.Helper()
		t.Fatalf("unexpected error: %v %s", err, join(msg))
	}
}

// Error fails the test if err is nil.
func Error(t *testing.T, err error, msg ...string) {
	if err == nil {
		t.Helper()
		t.Fatalf("expected error, got nil %s", join(msg))
	}
}

// True fails the test if condition is false.
func True(t *testing.T, cond bool, msg ...string) {
	if !cond {
		t.Helper()
		t.Fatalf("expected true, got false %s", join(msg))
	}
}

// False fails the test if condition is true.
func False(t *testing.T, cond bool, msg ...string) {
	if cond {
		t.Helper()
		t.Fatalf("expected false, got true %s", join(msg))
	}
}

// Equal fails the test if a != b.
func Equal[T comparable](t *testing.T, a, b T, msg ...string) {
	if a != b {
		t.Helper()
		t.Fatalf("expected %v == %v %s", a, b, join(msg))
	}
}

// NotEqual fails the test if a == b.
func NotEqual[T comparable](t *testing.T, a, b T, msg ...string) {
	if a == b {
		t.Helper()
		t.Fatalf("expected %v != %v %s", a, b, join(msg))
	}
}

// BytesEqual fails the test if byte slices are not equal.
func BytesEqual[T ~byte](t *testing.T, a, b []T, msg ...string) {
	aa := unsafe.Slice((*byte)(unsafe.Pointer(&a[0])), len(a))
	bb := unsafe.Slice((*byte)(unsafe.Pointer(&a[0])), len(a))
	if string(aa) != string(bb) {
		t.Helper()
		t.Fatalf("expected bytes equal:\n%v\n!=\n%v\n%s", a, b, join(msg))
	}
}

func ErrorIs(t *testing.T, err, target error, msg ...string) {
	if err == nil {
		t.Helper()
		t.Fatalf("expected error, got nil %s", join(msg))
	}
	if !strings.Contains(err.Error(), target.Error()) {
		t.Helper()
		t.Fatalf("expected error %v to contain %v %s", err, target, join(msg))
	}
}

func join(msg []string) string {
	if len(msg) == 0 {
		return ""
	}
	return " - " + msg[0]
}
