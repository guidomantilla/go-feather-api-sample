package config

import (
	"os"

	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	"github.com/guidomantilla/go-feather-commons/pkg/properties"
)

func InitEnv(cmdArgs []string) environment.Environment {

	// Load CMD and OS variables
	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OsPropertySourceName, properties.NewDefaultProperties(properties.FromArray(&osArgs)))
	cmdSource := properties.NewDefaultPropertySource(CmdPropertySourceName, properties.NewDefaultProperties(properties.FromArray(&cmdArgs)))
	environment := environment.NewDefaultEnvironment(environment.WithPropertySources(osSource, cmdSource))
	return environment
}
