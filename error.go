package clarck

import (
	"github.com/anoxia/clarck/errors"
)

func NewFrameworkError(code int, message ...string) errors.FramworkErrorInterface {
	return errors.NewFrameworkError(code, message...)
}

func NewConfigError(code int, message ...string) errors.FramworkErrorInterface {
	return errors.NewConfigError(code, message...)
}
