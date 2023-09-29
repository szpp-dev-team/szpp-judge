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

func (s *contestServiceServer) CreateClarification(ctx context.Context, req *backendv1.CreateClarificationRequest) (*backendv1.CreateClarificationResponse, error) {
	return s.interactor.CreateClarification(ctx, req)
}

func (s *contestServiceServer) ListClarifications(ctx context.Context, req *backendv1.ListClarificationsRequest) (*backendv1.ListClarificationsResponse, error) {
	return s.interactor.ListClarifications(ctx, req)
}

func (s *contestServiceServer) DeleteClarification(ctx context.Context, req *backendv1.DeleteClarificationRequest) (*backendv1.DeleteClarificationResponse, error) {
	return s.interactor.DeleteClarification(ctx, req)
}

func (s *contestServiceServer) CreateAnswer(ctx context.Context, req *backendv1.CreateAnswerRequest) (*backendv1.CreateAnswerResponse, error) {
	return s.interactor.CreateAnswer(ctx, req)
}

func (s *contestServiceServer) UpdateAnswer(ctx context.Context, req *backendv1.UpdateAnswerRequest) (*backendv1.UpdateAnswerResponse, error) {
	return s.interactor.UpdateAnswer(ctx, req)
}

func (s *contestServiceServer) DeleteAnswer(ctx context.Context, req *backendv1.DeleteAnswerRequest) (*backendv1.DeleteAnswerResponse, error) {
	return s.interactor.DeleteAnswer(ctx, req)
}

// GetMySubmissionStatuses implements backendv1.ContestServiceServer.
func (*contestServiceServer) GetMySubmissionStatuses(context.Context, *backendv1.GetMySubmissionStatusesRequest) (*backendv1.GetMySubmissionStatusesResponse, error) {
	panic("unimplemented")
}

// GetStandings implements backendv1.ContestServiceServer.
func (*contestServiceServer) GetStandings(context.Context, *backendv1.GetStandingsRequest) (*backendv1.GetStandingsResponse, error) {
	panic("unimplemented")
}

// RegisterMe implements backendv1.ContestServiceServer.
func (*contestServiceServer) RegisterMe(context.Context, *backendv1.RegisterMeRequest) (*backendv1.RegisterMeResponse, error) {
	panic("unimplemented")
}

// SyncContestTasks implements backendv1.ContestServiceServer.
func (*contestServiceServer) SyncContestTasks(context.Context, *backendv1.SyncContestTasksRequest) (*backendv1.SyncContestTasksResponse, error) {
	panic("unimplemented")
}
