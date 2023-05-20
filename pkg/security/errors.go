package security

import "errors"

var (
	ErrFailedAuthentication = errors.New("incorrect principal username or password")
	ErrPasswordNotStrong    = errors.New("principal password is not strong enough")
)
