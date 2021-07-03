package errors

// 框架底层错误
type FramworkError struct {
	*clarckError
}

func NewFrameworkError(code int, message ...string) ErrorInterface {
	return &FramworkError{code, message}
}
