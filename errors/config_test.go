package errors

import "testing"

func TestConfigErrorTypeIsFramworkInterface(t *testing.T) {
	e := NewConfigError(0, "")
	switch e.(type) {
	case ErrorInterface:
	default:
		t.Errorf("ConfigError 应为 ErrorInterface 实现类型")
	}
}
