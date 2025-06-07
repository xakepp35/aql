package vmi

type Iterator interface {
	Next() bool
	Item() any
}
