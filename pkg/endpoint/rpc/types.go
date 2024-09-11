package rpc

import (
	"context"
	"strings"

	feather_web_rest "github.com/guidomantilla/go-feather-lib/pkg/rest"
	feather_security "github.com/guidomantilla/go-feather-lib/pkg/security"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	_ ApiSampleServer = (*ApiSampleGrpcServer)(nil)
)

type ApiSampleGrpcServer struct {
	authenticationService feather_security.AuthenticationService
	authorizationService  feather_security.AuthorizationService
	principalManager      feather_security.PrincipalManager
}

func NewApiSampleGrpcServer(authenticationService feather_security.AuthenticationService, authorizationService feather_security.AuthorizationService, principalManager feather_security.PrincipalManager) *ApiSampleGrpcServer {
	return &ApiSampleGrpcServer{
		authenticationService: authenticationService,
		authorizationService:  authorizationService,
		principalManager:      principalManager,
	}
}

func (server *ApiSampleGrpcServer) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {

	principal := &feather_security.Principal{
		Username: &request.Username,
		Password: &request.Password,
	}
	var err error

	if errs := server.authenticationService.Validate(principal); errs != nil {
		ex := feather_web_rest.BadRequestException("error validating the principal", errs...)
		return nil, status.Errorf(codes.InvalidArgument, ex.Message)
	}

	if err = server.authenticationService.Authenticate(ctx, principal); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		return nil, status.Errorf(codes.Unauthenticated, ex.Message)
	}

	return &LoginResponse{
		Username:  *principal.Username,
		Role:      *principal.Role,
		Resources: principal.Resources,
		Token:     *principal.Token,
	}, nil
}

func (server *ApiSampleGrpcServer) GetPrincipal(ctx context.Context, _ *emptypb.Empty) (*Principal, error) {

	var ok bool
	var md metadata.MD
	if md, ok = metadata.FromIncomingContext(ctx); !ok {
		ex := feather_web_rest.UnauthorizedException("failed to get metadata")
		return nil, status.Errorf(codes.Unauthenticated, ex.Message)
	}

	bearer := md.Get("Authorization")
	if len(bearer) == 0 {
		ex := feather_web_rest.UnauthorizedException("invalid authorization header")
		return nil, status.Errorf(codes.Unauthenticated, ex.Message)
	}

	if !strings.HasPrefix(bearer[0], "Bearer ") {
		ex := feather_web_rest.UnauthorizedException("invalid authorization header")
		return nil, status.Errorf(codes.Unauthenticated, ex.Message)
	}

	splits := strings.Split(bearer[0], " ")
	if len(splits) != 2 {
		ex := feather_web_rest.UnauthorizedException("invalid authorization header")
		return nil, status.Errorf(codes.Unauthenticated, ex.Message)
	}

	var err error
	var principal *feather_security.Principal
	ctxWithResource := context.WithValue(ctx, feather_security.ResourceCtxKey{}, strings.Join([]string{"GET", "/principal"}, " "))
	if principal, err = server.authorizationService.Authorize(ctxWithResource, splits[1]); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		return nil, status.Errorf(codes.Unauthenticated, ex.Message)
	}

	if principal, err = server.principalManager.Find(ctx, *principal.Username); err != nil {
		ex := feather_web_rest.InternalServerErrorException(err.Error())
		return nil, status.Errorf(codes.Internal, ex.Message)
	}
	return &Principal{
		Username:           *principal.Username,
		Role:               *principal.Role,
		Password:           *principal.Password,
		Passphrase:         *principal.Passphrase,
		Enabled:            *principal.Enabled,
		NonLocked:          *principal.NonLocked,
		NonExpired:         *principal.NonExpired,
		PasswordNonExpired: *principal.PasswordNonExpired,
		SignupDone:         *principal.SignUpDone,
		Resources:          principal.Resources,
		Token:              *principal.Token,
	}, nil
}

func (server *ApiSampleGrpcServer) mustEmbedUnimplementedApiSampleServer() {
}
