package grpc

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/backend/usecases/user"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

type userServiceServer struct {
	interactor *user.Interactor
}

func (s *userServiceServer) GetUser(ctx context.Context, req *backendv1.GetUserRequest) (*backendv1.GetUserResponse, error) {
	return s.interactor.GetUser(ctx, req)
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *backendv1.CreateUserRequest) (*backendv1.CreateUserResponse, error) {
	return s.interactor.CreateUser(ctx, req)
}