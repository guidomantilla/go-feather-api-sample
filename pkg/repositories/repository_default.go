package repositories

import (
	"context"

	feather_sql_dao "github.com/guidomantilla/go-feather-sql/pkg/dao"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

type DefaultRepository struct {
	queriesMap QueriesMap
}

func NewDefaultRepository(queriesMap QueriesMap) *DefaultRepository {
	return &DefaultRepository{
		queriesMap: queriesMap,
	}
}

func (repository *DefaultRepository) FindPrincipalById(ctx context.Context, principal *models.AuthPrincipal) ([]models.AuthPrincipal, error) {
	return feather_sql_dao.QueryMany[models.AuthPrincipal](ctx, repository.queriesMap["FindPrincipalById"], principal)
}

//

func (repository *DefaultRepository) FindResourceById(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.QueryOne[models.AuthResource](ctx, repository.queriesMap["FindResourceById"], resource)
}

func (repository *DefaultRepository) SaveResource(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.MutateOne[models.AuthResource](ctx, repository.queriesMap["SaveResource"], resource)
}

func (repository *DefaultRepository) DeleteResource(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.MutateOne(ctx, "", resource)
}

//

func (repository *DefaultRepository) FindRoleById(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.QueryOne[models.AuthRole](ctx, "", role)
}

func (repository *DefaultRepository) SaveRole(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.MutateOne(ctx, "", role)
}

func (repository *DefaultRepository) DeleteRole(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.MutateOne(ctx, "", role)
}

//

func (repository *DefaultRepository) FindAccessControlListById(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.QueryOne[models.AuthAccessControlList](ctx, "", acl)
}

func (repository *DefaultRepository) SaveAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.MutateOne(ctx, "", acl)
}

func (repository *DefaultRepository) DeleteAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.MutateOne(ctx, "", acl)
}

//

func (repository *DefaultRepository) FindUserById(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.QueryOne[models.AuthUser](ctx, "", user)
}

func (repository *DefaultRepository) SaveUser(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.MutateOne(ctx, "", user)
}

func (repository *DefaultRepository) DeleteUser(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.MutateOne(ctx, "", user)
}
