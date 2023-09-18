package config

type Config struct {
	Host               *string `env:"HOST,default=localhost"`
	Port               *string `env:"PORT,default=8080"`
	TokenSignatureKey  *string `env:"TOKEN_SIGNATURE_KEY,default=SecretYouShouldHide"`
	ParamHolder        *string `env:"PARAM_HOLDER,default=named"`
	DatasourceDriver   *string `env:"DATASOURCE_DRIVER,required"`
	DatasourceUsername *string `env:"DATASOURCE_USERNAME,required"`
	DatasourcePassword *string `env:"DATASOURCE_PASSWORD,required"`
	DatasourceServer   *string `env:"DATASOURCE_SERVER,required"`
	DatasourceService  *string `env:"DATASOURCE_SERVICE,required"`
	DatasourceUrl      *string `env:"DATASOURCE_URL,required"`
}
