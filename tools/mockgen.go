package tools

//go:generate mockgen -package rpc -destination $PROJECT_HOME/pkg/endpoint/rpc/mocks.go -source $PROJECT_HOME/pkg/endpoint/rpc/api_grpc.pb.go
//go:generate mockgen -package rest -destination $PROJECT_HOME/pkg/endpoint/rest/mocks.go -source $PROJECT_HOME/pkg/endpoint/rest/types.go
//go:generate mockgen -package repositories -destination $PROJECT_HOME/pkg/repositories/mocks.go -source $PROJECT_HOME/pkg/repositories/types.go
