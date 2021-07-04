package clarck_test

import (
	"testing"

	"github.com/anoxia/clarck"
	"github.com/anoxia/clarck/errors"
)

func TestNewFrameworkErrorMethod(t *testing.T) {
	e := clarck.NewFrameworkError(0, "")
	if _, ok := e.(*errors.FramworkError); !ok {
		t.Error("NewFrameworkError() 函数应返回 FrameworkError 类型")
	}
}

func TestNewConfigErrorMethod(t *testing.T) {
	e := clarck.NewConfigError(0, "")
	if _, ok := e.(*errors.ConfigError); !ok {
		t.Error("NewConfigError() 函数应返回 ConfigError 类型")
	}
}
