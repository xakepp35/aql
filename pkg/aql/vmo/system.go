package vmo

import (
	"time"
)

//go:inline
func Nil(this *VM) {
	this.Push(nil)
}

//go:inline
func Now(this *VM) {
	this.Push(time.Now().UTC())
}
