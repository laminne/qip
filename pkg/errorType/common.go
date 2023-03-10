package errorType

type ErrNotFound struct {
	code    string
	message string
}

func NewErrNotFound(t string, m string) error {
	return &ErrNotFound{
		code:    t,
		message: m,
	}
}

func (n *ErrNotFound) Error() string {
	return n.message
}

type ErrExists struct {
	code    string
	message string
}

func NewErrExists(t string, m string) error {
	return &ErrExists{
		code:    t,
		message: m,
	}
}

func (e *ErrExists) Error() string {
	return e.message
}
