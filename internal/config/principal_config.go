package config

import (
	"context"

	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	feather_commons_util "github.com/guidomantilla/go-feather-commons/pkg/util"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	"go.uber.org/zap"
)

func InitPrincipal(_ environment.Environment, passwordManager feather_security.PasswordManager) feather_security.PrincipalManager {

	root := &feather_security.Principal{
		Username:           feather_commons_util.ValueToPtr("root"),
		Password:           feather_commons_util.ValueToPtr("RaveN123qweasd*+"),
		AccountNonExpired:  feather_commons_util.ValueToPtr(true),
		AccountNonLocked:   feather_commons_util.ValueToPtr(true),
		PasswordNonExpired: feather_commons_util.ValueToPtr(true),
		Enabled:            feather_commons_util.ValueToPtr(true),
		SignUpDone:         feather_commons_util.ValueToPtr(true),
		Authorities: feather_commons_util.ValueToPtr([]feather_security.GrantedAuthority{
			{
				Role: feather_commons_util.ValueToPtr("rol01"),
			},
			{
				Role: feather_commons_util.ValueToPtr("rol02"),
			},
			{
				Role: feather_commons_util.ValueToPtr("rol03"),
			},
		}),
	}
	principalManager := feather_security.NewInMemoryPrincipalManager(passwordManager)

	var err error
	if err = principalManager.Create(context.Background(), root); err != nil {
		zap.L().Fatal(err.Error())
		return nil
	}
	return principalManager
}
