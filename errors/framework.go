package errors

import "fmt"

// 框架底层错误
type ErrorInterface interface {
	error
	Code() int
	Message() []string
}

// 框架底层错误
type FramworkError struct {
	code    int
	message []string
}

func (f *FramworkError) Error() string {
	return fmt.Sprintf("%d %s", f.code, f.message)
}

func (f *FramworkError) Code() int {
	return f.code
}

func (f *FramworkError) Message() []string {
	return f.message
}

func NewFrameworkError(code int, message ...string) ErrorInterface {
	return &FramworkError{code, message}
}
