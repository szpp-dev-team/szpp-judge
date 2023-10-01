package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/user"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1/backendv1connect"
)

type userServiceServer struct {
	interactor *user.Interactor
}

func NewUserServiceServer(interactor *user.Interactor) backendv1connect.UserServiceHandler {
	return &userServiceServer{interactor}
}

func (s *userServiceServer) GetUser(ctx context.Context, req *connect.Request[backendv1.GetUserRequest]) (*connect.Response[backendv1.GetUserResponse], error) {
	resp, err := s.interactor.GetUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *connect.Request[backendv1.CreateUserRequest]) (*connect.Response[backendv1.CreateUserResponse], error) {
	resp, err := s.interactor.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *userServiceServer) ExistsUsername(ctx context.Context, req *connect.Request[backendv1.ExistsUsernameRequest]) (*connect.Response[backendv1.ExistsUsernameResponse], error) {
	resp, err := s.interactor.ExistsUsername(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *userServiceServer) ExistsEmail(ctx context.Context, req *connect.Request[backendv1.ExistsEmailRequest]) (*connect.Response[backendv1.ExistsEmailResponse], error) {
	resp, err := s.interactor.ExistsEmail(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
