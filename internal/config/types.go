package config

const (
	HostPort = "HOST_PORT"
)

var (
	EnvVarDefaultValuesMap = map[string]string{
		HostPort: ":8080",
	}
)
