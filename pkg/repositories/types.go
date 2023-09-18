package repositories

import (
	"context"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

type AuthResourceRepository interface {
	Create(ctx context.Context, authResource *models.AuthResource) error
	Update(ctx context.Context, authResource *models.AuthResource) error
	DeleteById(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (*models.AuthResource, error)
	FindAll(ctx context.Context) ([]models.AuthResource, error)
	FindByName(ctx context.Context, username string) (*models.AuthResource, error)
}

type AuthRoleRepository interface {
	Create(ctx context.Context, authRole *models.AuthRole) error
	Update(ctx context.Context, authRole *models.AuthRole) error
	DeleteById(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (*models.AuthRole, error)
	FindAll(ctx context.Context) ([]models.AuthRole, error)
	FindByName(ctx context.Context, username string) (*models.AuthRole, error)
}

type AuthAccessControlListRepository interface {
	Create(ctx context.Context, authAccessControlList *models.AuthAccessControlList) error
	Update(ctx context.Context, authAccessControlList *models.AuthAccessControlList) error
	DeleteById(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (*models.AuthAccessControlList, error)
	FindAll(ctx context.Context) ([]models.AuthAccessControlList, error)
	FindByRole(ctx context.Context, role string) ([]models.AuthAccessControlList, error)
}

type AuthUserRepository interface {
	Create(ctx context.Context, authUser *models.AuthUser) error
	Update(ctx context.Context, authUser *models.AuthUser) error
	DeleteById(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (*models.AuthUser, error)
	FindAll(ctx context.Context) ([]models.AuthUser, error)
	FindByRole(ctx context.Context, role string) ([]models.AuthUser, error)
	FindByUsername(ctx context.Context, username string) (*models.AuthUser, error)
}

type AuthPrincipalRepository interface {
	FindByUsername(ctx context.Context, username string) ([]models.AuthPrincipal, error)
	ExistsByUsername(ctx context.Context, username string) error
}
