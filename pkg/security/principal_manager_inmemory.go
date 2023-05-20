package security

import (
	"context"
	"errors"
	"reflect"

	"github.com/guidomantilla/go-feather-commons/pkg/security"
)

type InMemoryPrincipalManager struct {
	repository      map[string]*Principal
	passwordManager security.PasswordManager
}

func NewInMemoryPrincipalManager(passwordManager security.PasswordManager) *InMemoryPrincipalManager {
	return &InMemoryPrincipalManager{
		passwordManager: passwordManager,
		repository:      make(map[string]*Principal),
	}
}

func (manager *InMemoryPrincipalManager) Create(ctx context.Context, principal *Principal) error {

	var err error
	if err = manager.Exists(ctx, *principal.Username); err == nil {
		return errors.New("username already exists")
	}

	if err = manager.passwordManager.Validate(*principal.Password); err != nil {
		return err
	}

	if principal.Password, err = manager.passwordManager.Encode(*principal.Password); err != nil {
		return err
	}

	manager.repository[*principal.Username] = principal

	return nil
}

func (manager *InMemoryPrincipalManager) Update(ctx context.Context, principal *Principal) error {
	return manager.Create(ctx, principal)
}

func (manager *InMemoryPrincipalManager) Delete(_ context.Context, username string) error {
	delete(manager.repository, username)
	return nil
}

func (manager *InMemoryPrincipalManager) Find(_ context.Context, username string) (*Principal, error) {

	var ok bool
	var user *Principal
	if user, ok = manager.repository[username]; !ok {
		return nil, errors.New("username not found")
	}
	return user, nil
}

func (manager *InMemoryPrincipalManager) Exists(_ context.Context, username string) error {

	var ok bool
	if _, ok = manager.repository[username]; !ok {
		return errors.New("username not found")
	}
	return nil
}

func (manager *InMemoryPrincipalManager) ChangePassword(ctx context.Context, principal *Principal) error {

	var err error
	if err = manager.Exists(ctx, *principal.Username); err != nil {
		return err
	}

	if err = manager.passwordManager.Validate(*principal.Password); err != nil {
		return err
	}

	if principal.Password, err = manager.passwordManager.Encode(*principal.Password); err != nil {
		return err
	}

	manager.repository[*principal.Username] = principal

	return nil
}

func (manager *InMemoryPrincipalManager) Authenticate(ctx context.Context, principal *Principal) error {

	var err error
	var user *Principal
	if user, err = manager.Find(ctx, *principal.Username); err != nil {
		return ErrFailedAuthentication
	}

	if user.Password == nil || *(user.Password) == "" {
		return ErrFailedAuthentication
	}

	var matches *bool
	if matches, err = manager.passwordManager.Matches(*(user.Password), *principal.Password); err != nil || !*(matches) {
		return ErrFailedAuthentication
	}

	var needsUpgrade *bool
	if needsUpgrade, err = manager.passwordManager.UpgradeEncoding(*(user.Password)); err != nil || *(needsUpgrade) {
		return ErrFailedAuthentication
	}

	principal.Password = nil
	principal.Authorities = user.Authorities

	return nil
}

func (manager *InMemoryPrincipalManager) Authorize(ctx context.Context, principal *Principal) error {

	var err error
	var user *Principal
	if user, err = manager.Find(ctx, *principal.Username); err != nil {
		return ErrFailedAuthorization
	}

	if !reflect.DeepEqual(user.Authorities, principal.Authorities) {
		return ErrFailedAuthorization
	}

	return nil
}
