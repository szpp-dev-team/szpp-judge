// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: backend/v1/contest_service.proto

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
	ContestService_CreateContest_FullMethodName           = "/backend.v1.ContestService/CreateContest"
	ContestService_GetContest_FullMethodName              = "/backend.v1.ContestService/GetContest"
	ContestService_ListContests_FullMethodName            = "/backend.v1.ContestService/ListContests"
	ContestService_ListContestTasks_FullMethodName        = "/backend.v1.ContestService/ListContestTasks"
	ContestService_GetMySubmissionStatuses_FullMethodName = "/backend.v1.ContestService/GetMySubmissionStatuses"
	ContestService_GetStandings_FullMethodName            = "/backend.v1.ContestService/GetStandings"
)

// ContestServiceClient is the client API for ContestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContestServiceClient interface {
	CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*CreateContestResponse, error)
	GetContest(ctx context.Context, in *GetContestRequest, opts ...grpc.CallOption) (*GetContestResponse, error)
	ListContests(ctx context.Context, in *ListContestsRequest, opts ...grpc.CallOption) (*ListContestsResponse, error)
	ListContestTasks(ctx context.Context, in *ListContestTasksRequest, opts ...grpc.CallOption) (*ListContestTasksResponse, error)
	// 自分の問題ごとの結果情報を返す
	GetMySubmissionStatuses(ctx context.Context, in *GetMySubmissionStatusesRequest, opts ...grpc.CallOption) (*GetMySubmissionStatusesResponse, error)
	// 順位表取得
	GetStandings(ctx context.Context, in *GetStandingsRequest, opts ...grpc.CallOption) (*GetStandingsResponse, error)
}

type contestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContestServiceClient(cc grpc.ClientConnInterface) ContestServiceClient {
	return &contestServiceClient{cc}
}

func (c *contestServiceClient) CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*CreateContestResponse, error) {
	out := new(CreateContestResponse)
	err := c.cc.Invoke(ctx, ContestService_CreateContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) GetContest(ctx context.Context, in *GetContestRequest, opts ...grpc.CallOption) (*GetContestResponse, error) {
	out := new(GetContestResponse)
	err := c.cc.Invoke(ctx, ContestService_GetContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) ListContests(ctx context.Context, in *ListContestsRequest, opts ...grpc.CallOption) (*ListContestsResponse, error) {
	out := new(ListContestsResponse)
	err := c.cc.Invoke(ctx, ContestService_ListContests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) ListContestTasks(ctx context.Context, in *ListContestTasksRequest, opts ...grpc.CallOption) (*ListContestTasksResponse, error) {
	out := new(ListContestTasksResponse)
	err := c.cc.Invoke(ctx, ContestService_ListContestTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) GetMySubmissionStatuses(ctx context.Context, in *GetMySubmissionStatusesRequest, opts ...grpc.CallOption) (*GetMySubmissionStatusesResponse, error) {
	out := new(GetMySubmissionStatusesResponse)
	err := c.cc.Invoke(ctx, ContestService_GetMySubmissionStatuses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) GetStandings(ctx context.Context, in *GetStandingsRequest, opts ...grpc.CallOption) (*GetStandingsResponse, error) {
	out := new(GetStandingsResponse)
	err := c.cc.Invoke(ctx, ContestService_GetStandings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContestServiceServer is the server API for ContestService service.
// All implementations should embed UnimplementedContestServiceServer
// for forward compatibility
type ContestServiceServer interface {
	CreateContest(context.Context, *CreateContestRequest) (*CreateContestResponse, error)
	GetContest(context.Context, *GetContestRequest) (*GetContestResponse, error)
	ListContests(context.Context, *ListContestsRequest) (*ListContestsResponse, error)
	ListContestTasks(context.Context, *ListContestTasksRequest) (*ListContestTasksResponse, error)
	// 自分の問題ごとの結果情報を返す
	GetMySubmissionStatuses(context.Context, *GetMySubmissionStatusesRequest) (*GetMySubmissionStatusesResponse, error)
	// 順位表取得
	GetStandings(context.Context, *GetStandingsRequest) (*GetStandingsResponse, error)
}

// UnimplementedContestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedContestServiceServer struct {
}

func (UnimplementedContestServiceServer) CreateContest(context.Context, *CreateContestRequest) (*CreateContestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContest not implemented")
}
func (UnimplementedContestServiceServer) GetContest(context.Context, *GetContestRequest) (*GetContestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContest not implemented")
}
func (UnimplementedContestServiceServer) ListContests(context.Context, *ListContestsRequest) (*ListContestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContests not implemented")
}
func (UnimplementedContestServiceServer) ListContestTasks(context.Context, *ListContestTasksRequest) (*ListContestTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContestTasks not implemented")
}
func (UnimplementedContestServiceServer) GetMySubmissionStatuses(context.Context, *GetMySubmissionStatusesRequest) (*GetMySubmissionStatusesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMySubmissionStatuses not implemented")
}
func (UnimplementedContestServiceServer) GetStandings(context.Context, *GetStandingsRequest) (*GetStandingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStandings not implemented")
}

// UnsafeContestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContestServiceServer will
// result in compilation errors.
type UnsafeContestServiceServer interface {
	mustEmbedUnimplementedContestServiceServer()
}

func RegisterContestServiceServer(s grpc.ServiceRegistrar, srv ContestServiceServer) {
	s.RegisterService(&ContestService_ServiceDesc, srv)
}

func _ContestService_CreateContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).CreateContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_CreateContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).CreateContest(ctx, req.(*CreateContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_GetContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).GetContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_GetContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).GetContest(ctx, req.(*GetContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_ListContests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).ListContests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_ListContests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).ListContests(ctx, req.(*ListContestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_ListContestTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContestTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).ListContestTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_ListContestTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).ListContestTasks(ctx, req.(*ListContestTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_GetMySubmissionStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMySubmissionStatusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).GetMySubmissionStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_GetMySubmissionStatuses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).GetMySubmissionStatuses(ctx, req.(*GetMySubmissionStatusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_GetStandings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStandingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).GetStandings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_GetStandings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).GetStandings(ctx, req.(*GetStandingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContestService_ServiceDesc is the grpc.ServiceDesc for ContestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.v1.ContestService",
	HandlerType: (*ContestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContest",
			Handler:    _ContestService_CreateContest_Handler,
		},
		{
			MethodName: "GetContest",
			Handler:    _ContestService_GetContest_Handler,
		},
		{
			MethodName: "ListContests",
			Handler:    _ContestService_ListContests_Handler,
		},
		{
			MethodName: "ListContestTasks",
			Handler:    _ContestService_ListContestTasks_Handler,
		},
		{
			MethodName: "GetMySubmissionStatuses",
			Handler:    _ContestService_GetMySubmissionStatuses_Handler,
		},
		{
			MethodName: "GetStandings",
			Handler:    _ContestService_GetStandings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/v1/contest_service.proto",
}