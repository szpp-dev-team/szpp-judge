package grpc

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/judge"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1/backendv1connect"
)

type judgeServiceServer struct {
	interactor *judge.Interactor
}

func NewJudgeServiceServer(interactor *judge.Interactor) backendv1connect.JudgeServiceHandler {
	return &judgeServiceServer{interactor}
}

func (s *judgeServiceServer) Submit(ctx context.Context, req *connect.Request[backendv1.SubmitRequest]) (*connect.Response[backendv1.SubmitResponse], error) {
	resp, err := s.interactor.PostSubmit(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *judgeServiceServer) GetSubmissionDetail(ctx context.Context, req *connect.Request[backendv1.GetSubmissionDetailRequest]) (*connect.Response[backendv1.GetSubmissionDetailResponse], error) {
	resp, err := s.interactor.GetSubmissionDetail(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *judgeServiceServer) ListSubmissions(ctx context.Context, req *connect.Request[backendv1.ListSubmissionsRequest]) (*connect.Response[backendv1.ListSubmissionsResponse], error) {
	if req.Msg.ContestId == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("currently contestID must be set"))
	}
	resp, err := s.interactor.ListSubmissions(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *judgeServiceServer) GetJudgeProgress(ctx context.Context, req *connect.Request[backendv1.GetJudgeProgressRequest]) (*connect.Response[backendv1.GetJudgeProgressResponse], error) {
	resp, err := s.interactor.GetJudgeProgress(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
