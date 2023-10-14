package repositories

import (
	"context"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

var (
	_ Repository = (*DefaultRepository)(nil)
)

type Repository interface {
	FindPrincipalByUsernameAndApplication(ctx context.Context, principal *models.AuthPrincipal) ([]models.AuthPrincipal, error)

	FindResource(ctx context.Context, resource *models.AuthResource) error
	SaveResource(ctx context.Context, resource *models.AuthResource) error
	DeleteResource(ctx context.Context, resource *models.AuthResource) error

	FindRole(ctx context.Context, role *models.AuthRole) error
	SaveRole(ctx context.Context, role *models.AuthRole) error
	DeleteRole(ctx context.Context, role *models.AuthRole) error

	FindAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error
	SaveAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error
	DeleteAccessControlList(ctx context.Context, acl *models.AuthAccessControlList) error

	FindUser(ctx context.Context, user *models.AuthUser) error
	SaveUser(ctx context.Context, user *models.AuthUser) error
	DeleteUser(ctx context.Context, user *models.AuthUser) error
}
