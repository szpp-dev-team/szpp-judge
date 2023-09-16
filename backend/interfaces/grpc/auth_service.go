package grpc

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/backend/usecases/auth"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

type authServiceServer struct {
	interactor *auth.Interactor
}

func NewAuthServiceServer(interactor *auth.Interactor) backendv1.AuthServiceServer {
	return &authServiceServer{interactor}
}

func (s *authServiceServer) Login(ctx context.Context, req *backendv1.LoginRequest) (*backendv1.LoginResponse, error) {
	return s.interactor.Login(ctx, req, s.interactor.Secret)
}
