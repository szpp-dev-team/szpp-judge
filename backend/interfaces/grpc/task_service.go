package grpc

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/backend/usecases/tasks"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

type taskServiceServer struct {
	interactor *tasks.Interactor
}

func NewTaskServiceServer(interactor *tasks.Interactor) backendv1.TaskServiceServer {
	return &taskServiceServer{interactor}
}

func (s *taskServiceServer) CreateTask(ctx context.Context, req *backendv1.CreateTaskRequest) (*backendv1.CreateTaskResponse, error) {
	return s.interactor.CreateTask(ctx, req)
}

func (s *taskServiceServer) GetTask(ctx context.Context, req *backendv1.GetTaskRequest) (*backendv1.GetTaskResponse, error) {
	return s.interactor.GetTask(ctx, req)
}

func (s *taskServiceServer) UpdateTask(ctx context.Context, req *backendv1.UpdateTaskRequest) (*backendv1.UpdateTaskResponse, error) {
	return s.interactor.UpdateTask(ctx, req)
}

func (s *taskServiceServer) GetTestcaseSets(ctx context.Context, req *backendv1.GetTestcaseSetsRequest) (*backendv1.GetTestcaseSetsResponse, error) {
	return s.interactor.GetTestcaseSets(ctx, req)
}

func (s *taskServiceServer) SyncTestcaseSets(ctx context.Context, req *backendv1.SyncTestcaseSetsRequest) (*backendv1.SyncTestcaseSetsResponse, error) {
	return s.interactor.SyncTestcaseSets(ctx, req)
}

// CreateAnswer implements backendv1.TaskServiceServer.
func (*taskServiceServer) CreateAnswer(context.Context, *backendv1.CreateAnswerRequest) (*backendv1.CreateAnswerResponse, error) {
	panic("unimplemented")
}

// CreateClarification implements backendv1.TaskServiceServer.
func (*taskServiceServer) CreateClarification(context.Context, *backendv1.CreateClarificationRequest) (*backendv1.CreateClarificationResponse, error) {
	panic("unimplemented")
}

// DeleteAnswer implements backendv1.TaskServiceServer.
func (*taskServiceServer) DeleteAnswer(context.Context, *backendv1.DeleteAnswerRequest) (*backendv1.DeleteAnswerResponse, error) {
	panic("unimplemented")
}

// DeleteClarification implements backendv1.TaskServiceServer.
func (*taskServiceServer) DeleteClarification(context.Context, *backendv1.DeleteClarificationRequest) (*backendv1.DeleteClarificationResponse, error) {
	panic("unimplemented")
}

// GetAnswer implements backendv1.TaskServiceServer.
func (*taskServiceServer) GetAnswer(context.Context, *backendv1.GetAnswerRequest) (*backendv1.GetAnswerResponse, error) {
	panic("unimplemented")
}

// GetClarification implements backendv1.TaskServiceServer.
func (*taskServiceServer) GetClarification(context.Context, *backendv1.GetClarificationRequest) (*backendv1.GetClarificationResponse, error) {
	panic("unimplemented")
}

// GetJudgeProgress implements backendv1.TaskServiceServer.
func (*taskServiceServer) GetJudgeProgress(context.Context, *backendv1.GetJudgeProgressRequest) (*backendv1.GetJudgeProgressResponse, error) {
	panic("unimplemented")
}

// GetSubmissionDetail implements backendv1.TaskServiceServer.
func (*taskServiceServer) GetSubmissionDetail(context.Context, *backendv1.GetSubmissionDetailRequest) (*backendv1.GetSubmissionDetailResponse, error) {
	panic("unimplemented")
}

// ListClarifications implements backendv1.TaskServiceServer.
func (*taskServiceServer) ListClarifications(context.Context, *backendv1.ListClarificationsRequest) (*backendv1.ListClarificationsResponse, error) {
	panic("unimplemented")
}

// ListSubmissions implements backendv1.TaskServiceServer.
func (*taskServiceServer) ListSubmissions(context.Context, *backendv1.ListSubmissionsRequest) (*backendv1.ListSubmissionsResponse, error) {
	panic("unimplemented")
}

// Submit implements backendv1.TaskServiceServer.
func (*taskServiceServer) Submit(context.Context, *backendv1.SubmitRequest) (*backendv1.SubmitResponse, error) {
	panic("unimplemented")
}

// UpdateAnswer implements backendv1.TaskServiceServer.
func (*taskServiceServer) UpdateAnswer(context.Context, *backendv1.UpdateAnswerRequest) (*backendv1.UpdateAnswerResponse, error) {
	panic("unimplemented")
}

// UpdateClarification implements backendv1.TaskServiceServer.
func (*taskServiceServer) UpdateClarification(context.Context, *backendv1.UpdateClarificationRequest) (*backendv1.UpdateClarificationResponse, error) {
	panic("unimplemented")
}
