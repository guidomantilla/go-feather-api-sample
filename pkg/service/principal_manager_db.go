package service

import (
	"context"
	"errors"
	"strings"

	feather_commons_collections "github.com/guidomantilla/go-feather-commons/pkg/collections"
	feather_commons_errors "github.com/guidomantilla/go-feather-commons/pkg/errors"
	feather_commons_util "github.com/guidomantilla/go-feather-commons/pkg/util"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	feather_sql_datasource "github.com/guidomantilla/go-feather-sql/pkg/datasource"
	"github.com/jmoiron/sqlx"

	"github.com/guidomantilla/go-feather-api-sample/pkg/config"
	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
	"github.com/guidomantilla/go-feather-api-sample/pkg/repositories"
)

var (
	_ feather_security.PrincipalManager = (*DBPrincipalManager)(nil)
)

type DBPrincipalManager struct {
	transactionHandler feather_sql_datasource.TransactionHandler
	passwordManager    feather_security.PasswordManager
	repository         repositories.Repository
}

func NewDBPrincipalManager(transactionHandler feather_sql_datasource.TransactionHandler, passwordManager feather_security.PasswordManager, repository repositories.Repository) *DBPrincipalManager {
	return &DBPrincipalManager{
		transactionHandler: transactionHandler,
		passwordManager:    passwordManager,
		repository:         repository,
	}
}

func (manager *DBPrincipalManager) Create(ctx context.Context, principal *feather_security.Principal) error {

	err := manager.Upsert(ctx, principal, "Create")
	if err != nil {
		return feather_commons_errors.ErrJoin(errors.New("error creating principal"), err)
	}

	return nil
}

func (manager *DBPrincipalManager) Update(ctx context.Context, principal *feather_security.Principal) error {

	err := manager.Upsert(ctx, principal, "Update")
	if err != nil {
		return feather_commons_errors.ErrJoin(errors.New("error creating principal"), err)
	}

	return nil
}

func (manager *DBPrincipalManager) Upsert(ctx context.Context, principal *feather_security.Principal, mode string) error {
	return manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {

		var err error

		queryBy := &models.AuthUser{
			Username: principal.Username,
		}
		err = manager.repository.FindUserById(ctx, queryBy)
		if mode == "Create" && err != nil {
			return err
		} else if mode == "Update" && err == nil {
			return err
		}

		authRole := &models.AuthRole{
			Name:    principal.Role,
			Enabled: feather_commons_util.TruePrt(),
		}
		if err = manager.repository.SaveRole(ctx, authRole); err != nil {
			return err
		}

		for _, resource := range principal.Resources {
			resourceParts := strings.Split(resource, " ")
			if len(resourceParts) != 3 {
				return errors.New("invalid resource format")
			}

			authResource := &models.AuthResource{
				Name:        feather_commons_util.ValueToPtr(resourceParts[2]),
				Application: feather_commons_util.ValueToPtr(resourceParts[0]),
				Enabled:     feather_commons_util.TruePrt(),
			}
			if err = manager.repository.SaveResource(ctx, authResource); err != nil {
				return err
			}

			authAccessControlList := &models.AuthAccessControlList{
				Role:       authRole.Name,
				Resource:   authResource.Name,
				Permission: feather_commons_util.ValueToPtr(resourceParts[1]),
				Enabled:    feather_commons_util.TruePrt(),
			}
			if err = manager.repository.SaveAccessControlList(ctx, authAccessControlList); err != nil {
				return err
			}
		}

		encodedPassword, err := manager.passwordManager.Encode(*principal.Password)
		if err != nil {
			return err
		}
		encodedPassphrase, err := manager.passwordManager.Encode(*principal.Passphrase)
		if err != nil {
			return err
		}

		authUser := &models.AuthUser{
			Username:   principal.Username,
			Role:       principal.Role,
			Password:   encodedPassword,
			Passphrase: encodedPassphrase,
			Enabled:    feather_commons_util.TruePrt(),
		}
		if err = manager.repository.SaveUser(ctx, authUser); err != nil {
			return err
		}

		return nil
	})
}

func (manager *DBPrincipalManager) Delete(ctx context.Context, username string) error {

	return manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {

		return nil
	})
}

func (manager *DBPrincipalManager) Find(ctx context.Context, username string) (*feather_security.Principal, error) {

	var err error
	var principal *feather_security.Principal
	err = manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {

		queryBy := &models.AuthPrincipal{
			Username:    feather_commons_util.ValueToPtr(username),
			Application: feather_commons_util.ValueToPtr(config.Application),
		}

		var authPrincipals []models.AuthPrincipal
		if authPrincipals, err = manager.repository.FindPrincipalById(ctx, queryBy); err != nil {
			return err
		}
		if len(authPrincipals) == 0 {
			return errors.New("no principal found")
		}

		resources := feather_commons_collections.Map[models.AuthPrincipal, string](authPrincipals, func(principal models.AuthPrincipal, _ int) string {
			return strings.Join([]string{*principal.Application, *principal.Permission, *principal.Resource}, " ")
		})

		principal = &feather_security.Principal{
			Username:           authPrincipals[0].Username,
			Role:               authPrincipals[0].Role,
			Password:           authPrincipals[0].Password,
			Passphrase:         authPrincipals[0].Passphrase,
			Enabled:            authPrincipals[0].Enabled,
			NonLocked:          feather_commons_util.TruePrt(),
			NonExpired:         feather_commons_util.TruePrt(),
			PasswordNonExpired: feather_commons_util.TruePrt(),
			SignUpDone:         feather_commons_util.TruePrt(),
			Resources:          resources,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return principal, nil
}

func (manager *DBPrincipalManager) Exists(ctx context.Context, username string) error {
	return nil
}

func (manager *DBPrincipalManager) ChangePassword(ctx context.Context, username string, password string) error {
	return manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {

		return nil
	})
}

func (manager *DBPrincipalManager) VerifyResource(ctx context.Context, username string, resource string) error {

	var err error
	var principal *feather_security.Principal
	if principal, err = manager.Find(ctx, username); err != nil {
		return err
	}

	if feather_commons_collections.Contains(principal.Resources, resource) {
		return nil
	}

	return feather_security.ErrAccountInvalidAuthorities
}
