package vm

type Stream chan any

//go:inline
func (s Stream) Send(v any) {
	s <- v
}

//go:inline
func (s Stream) Recv() any {
	return <-s
}
