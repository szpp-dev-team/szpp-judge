package grpc

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/judge/usecases/healthcheck"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
)

type healthcheckServiceServer struct {
	healthcheckInteractor *healthcheck.Interactor
}

func NewHealthcheckServiceServer() pb.HealthcheckServiceServer {
	return &healthcheckServiceServer{
		healthcheckInteractor: healthcheck.NewInteractor(),
	}
}

func (s *healthcheckServiceServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return s.healthcheckInteractor.Ping(req), nil
}
