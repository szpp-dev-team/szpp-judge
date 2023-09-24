package grpc

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/backend/usecases/contests"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

type contestServiceServer struct {
	interactor *contests.Interactor
}

func NewContestServiceServer(interactor *contests.Interactor) backendv1.ContestServiceServer {
	return &contestServiceServer{interactor}
}

func (s *contestServiceServer) CreateContest(ctx context.Context, req *backendv1.CreateContestRequest) (*backendv1.CreateContestResponse, error) {
	return s.interactor.CreateContest(ctx, req)
}

func (s *contestServiceServer) GetContest(ctx context.Context, req *backendv1.GetContestRequest) (*backendv1.GetContestResponse, error) {
	return s.interactor.GetContest(ctx, req)
}

func (s *contestServiceServer) UpdateContest(ctx context.Context, req *backendv1.UpdateContestRequest) (*backendv1.UpdateContestResponse, error) {
	return s.interactor.UpdateContest(ctx, req)
}

func (s *contestServiceServer) ListContests(ctx context.Context, req *backendv1.ListContestsRequest) (*backendv1.ListContestsResponse, error) {
	return s.interactor.ListContests(ctx, req)
}

func (s *contestServiceServer) ListContestTasks(ctx context.Context, req *backendv1.ListContestTasksRequest) (*backendv1.ListContestTasksResponse, error) {
	return s.interactor.ListContestTasks(ctx, req)
}

func (s *contestServiceServer) GetContestTask(ctx context.Context, req *backendv1.GetContestTaskRequest) (*backendv1.GetContestTaskResponse, error) {
	return s.interactor.GetContestTask(ctx, req)
}

func (s *contestServiceServer) SyncContestTasks(ctx context.Context, req *backendv1.SyncContestTasksRequest) (*backendv1.SyncContestTasksResponse, error) {
	return s.interactor.SyncContestTasks(ctx, req)
}

func (s *contestServiceServer) GetMySubmissionStatuses(ctx context.Context, req *backendv1.GetMySubmissionStatusesRequest) (*backendv1.GetMySubmissionStatusesResponse, error) {
	return s.interactor.GetMySubmissionStatuses(ctx, req)
}

// GetStandings implements backendv1.ContestServiceServer.
func (*contestServiceServer) GetStandings(context.Context, *backendv1.GetStandingsRequest) (*backendv1.GetStandingsResponse, error) {
	panic("unimplemented")
}

func (s *contestServiceServer) RegisterMe(ctx context.Context, req *backendv1.RegisterMeRequest) (*backendv1.RegisterMeResponse, error) {
	return s.interactor.RegisterMe(ctx, req)
}

// UpdateAnswer implements backendv1.ContestServiceServer.
func (*contestServiceServer) UpdateAnswer(context.Context, *backendv1.UpdateAnswerRequest) (*backendv1.UpdateAnswerResponse, error) {
	panic("unimplemented")
}

// CreateAnswer implements backendv1.ContestServiceServer.
func (*contestServiceServer) CreateAnswer(context.Context, *backendv1.CreateAnswerRequest) (*backendv1.CreateAnswerResponse, error) {
	panic("unimplemented")
}

// CreateClarification implements backendv1.ContestServiceServer.
func (*contestServiceServer) CreateClarification(context.Context, *backendv1.CreateClarificationRequest) (*backendv1.CreateClarificationResponse, error) {
	panic("unimplemented")
}

// DeleteAnswer implements backendv1.ContestServiceServer.
func (*contestServiceServer) DeleteAnswer(context.Context, *backendv1.DeleteAnswerRequest) (*backendv1.DeleteAnswerResponse, error) {
	panic("unimplemented")
}

// DeleteClarification implements backendv1.ContestServiceServer.
func (*contestServiceServer) DeleteClarification(context.Context, *backendv1.DeleteClarificationRequest) (*backendv1.DeleteClarificationResponse, error) {
	panic("unimplemented")
}

// ListClarifications implements backendv1.ContestServiceServer.
func (*contestServiceServer) ListClarifications(context.Context, *backendv1.ListClarificationsRequest) (*backendv1.ListClarificationsResponse, error) {
	panic("unimplemented")
}
