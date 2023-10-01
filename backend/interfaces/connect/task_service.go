package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/tasks"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1/backendv1connect"
)

type taskServiceServer struct {
	interactor *tasks.Interactor
}

func NewTaskServiceServer(interactor *tasks.Interactor) backendv1connect.TaskServiceHandler {
	return &taskServiceServer{interactor}
}

func (s *taskServiceServer) CreateTask(ctx context.Context, req *connect.Request[backendv1.CreateTaskRequest]) (*connect.Response[backendv1.CreateTaskResponse], error) {
	resp, err := s.interactor.CreateTask(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *taskServiceServer) GetTask(ctx context.Context, req *connect.Request[backendv1.GetTaskRequest]) (*connect.Response[backendv1.GetTaskResponse], error) {
	resp, err := s.interactor.GetTask(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *taskServiceServer) UpdateTask(ctx context.Context, req *connect.Request[backendv1.UpdateTaskRequest]) (*connect.Response[backendv1.UpdateTaskResponse], error) {
	resp, err := s.interactor.UpdateTask(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *taskServiceServer) GetTestcaseSets(ctx context.Context, req *connect.Request[backendv1.GetTestcaseSetsRequest]) (*connect.Response[backendv1.GetTestcaseSetsResponse], error) {
	resp, err := s.interactor.GetTestcaseSets(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *taskServiceServer) SyncTestcaseSets(ctx context.Context, req *connect.Request[backendv1.SyncTestcaseSetsRequest]) (*connect.Response[backendv1.SyncTestcaseSetsResponse], error) {
	resp, err := s.interactor.SyncTestcaseSets(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
