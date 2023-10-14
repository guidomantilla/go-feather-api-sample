package repositories

import (
	"context"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

var (
	_ Repository = (*DefaultRepository)(nil)
)

type Repository interface {
	FindPrincipalById(ctx context.Context, principal *models.AuthPrincipal) ([]models.AuthPrincipal, error)

	FindResourceById(ctx context.Context, resource *models.AuthResource) error
	SaveResource(ctx context.Context, resource *models.AuthResource) error
	DeleteResource(ctx context.Context, resource *models.AuthResource) error

	FindRoleById(ctx context.Context, role *models.AuthRole) error
	SaveRole(ctx context.Context, role *models.AuthRole) error
	DeleteRole(ctx context.Context, role *models.AuthRole) error

	FindAccessControlListById(ctx context.Context, acl *models.AuthAccessControlList) error
	SaveAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error
	DeleteAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error

	FindUserById(ctx context.Context, user *models.AuthUser) error
	SaveUser(ctx context.Context, user *models.AuthUser) error
	DeleteUser(ctx context.Context, user *models.AuthUser) error
}
