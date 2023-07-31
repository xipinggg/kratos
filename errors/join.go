package errors

import (
	stderrors "errors"
)

func Join(errs ...error) error {
	return stderrors.Join(errs...)
}
