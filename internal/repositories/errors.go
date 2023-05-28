package repositories

import (
	"errors"

	"go.uber.org/multierr"
)

func ErrFindByUsername(errs ...error) error {
	return errors.New("db find by user name failed: " + multierr.Combine(errs...).Error())
}

func ErrExistsByUsername(errs ...error) error {
	return errors.New("db exists by user name failed: " + multierr.Combine(errs...).Error())
}
