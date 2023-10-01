package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/auth"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1/backendv1connect"
)

type authServiceServer struct {
	interactor *auth.Interactor
}

func NewAuthServiceServer(interactor *auth.Interactor) backendv1connect.AuthServiceHandler {
	return &authServiceServer{interactor}
}

func (s *authServiceServer) Login(ctx context.Context, req *connect.Request[backendv1.LoginRequest]) (*connect.Response[backendv1.LoginResponse], error) {
	resp, err := s.interactor.Login(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *authServiceServer) Logout(ctx context.Context, req *connect.Request[backendv1.LogoutRequest]) (*connect.Response[backendv1.LogoutResponse], error) {
	resp, err := s.interactor.Logout(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *authServiceServer) RefreshAccessToken(ctx context.Context, req *connect.Request[backendv1.RefreshAccessTokenRequest]) (*connect.Response[backendv1.RefreshAccessTokenResponse], error) {
	resp, err := s.interactor.RefreshAccessToken(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
