package rec

type WrappedError interface {
	Unwrap() error
	error
}

type StackError interface {
	Stack() string
	WrappedError
}

type stackErrorImpl struct {
	error
	stack []byte
}

func (e *stackErrorImpl) Stack() string { return string(e.stack) }

func (e *stackErrorImpl) Error() string { return e.error.Error() }

func (e *stackErrorImpl) Unwrap() error { return e.error }
