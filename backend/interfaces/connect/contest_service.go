package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/contests"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1/backendv1connect"
)

type contestServiceServer struct {
	interactor *contests.Interactor
}

func NewContestServiceServer(interactor *contests.Interactor) backendv1connect.ContestServiceHandler {
	return &contestServiceServer{interactor}
}

func (s *contestServiceServer) CreateContest(ctx context.Context, req *connect.Request[backendv1.CreateContestRequest]) (*connect.Response[backendv1.CreateContestResponse], error) {
	resp, err := s.interactor.CreateContest(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) GetContest(ctx context.Context, req *connect.Request[backendv1.GetContestRequest]) (*connect.Response[backendv1.GetContestResponse], error) {
	resp, err := s.interactor.GetContest(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) UpdateContest(ctx context.Context, req *connect.Request[backendv1.UpdateContestRequest]) (*connect.Response[backendv1.UpdateContestResponse], error) {
	resp, err := s.interactor.UpdateContest(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) ListContests(ctx context.Context, req *connect.Request[backendv1.ListContestsRequest]) (*connect.Response[backendv1.ListContestsResponse], error) {
	resp, err := s.interactor.ListContests(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) ListContestTasks(ctx context.Context, req *connect.Request[backendv1.ListContestTasksRequest]) (*connect.Response[backendv1.ListContestTasksResponse], error) {
	resp, err := s.interactor.ListContestTasks(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) GetContestTask(ctx context.Context, req *connect.Request[backendv1.GetContestTaskRequest]) (*connect.Response[backendv1.GetContestTaskResponse], error) {
	resp, err := s.interactor.GetContestTask(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) SyncContestTasks(ctx context.Context, req *connect.Request[backendv1.SyncContestTasksRequest]) (*connect.Response[backendv1.SyncContestTasksResponse], error) {
	resp, err := s.interactor.SyncContestTasks(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) GetMySubmissionStatuses(ctx context.Context, req *connect.Request[backendv1.GetMySubmissionStatusesRequest]) (*connect.Response[backendv1.GetMySubmissionStatusesResponse], error) {
	resp, err := s.interactor.GetMySubmissionStatuses(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) GetMyRegistrationStatus(ctx context.Context, req *connect.Request[backendv1.GetMyRegistrationStatusRequest]) (*connect.Response[backendv1.GetMyRegistrationStatusResponse], error) {
	resp, err := s.interactor.GetMyRegistrationStatus(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) GetStandings(ctx context.Context, req *connect.Request[backendv1.GetStandingsRequest]) (*connect.Response[backendv1.GetStandingsResponse], error) {
	resp, err := s.interactor.GetStandings(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *contestServiceServer) RegisterMe(ctx context.Context, req *connect.Request[backendv1.RegisterMeRequest]) (*connect.Response[backendv1.RegisterMeResponse], error) {
	resp, err := s.interactor.RegisterMe(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// UpdateAnswer implements backendv1.ContestServiceServer.
func (*contestServiceServer) UpdateAnswer(ctx context.Context, req *connect.Request[backendv1.UpdateAnswerRequest]) (*connect.Response[backendv1.UpdateAnswerResponse], error) {
	panic("unimplemented")
}

// CreateAnswer implements backendv1.ContestServiceServer.
func (*contestServiceServer) CreateAnswer(ctx context.Context, req *connect.Request[backendv1.CreateAnswerRequest]) (*connect.Response[backendv1.CreateAnswerResponse], error) {
	panic("unimplemented")
}

// CreateClarification implements backendv1.ContestServiceServer.
func (*contestServiceServer) CreateClarification(ctx context.Context, req *connect.Request[backendv1.CreateClarificationRequest]) (*connect.Response[backendv1.CreateClarificationResponse], error) {
	panic("unimplemented")
}

// DeleteAnswer implements backendv1.ContestServiceServer.
func (*contestServiceServer) DeleteAnswer(ctx context.Context, req *connect.Request[backendv1.DeleteAnswerRequest]) (*connect.Response[backendv1.DeleteAnswerResponse], error) {
	panic("unimplemented")
}

// DeleteClarification implements backendv1.ContestServiceServer.
func (*contestServiceServer) DeleteClarification(ctx context.Context, req *connect.Request[backendv1.DeleteClarificationRequest]) (*connect.Response[backendv1.DeleteClarificationResponse], error) {
	panic("unimplemented")
}

// ListClarifications implements backendv1.ContestServiceServer.
func (*contestServiceServer) ListClarifications(ctx context.Context, req *connect.Request[backendv1.ListClarificationsRequest]) (*connect.Response[backendv1.ListClarificationsResponse], error) {
	panic("unimplemented")
}
