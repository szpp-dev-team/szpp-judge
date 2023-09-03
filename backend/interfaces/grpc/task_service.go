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

// GetTask implements backendv1.TaskServiceServer.
func (*taskServiceServer) GetTask(context.Context, *backendv1.GetTaskRequest) (*backendv1.GetTaskResponse, error) {
	panic("unimplemented")
}

// UpdateTask implements backendv1.TaskServiceServer.
func (*taskServiceServer) UpdateTask(context.Context, *backendv1.UpdateTaskRequest) (*backendv1.UpdateTaskResponse, error) {
	panic("unimplemented")
}
