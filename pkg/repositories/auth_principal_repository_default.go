package repositories

import (
	"context"
	"database/sql"
	"errors"

	feather_sql_dao "github.com/guidomantilla/go-feather-sql/pkg/dao"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

type DefaultAuthPrincipalRepository struct {
	statementFindByUsername   string
	statementExistsByUsername string
}

func NewDefaultAuthPrincipalRepository() *DefaultAuthPrincipalRepository {
	return &DefaultAuthPrincipalRepository{
		statementFindByUsername:   "SELECT username, role, resource, permission, password, passphrase, enabled FROM auth_principal WHERE username = ?",
		statementExistsByUsername: "SELECT distinct(username) FROM auth_principal WHERE username = ?",
	}
}

func (repository *DefaultAuthPrincipalRepository) ExistsByUsername(ctx context.Context, username string) error {

	var err error
	var exists string
	if err = feather_sql_dao.ReadRowContext(ctx, repository.statementExistsByUsername, username, &exists); err != nil {
		return ErrExistsByUsername(errors.New("auth_principal"), err)
	}

	if exists == "" {
		return ErrExistsByUsername(errors.New("auth_principal"), err)
	}

	return nil
}

func (repository *DefaultAuthPrincipalRepository) FindByUsername(ctx context.Context, username string) ([]models.AuthPrincipal, error) {

	var err error
	principals := make([]models.AuthPrincipal, 0)
	err = feather_sql_dao.Context(ctx, repository.statementFindByUsername, func(statement *sql.Stmt) error {

		var rows *sql.Rows
		if rows, err = statement.Query(username); err != nil {
			return err
		}
		defer feather_sql_dao.CloseResultSet(rows)

		for rows.Next() {

			var principal models.AuthPrincipal
			if err = rows.Scan(&principal.Username, &principal.Role, &principal.Resource, &principal.Permission, &principal.Password, &principal.Passphrase, &principal.Enabled); err != nil {
				return err
			}
			principals = append(principals, principal)
		}

		return nil
	})

	if err != nil {
		return nil, ErrFindByUsername(errors.New("auth_principal"), err)
	}

	return principals, nil
}
