package grpc

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/backend/usecases/judge"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type judgeServiceServer struct {
	interactor *judge.Interactor
}

func NewJudgeServiceServer(interactor *judge.Interactor) backendv1.JudgeServiceServer {
	return &judgeServiceServer{interactor}
}

func (s *judgeServiceServer) Submit(ctx context.Context, req *backendv1.SubmitRequest) (*backendv1.SubmitResponse, error) {
	return s.interactor.PostSubmit(ctx, req)
}

func (s *judgeServiceServer) GetSubmissionDetail(ctx context.Context, req *backendv1.GetSubmissionDetailRequest) (*backendv1.GetSubmissionDetailResponse, error) {
	return s.interactor.GetSubmissionDetail(ctx, req)
}

func (s *judgeServiceServer) ListSubmissions(ctx context.Context, req *backendv1.ListSubmissionsRequest) (*backendv1.ListSubmissionsResponse, error) {
	if req.ContestId == nil {
		return nil, status.Error(codes.InvalidArgument, "currently contestID must be set")
	}
	return s.interactor.ListSubmissions(ctx, req)
}

func (s *judgeServiceServer) GetJudgeProgress(ctx context.Context, req *backendv1.GetJudgeProgressRequest) (*backendv1.GetJudgeProgressResponse, error) {
	return s.interactor.GetJudgeProgress(ctx, req)
}
