package grpc

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/judge/usecases/judge"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type judgeServiceServer struct {
	interactor *judge.Interactor
}

func NewJudgeServiceServer() pb.JudgeServiceServer {
	return &judgeServiceServer{}
}

func (s *judgeServiceServer) Judge(req *pb.JudgeRequest, stream pb.JudgeService_JudgeServer) error {
	return s.interactor.Judge(req, stream)
}

func (*judgeServiceServer) Run(context.Context, *pb.RunRequest) (*pb.RunResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
