package security

import (
	"context"
	"errors"

	"github.com/guidomantilla/go-feather-commons/pkg/security"
)

type InMemoryPrincipalService struct {
	repository      map[string]*Principal
	passwordManager security.PasswordManager
}

func NewInMemoryPrincipalService(passwordManager security.PasswordManager) *InMemoryPrincipalService {
	return &InMemoryPrincipalService{
		passwordManager: passwordManager,
		repository:      make(map[string]*Principal),
	}
}

func (service *InMemoryPrincipalService) Create(ctx context.Context, principal *Principal) error {

	var err error
	if err = service.Exists(ctx, *principal.Username); err == nil {
		return errors.New("username already exists")
	}

	if err = service.passwordManager.Validate(*principal.Password); err != nil {
		return err
	}

	if principal.Password, err = service.passwordManager.Encode(*principal.Password); err != nil {
		return err
	}

	service.repository[*principal.Username] = principal

	return nil
}

func (service *InMemoryPrincipalService) Update(ctx context.Context, principal *Principal) error {
	return service.Create(ctx, principal)
}

func (service *InMemoryPrincipalService) Delete(_ context.Context, username string) error {
	delete(service.repository, username)
	return nil
}

func (service *InMemoryPrincipalService) Find(_ context.Context, username string) (*Principal, error) {

	var ok bool
	var user *Principal
	if user, ok = service.repository[username]; !ok {
		return nil, errors.New("username not found")
	}
	return user, nil
}

func (service *InMemoryPrincipalService) Exists(_ context.Context, username string) error {

	var ok bool
	if _, ok = service.repository[username]; !ok {
		return errors.New("username not found")
	}
	return nil
}

func (service *InMemoryPrincipalService) ChangePassword(ctx context.Context, principal *Principal) error {

	var err error
	if err = service.Exists(ctx, *principal.Username); err != nil {
		return err
	}

	if err = service.passwordManager.Validate(*principal.Password); err != nil {
		return err
	}

	if principal.Password, err = service.passwordManager.Encode(*principal.Password); err != nil {
		return err
	}

	service.repository[*principal.Username] = principal

	return nil
}

func (service *InMemoryPrincipalService) Authenticate(_ context.Context, principal *Principal) error {

	var ok bool
	var user *Principal
	if user, ok = service.repository[*principal.Username]; !ok {
		return ErrFailedAuthentication
	}

	if user.Password == nil || *(user.Password) == "" {
		return ErrFailedAuthentication
	}

	var err error
	var matches *bool
	if matches, err = service.passwordManager.Matches(*(user.Password), *principal.Password); err != nil || !*(matches) {
		return ErrFailedAuthentication
	}

	var needsUpgrade *bool
	if needsUpgrade, err = service.passwordManager.UpgradeEncoding(*(user.Password)); err != nil || *(needsUpgrade) {
		return ErrFailedAuthentication
	}

	return nil
}
