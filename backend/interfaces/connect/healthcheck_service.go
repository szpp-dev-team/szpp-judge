package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/healthcheck"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1/backendv1connect"
)

type healthcheckServiceServer struct {
	healthcheckInteractor *healthcheck.Interactor
}

func NewHealthcheckServiceServer() backendv1connect.HealthcheckServiceHandler {
	return &healthcheckServiceServer{
		healthcheckInteractor: healthcheck.NewInteractor(),
	}
}

func (s *healthcheckServiceServer) Ping(ctx context.Context, req *connect.Request[backendv1.PingRequest]) (*connect.Response[backendv1.PingResponse], error) {
	resp := s.healthcheckInteractor.Ping(req.Msg)
	return connect.NewResponse(resp), nil
}
