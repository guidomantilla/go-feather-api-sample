package repositories

import (
	"context"

	feather_boot "github.com/guidomantilla/go-feather-boot/pkg/boot"
	feather_sql_dao "github.com/guidomantilla/go-feather-sql/pkg/dao"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

type DefaultRepository struct {
	findPrincipalById       string
	findResourceById        string
	saveResource            string
	deleteResource          string
	saveRole                string
	deleteRole              string
	saveAccessControlList   string
	deleteAccessControlList string
	saveUser                string
	deleteUser              string
}

func NewDefaultRepository(databaseConfig *feather_boot.DatabaseConfig) *DefaultRepository {
	return &DefaultRepository{
		findPrincipalById:       feather_sql.CreateSelectSQL("auth_principals", models.AuthPrincipal{}, databaseConfig.Driver, databaseConfig.ParamHolder, feather_sql.PkColumnFilter),
		findResourceById:        feather_sql.CreateSelectSQL("auth_resources", models.AuthResource{}, databaseConfig.Driver, databaseConfig.ParamHolder, feather_sql.PkColumnFilter),
		saveResource:            feather_sql.CreateInsertSQL("auth_resources", models.AuthResource{}, databaseConfig.Driver, databaseConfig.ParamHolder) + " ON DUPLICATE KEY UPDATE name = :name, application = :application",
		deleteResource:          feather_sql.CreateDeleteSQL("auth_resources", models.AuthResource{}, databaseConfig.Driver, databaseConfig.ParamHolder, feather_sql.PkColumnFilter),
		saveRole:                feather_sql.CreateInsertSQL("auth_roles", models.AuthRole{}, databaseConfig.Driver, databaseConfig.ParamHolder) + " ON DUPLICATE KEY UPDATE name = :name",
		deleteRole:              feather_sql.CreateDeleteSQL("auth_roles", models.AuthRole{}, databaseConfig.Driver, databaseConfig.ParamHolder, feather_sql.PkColumnFilter),
		saveAccessControlList:   feather_sql.CreateInsertSQL("auth_access_control_list", models.AuthAccessControlList{}, databaseConfig.Driver, databaseConfig.ParamHolder) + " ON DUPLICATE KEY UPDATE role = :role, resource = :resource, permission = :permission",
		deleteAccessControlList: feather_sql.CreateDeleteSQL("auth_access_control_list", models.AuthAccessControlList{}, databaseConfig.Driver, databaseConfig.ParamHolder, feather_sql.PkColumnFilter),
		saveUser:                feather_sql.CreateInsertSQL("auth_users", models.AuthUser{}, databaseConfig.Driver, databaseConfig.ParamHolder) + " ON DUPLICATE KEY UPDATE username = :username, role = :role, password = :password, passphrase = :passphrase",
		deleteUser:              feather_sql.CreateDeleteSQL("auth_users", models.AuthUser{}, databaseConfig.Driver, databaseConfig.ParamHolder, feather_sql.PkColumnFilter),
	}
}

func (repository *DefaultRepository) FindPrincipalById(ctx context.Context, principal *models.AuthPrincipal) ([]models.AuthPrincipal, error) {
	return feather_sql_dao.QueryMany[models.AuthPrincipal](ctx, repository.findPrincipalById, principal)
}

//

func (repository *DefaultRepository) FindResourceById(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.QueryOne[models.AuthResource](ctx, repository.findResourceById, resource)
}

func (repository *DefaultRepository) SaveResource(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.MutateOne[models.AuthResource](ctx, repository.saveResource, resource)
}

func (repository *DefaultRepository) DeleteResource(ctx context.Context, resource *models.AuthResource) error {
	return feather_sql_dao.MutateOne(ctx, repository.deleteResource, resource)
}

//

func (repository *DefaultRepository) FindRoleById(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.QueryOne[models.AuthRole](ctx, "", role)
}

func (repository *DefaultRepository) SaveRole(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.MutateOne(ctx, repository.saveRole, role)
}

func (repository *DefaultRepository) DeleteRole(ctx context.Context, role *models.AuthRole) error {
	return feather_sql_dao.MutateOne(ctx, repository.deleteRole, role)
}

//

func (repository *DefaultRepository) FindAccessControlListById(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.QueryOne[models.AuthAccessControlList](ctx, "", acl)
}

func (repository *DefaultRepository) SaveAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.MutateOne(ctx, repository.saveAccessControlList, acl)
}

func (repository *DefaultRepository) DeleteAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error {
	return feather_sql_dao.MutateOne(ctx, repository.deleteAccessControlList, acl)
}

//

func (repository *DefaultRepository) FindUserById(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.QueryOne[models.AuthUser](ctx, "", user)
}

func (repository *DefaultRepository) SaveUser(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.MutateOne(ctx, repository.saveUser, user)
}

func (repository *DefaultRepository) DeleteUser(ctx context.Context, user *models.AuthUser) error {
	return feather_sql_dao.MutateOne(ctx, repository.deleteUser, user)
}
