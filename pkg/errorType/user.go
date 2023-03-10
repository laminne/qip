package errorType

type ErrMissingPassword struct {
	code    string
	message string
}

func NewErrMissingPassword(c string, m string) error {
	return &ErrMissingPassword{
		code:    c,
		message: m,
	}
}

func (e *ErrMissingPassword) Error() string {
	return e.message
}
