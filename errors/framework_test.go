package errors

import "testing"

func TestFrameworkErrorTypeIsFramworkInterface(t *testing.T) {
	e := NewFrameworkError(0, "")
	switch e.(type) {
	case ErrorInterface:
	default:
		t.Errorf("FrameworkError 应为 ErrorInterface 实现类型")
	}
}

func TestFrameworkErrorTypeIsError(t *testing.T) {
	e := NewFrameworkError(0, "")
	if _, ok := e.(error); !ok {
		t.Errorf("FrameworkError 类型不是 error")
	}
}

func TestCodeMethod(t *testing.T) {
	frameworkError := NewFrameworkError(1000, "frameworkError")
	if frameworkError.Code() != 1000 {
		t.Error("code 返回错误")
	}
}
