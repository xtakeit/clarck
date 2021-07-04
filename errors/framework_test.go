package errors

import (
	"fmt"
	"testing"
)

func TestFrameworkErrorTypeIsErrorInterface(t *testing.T) {
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
	e := NewFrameworkError(1000, "frameworkError")
	if e.Code() != 1000 {
		t.Error("code 返回错误")
	}
}

func TestErrorMethod(t *testing.T) {
	e := NewFrameworkError(0, "abc", "efg")
	c := e.Error()
	if c != "0 [abc efg]" {
		t.Error("Error 返回值不符合预期")
	}

	if fmt.Sprint(e) != e.Error() {
		t.Error("Error 返回值不符合预期")
	}
}
