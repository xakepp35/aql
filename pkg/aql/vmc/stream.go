package vmc

type RecvStream chan any

//go:inline
func (s RecvStream) Recv() any {
	return <-s
}

type SendStream chan any

//go:inline
func (s SendStream) Send(v any) {
	s <- v
}
