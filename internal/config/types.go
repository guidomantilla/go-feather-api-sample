package config

const (
	AppName               = "go-feather-api-sample"
	OsPropertySourceName  = "OS_PROPERTY_SOURCE_NAME"
	CmdPropertySourceName = "CMD_PROPERTY_SOURCE_NAME"
	HostPort              = "HOST_PORT"
	TokenSignatureKey     = "TOKEN_SIGNATURE_KEY"
)

var (
	EnvVarDefaultValuesMap = map[string]string{
		HostPort:          ":8080",
		TokenSignatureKey: "SecretYouShouldHide",
	}
)
