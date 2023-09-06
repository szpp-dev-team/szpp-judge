// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: backend/v1/services.proto

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
	UserService_GetUser_FullMethodName    = "/backend.v1.UserService/GetUser"
	UserService_CreateUser_FullMethodName = "/backend.v1.UserService/CreateUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 指定された User を取得する
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// User を新たに作成する
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 指定された User を取得する
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// User を新たに作成する
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/v1/services.proto",
}

const (
	AuthService_Login_FullMethodName  = "/backend.v1.AuthService/Login"
	AuthService_Logout_FullMethodName = "/backend.v1.AuthService/Logout"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// ログイン
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// ログアウト
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, AuthService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations should embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// ログイン
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// ログアウト
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
}

// UnimplementedAuthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/v1/services.proto",
}

const (
	TaskService_CreateTask_FullMethodName          = "/backend.v1.TaskService/CreateTask"
	TaskService_GetTask_FullMethodName             = "/backend.v1.TaskService/GetTask"
	TaskService_UpdateTask_FullMethodName          = "/backend.v1.TaskService/UpdateTask"
	TaskService_CreateClarification_FullMethodName = "/backend.v1.TaskService/CreateClarification"
	TaskService_GetClarification_FullMethodName    = "/backend.v1.TaskService/GetClarification"
	TaskService_ListClarifications_FullMethodName  = "/backend.v1.TaskService/ListClarifications"
	TaskService_UpdateClarification_FullMethodName = "/backend.v1.TaskService/UpdateClarification"
	TaskService_DeleteClarification_FullMethodName = "/backend.v1.TaskService/DeleteClarification"
	TaskService_CreateAnswer_FullMethodName        = "/backend.v1.TaskService/CreateAnswer"
	TaskService_GetAnswer_FullMethodName           = "/backend.v1.TaskService/GetAnswer"
	TaskService_UpdateAnswer_FullMethodName        = "/backend.v1.TaskService/UpdateAnswer"
	TaskService_DeleteAnswer_FullMethodName        = "/backend.v1.TaskService/DeleteAnswer"
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
	// Clarification を作成する
	CreateClarification(ctx context.Context, in *CreateClarificationRequest, opts ...grpc.CallOption) (*CreateClarificationResponse, error)
	// 指定された Clarification を取得する
	GetClarification(ctx context.Context, in *GetClarificationRequest, opts ...grpc.CallOption) (*GetClarificationResponse, error)
	// ClarificationListを取得する
	ListClarifications(ctx context.Context, in *ListClarificationsRequest, opts ...grpc.CallOption) (*ListClarificationsResponse, error)
	// Clarification を更新する
	UpdateClarification(ctx context.Context, in *UpdateClarificationRequest, opts ...grpc.CallOption) (*UpdateClarificationResponse, error)
	// Clarification を削除する
	DeleteClarification(ctx context.Context, in *DeleteClarificationRequest, opts ...grpc.CallOption) (*DeleteClarificationResponse, error)
	// Answerを追加する
	CreateAnswer(ctx context.Context, in *CreateAnswerRequest, opts ...grpc.CallOption) (*CreateAnswerResponse, error)
	// Answerを取得する
	GetAnswer(ctx context.Context, in *GetAnswerRequest, opts ...grpc.CallOption) (*GetAnswerResponse, error)
	// Answerを更新する
	UpdateAnswer(ctx context.Context, in *UpdateAnswerRequest, opts ...grpc.CallOption) (*UpdateAnswerResponse, error)
	// Answerを削除する
	DeleteAnswer(ctx context.Context, in *DeleteAnswerRequest, opts ...grpc.CallOption) (*DeleteAnswerResponse, error)
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

func (c *taskServiceClient) CreateClarification(ctx context.Context, in *CreateClarificationRequest, opts ...grpc.CallOption) (*CreateClarificationResponse, error) {
	out := new(CreateClarificationResponse)
	err := c.cc.Invoke(ctx, TaskService_CreateClarification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetClarification(ctx context.Context, in *GetClarificationRequest, opts ...grpc.CallOption) (*GetClarificationResponse, error) {
	out := new(GetClarificationResponse)
	err := c.cc.Invoke(ctx, TaskService_GetClarification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListClarifications(ctx context.Context, in *ListClarificationsRequest, opts ...grpc.CallOption) (*ListClarificationsResponse, error) {
	out := new(ListClarificationsResponse)
	err := c.cc.Invoke(ctx, TaskService_ListClarifications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateClarification(ctx context.Context, in *UpdateClarificationRequest, opts ...grpc.CallOption) (*UpdateClarificationResponse, error) {
	out := new(UpdateClarificationResponse)
	err := c.cc.Invoke(ctx, TaskService_UpdateClarification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteClarification(ctx context.Context, in *DeleteClarificationRequest, opts ...grpc.CallOption) (*DeleteClarificationResponse, error) {
	out := new(DeleteClarificationResponse)
	err := c.cc.Invoke(ctx, TaskService_DeleteClarification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) CreateAnswer(ctx context.Context, in *CreateAnswerRequest, opts ...grpc.CallOption) (*CreateAnswerResponse, error) {
	out := new(CreateAnswerResponse)
	err := c.cc.Invoke(ctx, TaskService_CreateAnswer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetAnswer(ctx context.Context, in *GetAnswerRequest, opts ...grpc.CallOption) (*GetAnswerResponse, error) {
	out := new(GetAnswerResponse)
	err := c.cc.Invoke(ctx, TaskService_GetAnswer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateAnswer(ctx context.Context, in *UpdateAnswerRequest, opts ...grpc.CallOption) (*UpdateAnswerResponse, error) {
	out := new(UpdateAnswerResponse)
	err := c.cc.Invoke(ctx, TaskService_UpdateAnswer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteAnswer(ctx context.Context, in *DeleteAnswerRequest, opts ...grpc.CallOption) (*DeleteAnswerResponse, error) {
	out := new(DeleteAnswerResponse)
	err := c.cc.Invoke(ctx, TaskService_DeleteAnswer_FullMethodName, in, out, opts...)
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
	// Clarification を作成する
	CreateClarification(context.Context, *CreateClarificationRequest) (*CreateClarificationResponse, error)
	// 指定された Clarification を取得する
	GetClarification(context.Context, *GetClarificationRequest) (*GetClarificationResponse, error)
	// ClarificationListを取得する
	ListClarifications(context.Context, *ListClarificationsRequest) (*ListClarificationsResponse, error)
	// Clarification を更新する
	UpdateClarification(context.Context, *UpdateClarificationRequest) (*UpdateClarificationResponse, error)
	// Clarification を削除する
	DeleteClarification(context.Context, *DeleteClarificationRequest) (*DeleteClarificationResponse, error)
	// Answerを追加する
	CreateAnswer(context.Context, *CreateAnswerRequest) (*CreateAnswerResponse, error)
	// Answerを取得する
	GetAnswer(context.Context, *GetAnswerRequest) (*GetAnswerResponse, error)
	// Answerを更新する
	UpdateAnswer(context.Context, *UpdateAnswerRequest) (*UpdateAnswerResponse, error)
	// Answerを削除する
	DeleteAnswer(context.Context, *DeleteAnswerRequest) (*DeleteAnswerResponse, error)
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
func (UnimplementedTaskServiceServer) CreateClarification(context.Context, *CreateClarificationRequest) (*CreateClarificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClarification not implemented")
}
func (UnimplementedTaskServiceServer) GetClarification(context.Context, *GetClarificationRequest) (*GetClarificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClarification not implemented")
}
func (UnimplementedTaskServiceServer) ListClarifications(context.Context, *ListClarificationsRequest) (*ListClarificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClarifications not implemented")
}
func (UnimplementedTaskServiceServer) UpdateClarification(context.Context, *UpdateClarificationRequest) (*UpdateClarificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClarification not implemented")
}
func (UnimplementedTaskServiceServer) DeleteClarification(context.Context, *DeleteClarificationRequest) (*DeleteClarificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClarification not implemented")
}
func (UnimplementedTaskServiceServer) CreateAnswer(context.Context, *CreateAnswerRequest) (*CreateAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnswer not implemented")
}
func (UnimplementedTaskServiceServer) GetAnswer(context.Context, *GetAnswerRequest) (*GetAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswer not implemented")
}
func (UnimplementedTaskServiceServer) UpdateAnswer(context.Context, *UpdateAnswerRequest) (*UpdateAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnswer not implemented")
}
func (UnimplementedTaskServiceServer) DeleteAnswer(context.Context, *DeleteAnswerRequest) (*DeleteAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnswer not implemented")
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

func _TaskService_CreateClarification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClarificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateClarification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_CreateClarification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateClarification(ctx, req.(*CreateClarificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetClarification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClarificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetClarification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetClarification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetClarification(ctx, req.(*GetClarificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListClarifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClarificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListClarifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_ListClarifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListClarifications(ctx, req.(*ListClarificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateClarification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClarificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateClarification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_UpdateClarification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateClarification(ctx, req.(*UpdateClarificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteClarification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClarificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteClarification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_DeleteClarification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteClarification(ctx, req.(*DeleteClarificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_CreateAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_CreateAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateAnswer(ctx, req.(*CreateAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetAnswer(ctx, req.(*GetAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_UpdateAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateAnswer(ctx, req.(*UpdateAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_DeleteAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteAnswer(ctx, req.(*DeleteAnswerRequest))
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
			MethodName: "CreateClarification",
			Handler:    _TaskService_CreateClarification_Handler,
		},
		{
			MethodName: "GetClarification",
			Handler:    _TaskService_GetClarification_Handler,
		},
		{
			MethodName: "ListClarifications",
			Handler:    _TaskService_ListClarifications_Handler,
		},
		{
			MethodName: "UpdateClarification",
			Handler:    _TaskService_UpdateClarification_Handler,
		},
		{
			MethodName: "DeleteClarification",
			Handler:    _TaskService_DeleteClarification_Handler,
		},
		{
			MethodName: "CreateAnswer",
			Handler:    _TaskService_CreateAnswer_Handler,
		},
		{
			MethodName: "GetAnswer",
			Handler:    _TaskService_GetAnswer_Handler,
		},
		{
			MethodName: "UpdateAnswer",
			Handler:    _TaskService_UpdateAnswer_Handler,
		},
		{
			MethodName: "DeleteAnswer",
			Handler:    _TaskService_DeleteAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/v1/services.proto",
}

const (
	HealthcheckService_Ping_FullMethodName = "/backend.v1.HealthcheckService/Ping"
)

// HealthcheckServiceClient is the client API for HealthcheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthcheckServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type healthcheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthcheckServiceClient(cc grpc.ClientConnInterface) HealthcheckServiceClient {
	return &healthcheckServiceClient{cc}
}

func (c *healthcheckServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, HealthcheckService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthcheckServiceServer is the server API for HealthcheckService service.
// All implementations should embed UnimplementedHealthcheckServiceServer
// for forward compatibility
type HealthcheckServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedHealthcheckServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHealthcheckServiceServer struct {
}

func (UnimplementedHealthcheckServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

// UnsafeHealthcheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthcheckServiceServer will
// result in compilation errors.
type UnsafeHealthcheckServiceServer interface {
	mustEmbedUnimplementedHealthcheckServiceServer()
}

func RegisterHealthcheckServiceServer(s grpc.ServiceRegistrar, srv HealthcheckServiceServer) {
	s.RegisterService(&HealthcheckService_ServiceDesc, srv)
}

func _HealthcheckService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthcheckServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthcheckService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthcheckServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthcheckService_ServiceDesc is the grpc.ServiceDesc for HealthcheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthcheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.v1.HealthcheckService",
	HandlerType: (*HealthcheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HealthcheckService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/v1/services.proto",
}
