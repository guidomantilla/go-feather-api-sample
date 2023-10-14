package repositories

import (
	"context"
	"fmt"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
	feather_sql_dao "github.com/guidomantilla/go-feather-sql/pkg/dao"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

type DefaultRepository struct {
	statementFindPrincipal string
	statementFindResource  string
	statementSaveResource  string
}

func NewDefaultRepository() *DefaultRepository {
	driverName := feather_sql.MysqlDriverName
	paramHolder := feather_sql.NamedParamHolder

	statementFindPrincipal, err := feather_sql.CreateSelectSQL("auth_principals", models.AuthPrincipal{}, driverName, paramHolder, feather_sql.PkColumnFilter)
	if err != nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: %s", "auth_principals", err.Error()))
	}

	statementFindResource, err := feather_sql.CreateSelectSQL("auth_resources", models.AuthResource{}, driverName, paramHolder, feather_sql.PkColumnFilter)
	if err != nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: %s", "auth_principals", err.Error()))
	}

	statementSaveResource, err := feather_sql.CreateInsertSQL("auth_resources", models.AuthResource{}, driverName, paramHolder)
	if err != nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: %s", "auth_principals", err.Error()))
	}

	return &DefaultRepository{
		statementFindPrincipal: statementFindPrincipal,
		statementFindResource:  statementFindResource,
		statementSaveResource:  statementSaveResource + " ON DUPLICATE KEY UPDATE name = :name, application = :application",
	}
}

func (repository *DefaultRepository) FindPrincipal(ctx context.Context, principal *models.AuthPrincipal) ([]models.AuthPrincipal, error) {
	return feather_sql_dao.QueryMany[models.AuthPrincipal](ctx, repository.statementFindPrincipal, principal)
}

//

func (repository *DefaultRepository) FindResource(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.QueryOne[models.AuthResource](ctx, repository.statementFindResource, resource)
}

func (repository *DefaultRepository) SaveResource(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.MutateOne[models.AuthResource](ctx, repository.statementSaveResource, resource)
}

func (repository *DefaultRepository) DeleteResource(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.MutateOne(ctx, "", resource)
}

//

func (repository *DefaultRepository) FindRole(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.QueryOne[models.AuthRole](ctx, "", role)
}

func (repository *DefaultRepository) SaveRole(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.MutateOne(ctx, "", role)
}

func (repository *DefaultRepository) DeleteRole(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.MutateOne(ctx, "", role)
}

//

func (repository *DefaultRepository) FindAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.QueryOne[models.AuthAccessControlList](ctx, "", acl)
}

func (repository *DefaultRepository) SaveAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.MutateOne(ctx, "", acl)
}

func (repository *DefaultRepository) DeleteAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.MutateOne(ctx, "", acl)
}

//

func (repository *DefaultRepository) FindUser(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.QueryOne[models.AuthUser](ctx, "", user)
}

func (repository *DefaultRepository) SaveUser(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.MutateOne(ctx, "", user)
}

func (repository *DefaultRepository) DeleteUser(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.MutateOne(ctx, "", user)
}
