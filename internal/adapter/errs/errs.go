package errs

type Err struct {
	Code    int
	Message string
}

func (e Err) Error() string {
	return e.Message
}
