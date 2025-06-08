package vm

type ErrStatus struct {
	err error
}

//go:inline
func (s *ErrStatus) Fail(err error) {
	s.err = err
}

//go:inline
func (s ErrStatus) Status() error {
	return s.err
}
