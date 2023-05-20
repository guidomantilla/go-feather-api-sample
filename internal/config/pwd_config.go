package config

import (
	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
)

func InitPwd(_ environment.Environment) feather_security.PasswordManager {

	passwordGenerator := feather_security.NewDefaultPasswordGenerator()
	bcryptPasswordEncoder := feather_security.NewBcryptPasswordEncoder()
	passwordEncoder := feather_security.NewDelegatingPasswordEncoder(bcryptPasswordEncoder)
	passwordManager := feather_security.NewDefaultPasswordManager(passwordEncoder, passwordGenerator)
	return passwordManager
}
