package errors

import "fmt"

type ErrorInterface interface {
	error
	Code() int
	Message() []string
}

// 框架底层错误
type clarckError struct {
	code    int
	message []string
}

func (e *clarckError) Error() string {
	return fmt.Sprintf("%d %s", e.code, e.message)
}

func (e *clarckError) Code() int {
	return e.code
}

func (e *clarckError) Message() []string {
	return e.message
}
