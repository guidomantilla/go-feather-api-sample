package security

import (
	"context"

	"github.com/guidomantilla/go-feather-api-sample/pkg/validation"
)

type PrincipalManager interface {
	Create(ctx context.Context, principal *Principal) error
	Update(ctx context.Context, principal *Principal) error
	Delete(ctx context.Context, username string) error
	Find(ctx context.Context, username string) (*Principal, error)
	Exists(ctx context.Context, username string) error
	ChangePassword(ctx context.Context, principal *Principal) error
}

//

type GrantedAuthority struct {
	Role *string `json:"role,omitempty"`
}

type Principal struct {
	Username           *string             `json:"username,omitempty" binding:"required"`
	Password           *string             `json:"password,omitempty" binding:"required"`
	AccountNonExpired  *bool               `json:"account_non_expired,omitempty"`
	AccountNonLocked   *bool               `json:"account_non_locked,omitempty"`
	PasswordNonExpired *bool               `json:"password_non_expired,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	SignUpDone         *bool               `json:"signup_done,omitempty"`
	Authorities        *[]GrantedAuthority `json:"authorities,omitempty"`
}

func (principal *Principal) Validate() []error {

	var errors []error

	if err := validation.ValidateFieldIsRequired("this", "username", principal.Username); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldIsRequired("this", "password", principal.Password); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "account_non_expired", principal.AccountNonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "account_non_locked", principal.AccountNonLocked); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "password_non_expired", principal.PasswordNonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "enabled", principal.Enabled); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "signup_done", principal.SignUpDone); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateStructMustBeUndefined("this", "authorities", principal.Authorities); err != nil {
		errors = append(errors, err)
		return errors
	}

	return errors
}
