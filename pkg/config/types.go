package config

const (
	Application = "go-feather-api-sample"
	Version     = "v0.3.0"
)

type Config struct {
	Host               *string `env:"HOST,default=localhost"`
	HttpPort           *string `env:"HTTP_PORT,default=8080"`
	GrpcPort           *string `env:"GRPC_PORT,default=50051"`
	TokenSignatureKey  *string `env:"TOKEN_SIGNATURE_KEY,default=SecretYouShouldHide"`
	ParamHolder        *string `env:"PARAM_HOLDER,default=named"`
	DatasourceDriver   *string `env:"DATASOURCE_DRIVER,required"`
	DatasourceUsername *string `env:"DATASOURCE_USERNAME,required"`
	DatasourcePassword *string `env:"DATASOURCE_PASSWORD,required"`
	DatasourceServer   *string `env:"DATASOURCE_SERVER,required"`
	DatasourceService  *string `env:"DATASOURCE_SERVICE,required"`
	DatasourceUrl      *string `env:"DATASOURCE_URL,required"`
}
