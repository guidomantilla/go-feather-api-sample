package security

import "errors"

var (
	ErrFailedAuthentication = errors.New("incorrect principal username or password")
	ErrFailedAuthorization  = errors.New("invalid authorities")
	ErrPasswordNotStrong    = errors.New("principal password is not strong enough")
)
