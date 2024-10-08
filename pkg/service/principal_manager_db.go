package service

import (
	"context"
	"errors"
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/config"
	feather_sql_datasource "github.com/guidomantilla/go-feather-lib/pkg/datasource"
	feather_security "github.com/guidomantilla/go-feather-lib/pkg/security"
	feather_commons_util "github.com/guidomantilla/go-feather-lib/pkg/util"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

var (
	_ feather_security.PrincipalManager = (*DBPrincipalManager)(nil)
)

type DBPrincipalManager struct {
	transactionHandler feather_sql_datasource.TransactionHandler
	passwordManager    feather_security.PasswordManager
}

func NewDBPrincipalManager(transactionHandler feather_sql_datasource.TransactionHandler, passwordManager feather_security.PasswordManager) *DBPrincipalManager {
	return &DBPrincipalManager{
		transactionHandler: transactionHandler,
		passwordManager:    passwordManager,
	}
}

func (manager *DBPrincipalManager) Create(ctx context.Context, principal *feather_security.Principal) error {
	return manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		return tx.Create(principal).Error
	})
}

func (manager *DBPrincipalManager) Update(ctx context.Context, principal *feather_security.Principal) error {
	return manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		return tx.Save(principal).Error
	})
}

func (manager *DBPrincipalManager) Upsert(ctx context.Context, principal *feather_security.Principal, mode string) error {
	return manager.transactionHandler.HandleTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {

		var err error

		queryBy := &models.AuthUser{
			Username: principal.Username,
		}
		err = manager.repository.FindUserById(ctx, queryBy)
		if mode == "Create" && err == nil {
			return errors.New("principal already exists")
		} else if mode == "Update" && err != nil {
			return errors.New("principal does not exists")
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

		queryBy := &models.AuthPrincipal{
			Username:    feather_commons_util.ValueToPtr(username),
			Application: feather_commons_util.ValueToPtr(config.Application),
		}

		var authPrincipals []models.AuthPrincipal
		authPrincipals, err := manager.repository.FindPrincipalById(ctx, queryBy)
		if err != nil {
			return err
		}
		if len(authPrincipals) == 0 {
			return errors.New("principal does not exists")
		}

		if err = manager.repository.DeleteUser(ctx, &models.AuthUser{Username: &username}); err != nil {
			return err
		}

		for _, principal := range authPrincipals {
			if err = manager.repository.DeleteAccessControlList(ctx, &models.AuthAccessControlList{Role: principal.Role, Resource: principal.Resource, Permission: principal.Permission}); err != nil {
				return err
			}

			if err = manager.repository.DeleteResource(ctx, &models.AuthResource{Name: principal.Resource}); err != nil {
				return err
			}
		}

		if err = manager.repository.DeleteRole(ctx, &models.AuthRole{Name: authPrincipals[0].Role}); err != nil {
			return err
		}

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
			return errors.New("principal does not exists")
		}

		resources := feather_commons_streams.Map[models.AuthPrincipal, string](feather_commons_streams.Build(authPrincipals), func(principal models.AuthPrincipal) string {
			return strings.Join([]string{*principal.Application, *principal.Permission, *principal.Resource}, " ")
		}).ToArray()

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

		var err error
		authUser := &models.AuthUser{
			Username: &username,
		}
		if err = manager.repository.FindUserById(ctx, authUser); err != nil {
			return errors.New("principal does not exists")
		}

		authUser.Password, err = manager.passwordManager.Encode(password)
		if err != nil {
			return err
		}

		if err = manager.repository.SaveUser(ctx, authUser); err != nil {
			return err
		}

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
