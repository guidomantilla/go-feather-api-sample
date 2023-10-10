package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	feather_commons_collections "github.com/guidomantilla/go-feather-commons/pkg/collections"
	feather_commons_util "github.com/guidomantilla/go-feather-commons/pkg/util"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	feather_sql_datasource "github.com/guidomantilla/go-feather-sql/pkg/datasource"

	"github.com/guidomantilla/go-feather-api-sample/pkg/config"
	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
	"github.com/guidomantilla/go-feather-api-sample/pkg/repositories"
)

var (
	_ feather_security.PrincipalManager = (*DBPrincipalManager)(nil)
)

type DBPrincipalManager struct {
	transactionHandler      feather_sql_datasource.TransactionHandler
	authPrincipalRepository repositories.AuthPrincipalRepository
}

func NewDBPrincipalManager(transactionHandler feather_sql_datasource.TransactionHandler, authPrincipalRepository repositories.AuthPrincipalRepository) *DBPrincipalManager {
	return &DBPrincipalManager{
		transactionHandler:      transactionHandler,
		authPrincipalRepository: authPrincipalRepository,
	}
}

func (manager *DBPrincipalManager) Create(ctx context.Context, principal *feather_security.Principal) error {
	var err error
	err = manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (manager *DBPrincipalManager) Update(ctx context.Context, principal *feather_security.Principal) error {
	var err error
	err = manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (manager *DBPrincipalManager) Delete(ctx context.Context, username string) error {
	var err error
	err = manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (manager *DBPrincipalManager) Find(ctx context.Context, username string) (*feather_security.Principal, error) {

	var err error
	var principal *feather_security.Principal
	err = manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {

		var authPrincipals []models.AuthPrincipal
		if authPrincipals, err = manager.authPrincipalRepository.FindByUsername(ctx, username); err != nil {
			return err
		}
		authPrincipals = feather_commons_collections.Filter[models.AuthPrincipal](authPrincipals, func(principal models.AuthPrincipal, _ int) bool {
			return (*principal.Application) == config.Application
		})

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

	if err := manager.authPrincipalRepository.ExistsByUsername(ctx, username); err != nil {
		return err
	}

	return nil
}

func (manager *DBPrincipalManager) ChangePassword(ctx context.Context, username string, password string) error {
	var err error
	err = manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {

		return nil
	})
	if err != nil {
		return err
	}

	return nil
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
