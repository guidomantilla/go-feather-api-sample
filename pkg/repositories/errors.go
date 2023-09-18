package repositories

import (
	"errors"
)

func ErrFindByUsername(errs ...error) error {
	return errors.New("db find by user name failed: " + errors.Join(errs...).Error())
}

func ErrExistsByUsername(errs ...error) error {
	return errors.New("db exists by user name failed: " + errors.Join(errs...).Error())
}
