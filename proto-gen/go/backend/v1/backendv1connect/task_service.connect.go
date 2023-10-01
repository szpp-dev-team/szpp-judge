// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: backend/v1/task_service.proto

package backendv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// TaskServiceName is the fully-qualified name of the TaskService service.
	TaskServiceName = "backend.v1.TaskService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TaskServiceCreateTaskProcedure is the fully-qualified name of the TaskService's CreateTask RPC.
	TaskServiceCreateTaskProcedure = "/backend.v1.TaskService/CreateTask"
	// TaskServiceGetTaskProcedure is the fully-qualified name of the TaskService's GetTask RPC.
	TaskServiceGetTaskProcedure = "/backend.v1.TaskService/GetTask"
	// TaskServiceUpdateTaskProcedure is the fully-qualified name of the TaskService's UpdateTask RPC.
	TaskServiceUpdateTaskProcedure = "/backend.v1.TaskService/UpdateTask"
	// TaskServiceGetTestcaseSetsProcedure is the fully-qualified name of the TaskService's
	// GetTestcaseSets RPC.
	TaskServiceGetTestcaseSetsProcedure = "/backend.v1.TaskService/GetTestcaseSets"
	// TaskServiceSyncTestcaseSetsProcedure is the fully-qualified name of the TaskService's
	// SyncTestcaseSets RPC.
	TaskServiceSyncTestcaseSetsProcedure = "/backend.v1.TaskService/SyncTestcaseSets"
)

// TaskServiceClient is a client for the backend.v1.TaskService service.
type TaskServiceClient interface {
	// Task を作成する
	CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.CreateTaskResponse], error)
	// 指定された Task を取得する
	GetTask(context.Context, *connect.Request[v1.GetTaskRequest]) (*connect.Response[v1.GetTaskResponse], error)
	// Task を更新する
	UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.UpdateTaskResponse], error)
	// TestcaseSet の一覧を取得する。また、Testcase の一覧も取得する。
	// contestant によるリクエストの場合は sample のみ取得する。
	GetTestcaseSets(context.Context, *connect.Request[v1.GetTestcaseSetsRequest]) (*connect.Response[v1.GetTestcaseSetsResponse], error)
	// TestcaseSet を同期する。全てのリソースは上書きされ、このリクエストに含まれないリソースは削除される。
	SyncTestcaseSets(context.Context, *connect.Request[v1.SyncTestcaseSetsRequest]) (*connect.Response[v1.SyncTestcaseSetsResponse], error)
}

// NewTaskServiceClient constructs a client for the backend.v1.TaskService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTaskServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TaskServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &taskServiceClient{
		createTask: connect.NewClient[v1.CreateTaskRequest, v1.CreateTaskResponse](
			httpClient,
			baseURL+TaskServiceCreateTaskProcedure,
			opts...,
		),
		getTask: connect.NewClient[v1.GetTaskRequest, v1.GetTaskResponse](
			httpClient,
			baseURL+TaskServiceGetTaskProcedure,
			opts...,
		),
		updateTask: connect.NewClient[v1.UpdateTaskRequest, v1.UpdateTaskResponse](
			httpClient,
			baseURL+TaskServiceUpdateTaskProcedure,
			opts...,
		),
		getTestcaseSets: connect.NewClient[v1.GetTestcaseSetsRequest, v1.GetTestcaseSetsResponse](
			httpClient,
			baseURL+TaskServiceGetTestcaseSetsProcedure,
			opts...,
		),
		syncTestcaseSets: connect.NewClient[v1.SyncTestcaseSetsRequest, v1.SyncTestcaseSetsResponse](
			httpClient,
			baseURL+TaskServiceSyncTestcaseSetsProcedure,
			opts...,
		),
	}
}

// taskServiceClient implements TaskServiceClient.
type taskServiceClient struct {
	createTask       *connect.Client[v1.CreateTaskRequest, v1.CreateTaskResponse]
	getTask          *connect.Client[v1.GetTaskRequest, v1.GetTaskResponse]
	updateTask       *connect.Client[v1.UpdateTaskRequest, v1.UpdateTaskResponse]
	getTestcaseSets  *connect.Client[v1.GetTestcaseSetsRequest, v1.GetTestcaseSetsResponse]
	syncTestcaseSets *connect.Client[v1.SyncTestcaseSetsRequest, v1.SyncTestcaseSetsResponse]
}

// CreateTask calls backend.v1.TaskService.CreateTask.
func (c *taskServiceClient) CreateTask(ctx context.Context, req *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.CreateTaskResponse], error) {
	return c.createTask.CallUnary(ctx, req)
}

// GetTask calls backend.v1.TaskService.GetTask.
func (c *taskServiceClient) GetTask(ctx context.Context, req *connect.Request[v1.GetTaskRequest]) (*connect.Response[v1.GetTaskResponse], error) {
	return c.getTask.CallUnary(ctx, req)
}

// UpdateTask calls backend.v1.TaskService.UpdateTask.
func (c *taskServiceClient) UpdateTask(ctx context.Context, req *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.UpdateTaskResponse], error) {
	return c.updateTask.CallUnary(ctx, req)
}

// GetTestcaseSets calls backend.v1.TaskService.GetTestcaseSets.
func (c *taskServiceClient) GetTestcaseSets(ctx context.Context, req *connect.Request[v1.GetTestcaseSetsRequest]) (*connect.Response[v1.GetTestcaseSetsResponse], error) {
	return c.getTestcaseSets.CallUnary(ctx, req)
}

// SyncTestcaseSets calls backend.v1.TaskService.SyncTestcaseSets.
func (c *taskServiceClient) SyncTestcaseSets(ctx context.Context, req *connect.Request[v1.SyncTestcaseSetsRequest]) (*connect.Response[v1.SyncTestcaseSetsResponse], error) {
	return c.syncTestcaseSets.CallUnary(ctx, req)
}

// TaskServiceHandler is an implementation of the backend.v1.TaskService service.
type TaskServiceHandler interface {
	// Task を作成する
	CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.CreateTaskResponse], error)
	// 指定された Task を取得する
	GetTask(context.Context, *connect.Request[v1.GetTaskRequest]) (*connect.Response[v1.GetTaskResponse], error)
	// Task を更新する
	UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.UpdateTaskResponse], error)
	// TestcaseSet の一覧を取得する。また、Testcase の一覧も取得する。
	// contestant によるリクエストの場合は sample のみ取得する。
	GetTestcaseSets(context.Context, *connect.Request[v1.GetTestcaseSetsRequest]) (*connect.Response[v1.GetTestcaseSetsResponse], error)
	// TestcaseSet を同期する。全てのリソースは上書きされ、このリクエストに含まれないリソースは削除される。
	SyncTestcaseSets(context.Context, *connect.Request[v1.SyncTestcaseSetsRequest]) (*connect.Response[v1.SyncTestcaseSetsResponse], error)
}

// NewTaskServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTaskServiceHandler(svc TaskServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	taskServiceCreateTaskHandler := connect.NewUnaryHandler(
		TaskServiceCreateTaskProcedure,
		svc.CreateTask,
		opts...,
	)
	taskServiceGetTaskHandler := connect.NewUnaryHandler(
		TaskServiceGetTaskProcedure,
		svc.GetTask,
		opts...,
	)
	taskServiceUpdateTaskHandler := connect.NewUnaryHandler(
		TaskServiceUpdateTaskProcedure,
		svc.UpdateTask,
		opts...,
	)
	taskServiceGetTestcaseSetsHandler := connect.NewUnaryHandler(
		TaskServiceGetTestcaseSetsProcedure,
		svc.GetTestcaseSets,
		opts...,
	)
	taskServiceSyncTestcaseSetsHandler := connect.NewUnaryHandler(
		TaskServiceSyncTestcaseSetsProcedure,
		svc.SyncTestcaseSets,
		opts...,
	)
	return "/backend.v1.TaskService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TaskServiceCreateTaskProcedure:
			taskServiceCreateTaskHandler.ServeHTTP(w, r)
		case TaskServiceGetTaskProcedure:
			taskServiceGetTaskHandler.ServeHTTP(w, r)
		case TaskServiceUpdateTaskProcedure:
			taskServiceUpdateTaskHandler.ServeHTTP(w, r)
		case TaskServiceGetTestcaseSetsProcedure:
			taskServiceGetTestcaseSetsHandler.ServeHTTP(w, r)
		case TaskServiceSyncTestcaseSetsProcedure:
			taskServiceSyncTestcaseSetsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTaskServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTaskServiceHandler struct{}

func (UnimplementedTaskServiceHandler) CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.CreateTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.TaskService.CreateTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) GetTask(context.Context, *connect.Request[v1.GetTaskRequest]) (*connect.Response[v1.GetTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.TaskService.GetTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) UpdateTask(context.Context, *connect.Request[v1.UpdateTaskRequest]) (*connect.Response[v1.UpdateTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.TaskService.UpdateTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) GetTestcaseSets(context.Context, *connect.Request[v1.GetTestcaseSetsRequest]) (*connect.Response[v1.GetTestcaseSetsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.TaskService.GetTestcaseSets is not implemented"))
}

func (UnimplementedTaskServiceHandler) SyncTestcaseSets(context.Context, *connect.Request[v1.SyncTestcaseSetsRequest]) (*connect.Response[v1.SyncTestcaseSetsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.TaskService.SyncTestcaseSets is not implemented"))
}