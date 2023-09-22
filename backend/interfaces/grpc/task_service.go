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
