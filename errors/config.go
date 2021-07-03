package errors

// 配置错误
type ConfigError struct {
	*FramworkError
}

func NewConfigError(code int, message ...string) FramworkErrorInterface {
	return &ConfigError{
		&FramworkError{code: code, message: message},
	}
}
