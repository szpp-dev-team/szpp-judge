// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: backend/v1/task_service.proto

package backendv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TaskService_CreateTask_FullMethodName       = "/backend.v1.TaskService/CreateTask"
	TaskService_GetTask_FullMethodName          = "/backend.v1.TaskService/GetTask"
	TaskService_UpdateTask_FullMethodName       = "/backend.v1.TaskService/UpdateTask"
	TaskService_GetTestcaseSets_FullMethodName  = "/backend.v1.TaskService/GetTestcaseSets"
	TaskService_SyncTestcaseSets_FullMethodName = "/backend.v1.TaskService/SyncTestcaseSets"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	// Task を作成する
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	// 指定された Task を取得する
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	// Task を更新する
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	// TestcaseSet の一覧を取得する。また、Testcase の一覧も取得する。
	GetTestcaseSets(ctx context.Context, in *GetTestcaseSetsRequest, opts ...grpc.CallOption) (*GetTestcaseSetsResponse, error)
	// TestcaseSet を同期する。全てのリソースは上書きされ、このリクエストに含まれないリソースは削除される。
	SyncTestcaseSets(ctx context.Context, in *SyncTestcaseSetsRequest, opts ...grpc.CallOption) (*SyncTestcaseSetsResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_GetTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_UpdateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTestcaseSets(ctx context.Context, in *GetTestcaseSetsRequest, opts ...grpc.CallOption) (*GetTestcaseSetsResponse, error) {
	out := new(GetTestcaseSetsResponse)
	err := c.cc.Invoke(ctx, TaskService_GetTestcaseSets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SyncTestcaseSets(ctx context.Context, in *SyncTestcaseSetsRequest, opts ...grpc.CallOption) (*SyncTestcaseSetsResponse, error) {
	out := new(SyncTestcaseSetsResponse)
	err := c.cc.Invoke(ctx, TaskService_SyncTestcaseSets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations should embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	// Task を作成する
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	// 指定された Task を取得する
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	// Task を更新する
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	// TestcaseSet の一覧を取得する。また、Testcase の一覧も取得する。
	GetTestcaseSets(context.Context, *GetTestcaseSetsRequest) (*GetTestcaseSetsResponse, error)
	// TestcaseSet を同期する。全てのリソースは上書きされ、このリクエストに含まれないリソースは削除される。
	SyncTestcaseSets(context.Context, *SyncTestcaseSetsRequest) (*SyncTestcaseSetsResponse, error)
}

// UnimplementedTaskServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskServiceServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskServiceServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskServiceServer) GetTestcaseSets(context.Context, *GetTestcaseSetsRequest) (*GetTestcaseSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestcaseSets not implemented")
}
func (UnimplementedTaskServiceServer) SyncTestcaseSets(context.Context, *SyncTestcaseSetsRequest) (*SyncTestcaseSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncTestcaseSets not implemented")
}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTestcaseSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestcaseSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTestcaseSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetTestcaseSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTestcaseSets(ctx, req.(*GetTestcaseSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SyncTestcaseSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncTestcaseSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SyncTestcaseSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_SyncTestcaseSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SyncTestcaseSets(ctx, req.(*SyncTestcaseSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.v1.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskService_UpdateTask_Handler,
		},
		{
			MethodName: "GetTestcaseSets",
			Handler:    _TaskService_GetTestcaseSets_Handler,
		},
		{
			MethodName: "SyncTestcaseSets",
			Handler:    _TaskService_SyncTestcaseSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/v1/task_service.proto",
}
