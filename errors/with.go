package errors

import (
	"fmt"
	"runtime/debug"
)

func WithStack(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w\nstack: %s", err, debug.Stack())
}

func WithMessage(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\n%w", msg, err)
}
