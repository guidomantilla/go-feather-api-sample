package repositories

import (
	"errors"

	feather_commons_errors "github.com/guidomantilla/go-feather-commons/pkg/errors"
)

func ErrFindPrincipal(errs ...error) error {
	return errors.New("find principal failed: " + feather_commons_errors.ErrJoin(errs...).Error())
}
