// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: backend/v1/contest_service.proto

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
	// ContestServiceName is the fully-qualified name of the ContestService service.
	ContestServiceName = "backend.v1.ContestService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ContestServiceCreateContestProcedure is the fully-qualified name of the ContestService's
	// CreateContest RPC.
	ContestServiceCreateContestProcedure = "/backend.v1.ContestService/CreateContest"
	// ContestServiceGetContestProcedure is the fully-qualified name of the ContestService's GetContest
	// RPC.
	ContestServiceGetContestProcedure = "/backend.v1.ContestService/GetContest"
	// ContestServiceUpdateContestProcedure is the fully-qualified name of the ContestService's
	// UpdateContest RPC.
	ContestServiceUpdateContestProcedure = "/backend.v1.ContestService/UpdateContest"
	// ContestServiceListContestsProcedure is the fully-qualified name of the ContestService's
	// ListContests RPC.
	ContestServiceListContestsProcedure = "/backend.v1.ContestService/ListContests"
	// ContestServiceListContestTasksProcedure is the fully-qualified name of the ContestService's
	// ListContestTasks RPC.
	ContestServiceListContestTasksProcedure = "/backend.v1.ContestService/ListContestTasks"
	// ContestServiceGetContestTaskProcedure is the fully-qualified name of the ContestService's
	// GetContestTask RPC.
	ContestServiceGetContestTaskProcedure = "/backend.v1.ContestService/GetContestTask"
	// ContestServiceSyncContestTasksProcedure is the fully-qualified name of the ContestService's
	// SyncContestTasks RPC.
	ContestServiceSyncContestTasksProcedure = "/backend.v1.ContestService/SyncContestTasks"
	// ContestServiceGetMySubmissionStatusesProcedure is the fully-qualified name of the
	// ContestService's GetMySubmissionStatuses RPC.
	ContestServiceGetMySubmissionStatusesProcedure = "/backend.v1.ContestService/GetMySubmissionStatuses"
	// ContestServiceGetStandingsProcedure is the fully-qualified name of the ContestService's
	// GetStandings RPC.
	ContestServiceGetStandingsProcedure = "/backend.v1.ContestService/GetStandings"
	// ContestServiceRegisterMeProcedure is the fully-qualified name of the ContestService's RegisterMe
	// RPC.
	ContestServiceRegisterMeProcedure = "/backend.v1.ContestService/RegisterMe"
	// ContestServiceCreateClarificationProcedure is the fully-qualified name of the ContestService's
	// CreateClarification RPC.
	ContestServiceCreateClarificationProcedure = "/backend.v1.ContestService/CreateClarification"
	// ContestServiceListClarificationsProcedure is the fully-qualified name of the ContestService's
	// ListClarifications RPC.
	ContestServiceListClarificationsProcedure = "/backend.v1.ContestService/ListClarifications"
	// ContestServiceDeleteClarificationProcedure is the fully-qualified name of the ContestService's
	// DeleteClarification RPC.
	ContestServiceDeleteClarificationProcedure = "/backend.v1.ContestService/DeleteClarification"
	// ContestServiceCreateAnswerProcedure is the fully-qualified name of the ContestService's
	// CreateAnswer RPC.
	ContestServiceCreateAnswerProcedure = "/backend.v1.ContestService/CreateAnswer"
	// ContestServiceUpdateAnswerProcedure is the fully-qualified name of the ContestService's
	// UpdateAnswer RPC.
	ContestServiceUpdateAnswerProcedure = "/backend.v1.ContestService/UpdateAnswer"
	// ContestServiceDeleteAnswerProcedure is the fully-qualified name of the ContestService's
	// DeleteAnswer RPC.
	ContestServiceDeleteAnswerProcedure = "/backend.v1.ContestService/DeleteAnswer"
)

// ContestServiceClient is a client for the backend.v1.ContestService service.
type ContestServiceClient interface {
	// コンテストを作成する
	CreateContest(context.Context, *connect.Request[v1.CreateContestRequest]) (*connect.Response[v1.CreateContestResponse], error)
	// コンテスト情報を取得する
	GetContest(context.Context, *connect.Request[v1.GetContestRequest]) (*connect.Response[v1.GetContestResponse], error)
	// コンテスト情報を更新する
	UpdateContest(context.Context, *connect.Request[v1.UpdateContestRequest]) (*connect.Response[v1.UpdateContestResponse], error)
	// コンテストの一覧を取得する
	ListContests(context.Context, *connect.Request[v1.ListContestsRequest]) (*connect.Response[v1.ListContestsResponse], error)
	// コンテストに紐づく問題の一覧を取得する
	ListContestTasks(context.Context, *connect.Request[v1.ListContestTasksRequest]) (*connect.Response[v1.ListContestTasksResponse], error)
	// コンテストに紐づく問題を取得する
	GetContestTask(context.Context, *connect.Request[v1.GetContestTaskRequest]) (*connect.Response[v1.GetContestTaskResponse], error)
	// 問題をコンテストに紐づかせる
	SyncContestTasks(context.Context, *connect.Request[v1.SyncContestTasksRequest]) (*connect.Response[v1.SyncContestTasksResponse], error)
	// 自分の問題ごとの結果情報を取得する
	GetMySubmissionStatuses(context.Context, *connect.Request[v1.GetMySubmissionStatusesRequest]) (*connect.Response[v1.GetMySubmissionStatusesResponse], error)
	// 順位表を取得する
	GetStandings(context.Context, *connect.Request[v1.GetStandingsRequest]) (*connect.Response[v1.GetStandingsResponse], error)
	// 参加登録をする
	RegisterMe(context.Context, *connect.Request[v1.RegisterMeRequest]) (*connect.Response[v1.RegisterMeResponse], error)
	// Clarification を作成する
	CreateClarification(context.Context, *connect.Request[v1.CreateClarificationRequest]) (*connect.Response[v1.CreateClarificationResponse], error)
	// ClarificationListを取得する
	ListClarifications(context.Context, *connect.Request[v1.ListClarificationsRequest]) (*connect.Response[v1.ListClarificationsResponse], error)
	// Clarification を削除する
	DeleteClarification(context.Context, *connect.Request[v1.DeleteClarificationRequest]) (*connect.Response[v1.DeleteClarificationResponse], error)
	// Answerを追加する
	CreateAnswer(context.Context, *connect.Request[v1.CreateAnswerRequest]) (*connect.Response[v1.CreateAnswerResponse], error)
	// Answerを更新する
	UpdateAnswer(context.Context, *connect.Request[v1.UpdateAnswerRequest]) (*connect.Response[v1.UpdateAnswerResponse], error)
	// Answerを削除する
	DeleteAnswer(context.Context, *connect.Request[v1.DeleteAnswerRequest]) (*connect.Response[v1.DeleteAnswerResponse], error)
}

// NewContestServiceClient constructs a client for the backend.v1.ContestService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewContestServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ContestServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &contestServiceClient{
		createContest: connect.NewClient[v1.CreateContestRequest, v1.CreateContestResponse](
			httpClient,
			baseURL+ContestServiceCreateContestProcedure,
			opts...,
		),
		getContest: connect.NewClient[v1.GetContestRequest, v1.GetContestResponse](
			httpClient,
			baseURL+ContestServiceGetContestProcedure,
			opts...,
		),
		updateContest: connect.NewClient[v1.UpdateContestRequest, v1.UpdateContestResponse](
			httpClient,
			baseURL+ContestServiceUpdateContestProcedure,
			opts...,
		),
		listContests: connect.NewClient[v1.ListContestsRequest, v1.ListContestsResponse](
			httpClient,
			baseURL+ContestServiceListContestsProcedure,
			opts...,
		),
		listContestTasks: connect.NewClient[v1.ListContestTasksRequest, v1.ListContestTasksResponse](
			httpClient,
			baseURL+ContestServiceListContestTasksProcedure,
			opts...,
		),
		getContestTask: connect.NewClient[v1.GetContestTaskRequest, v1.GetContestTaskResponse](
			httpClient,
			baseURL+ContestServiceGetContestTaskProcedure,
			opts...,
		),
		syncContestTasks: connect.NewClient[v1.SyncContestTasksRequest, v1.SyncContestTasksResponse](
			httpClient,
			baseURL+ContestServiceSyncContestTasksProcedure,
			opts...,
		),
		getMySubmissionStatuses: connect.NewClient[v1.GetMySubmissionStatusesRequest, v1.GetMySubmissionStatusesResponse](
			httpClient,
			baseURL+ContestServiceGetMySubmissionStatusesProcedure,
			opts...,
		),
		getStandings: connect.NewClient[v1.GetStandingsRequest, v1.GetStandingsResponse](
			httpClient,
			baseURL+ContestServiceGetStandingsProcedure,
			opts...,
		),
		registerMe: connect.NewClient[v1.RegisterMeRequest, v1.RegisterMeResponse](
			httpClient,
			baseURL+ContestServiceRegisterMeProcedure,
			opts...,
		),
		createClarification: connect.NewClient[v1.CreateClarificationRequest, v1.CreateClarificationResponse](
			httpClient,
			baseURL+ContestServiceCreateClarificationProcedure,
			opts...,
		),
		listClarifications: connect.NewClient[v1.ListClarificationsRequest, v1.ListClarificationsResponse](
			httpClient,
			baseURL+ContestServiceListClarificationsProcedure,
			opts...,
		),
		deleteClarification: connect.NewClient[v1.DeleteClarificationRequest, v1.DeleteClarificationResponse](
			httpClient,
			baseURL+ContestServiceDeleteClarificationProcedure,
			opts...,
		),
		createAnswer: connect.NewClient[v1.CreateAnswerRequest, v1.CreateAnswerResponse](
			httpClient,
			baseURL+ContestServiceCreateAnswerProcedure,
			opts...,
		),
		updateAnswer: connect.NewClient[v1.UpdateAnswerRequest, v1.UpdateAnswerResponse](
			httpClient,
			baseURL+ContestServiceUpdateAnswerProcedure,
			opts...,
		),
		deleteAnswer: connect.NewClient[v1.DeleteAnswerRequest, v1.DeleteAnswerResponse](
			httpClient,
			baseURL+ContestServiceDeleteAnswerProcedure,
			opts...,
		),
	}
}

// contestServiceClient implements ContestServiceClient.
type contestServiceClient struct {
	createContest           *connect.Client[v1.CreateContestRequest, v1.CreateContestResponse]
	getContest              *connect.Client[v1.GetContestRequest, v1.GetContestResponse]
	updateContest           *connect.Client[v1.UpdateContestRequest, v1.UpdateContestResponse]
	listContests            *connect.Client[v1.ListContestsRequest, v1.ListContestsResponse]
	listContestTasks        *connect.Client[v1.ListContestTasksRequest, v1.ListContestTasksResponse]
	getContestTask          *connect.Client[v1.GetContestTaskRequest, v1.GetContestTaskResponse]
	syncContestTasks        *connect.Client[v1.SyncContestTasksRequest, v1.SyncContestTasksResponse]
	getMySubmissionStatuses *connect.Client[v1.GetMySubmissionStatusesRequest, v1.GetMySubmissionStatusesResponse]
	getStandings            *connect.Client[v1.GetStandingsRequest, v1.GetStandingsResponse]
	registerMe              *connect.Client[v1.RegisterMeRequest, v1.RegisterMeResponse]
	createClarification     *connect.Client[v1.CreateClarificationRequest, v1.CreateClarificationResponse]
	listClarifications      *connect.Client[v1.ListClarificationsRequest, v1.ListClarificationsResponse]
	deleteClarification     *connect.Client[v1.DeleteClarificationRequest, v1.DeleteClarificationResponse]
	createAnswer            *connect.Client[v1.CreateAnswerRequest, v1.CreateAnswerResponse]
	updateAnswer            *connect.Client[v1.UpdateAnswerRequest, v1.UpdateAnswerResponse]
	deleteAnswer            *connect.Client[v1.DeleteAnswerRequest, v1.DeleteAnswerResponse]
}

// CreateContest calls backend.v1.ContestService.CreateContest.
func (c *contestServiceClient) CreateContest(ctx context.Context, req *connect.Request[v1.CreateContestRequest]) (*connect.Response[v1.CreateContestResponse], error) {
	return c.createContest.CallUnary(ctx, req)
}

// GetContest calls backend.v1.ContestService.GetContest.
func (c *contestServiceClient) GetContest(ctx context.Context, req *connect.Request[v1.GetContestRequest]) (*connect.Response[v1.GetContestResponse], error) {
	return c.getContest.CallUnary(ctx, req)
}

// UpdateContest calls backend.v1.ContestService.UpdateContest.
func (c *contestServiceClient) UpdateContest(ctx context.Context, req *connect.Request[v1.UpdateContestRequest]) (*connect.Response[v1.UpdateContestResponse], error) {
	return c.updateContest.CallUnary(ctx, req)
}

// ListContests calls backend.v1.ContestService.ListContests.
func (c *contestServiceClient) ListContests(ctx context.Context, req *connect.Request[v1.ListContestsRequest]) (*connect.Response[v1.ListContestsResponse], error) {
	return c.listContests.CallUnary(ctx, req)
}

// ListContestTasks calls backend.v1.ContestService.ListContestTasks.
func (c *contestServiceClient) ListContestTasks(ctx context.Context, req *connect.Request[v1.ListContestTasksRequest]) (*connect.Response[v1.ListContestTasksResponse], error) {
	return c.listContestTasks.CallUnary(ctx, req)
}

// GetContestTask calls backend.v1.ContestService.GetContestTask.
func (c *contestServiceClient) GetContestTask(ctx context.Context, req *connect.Request[v1.GetContestTaskRequest]) (*connect.Response[v1.GetContestTaskResponse], error) {
	return c.getContestTask.CallUnary(ctx, req)
}

// SyncContestTasks calls backend.v1.ContestService.SyncContestTasks.
func (c *contestServiceClient) SyncContestTasks(ctx context.Context, req *connect.Request[v1.SyncContestTasksRequest]) (*connect.Response[v1.SyncContestTasksResponse], error) {
	return c.syncContestTasks.CallUnary(ctx, req)
}

// GetMySubmissionStatuses calls backend.v1.ContestService.GetMySubmissionStatuses.
func (c *contestServiceClient) GetMySubmissionStatuses(ctx context.Context, req *connect.Request[v1.GetMySubmissionStatusesRequest]) (*connect.Response[v1.GetMySubmissionStatusesResponse], error) {
	return c.getMySubmissionStatuses.CallUnary(ctx, req)
}

// GetStandings calls backend.v1.ContestService.GetStandings.
func (c *contestServiceClient) GetStandings(ctx context.Context, req *connect.Request[v1.GetStandingsRequest]) (*connect.Response[v1.GetStandingsResponse], error) {
	return c.getStandings.CallUnary(ctx, req)
}

// RegisterMe calls backend.v1.ContestService.RegisterMe.
func (c *contestServiceClient) RegisterMe(ctx context.Context, req *connect.Request[v1.RegisterMeRequest]) (*connect.Response[v1.RegisterMeResponse], error) {
	return c.registerMe.CallUnary(ctx, req)
}

// CreateClarification calls backend.v1.ContestService.CreateClarification.
func (c *contestServiceClient) CreateClarification(ctx context.Context, req *connect.Request[v1.CreateClarificationRequest]) (*connect.Response[v1.CreateClarificationResponse], error) {
	return c.createClarification.CallUnary(ctx, req)
}

// ListClarifications calls backend.v1.ContestService.ListClarifications.
func (c *contestServiceClient) ListClarifications(ctx context.Context, req *connect.Request[v1.ListClarificationsRequest]) (*connect.Response[v1.ListClarificationsResponse], error) {
	return c.listClarifications.CallUnary(ctx, req)
}

// DeleteClarification calls backend.v1.ContestService.DeleteClarification.
func (c *contestServiceClient) DeleteClarification(ctx context.Context, req *connect.Request[v1.DeleteClarificationRequest]) (*connect.Response[v1.DeleteClarificationResponse], error) {
	return c.deleteClarification.CallUnary(ctx, req)
}

// CreateAnswer calls backend.v1.ContestService.CreateAnswer.
func (c *contestServiceClient) CreateAnswer(ctx context.Context, req *connect.Request[v1.CreateAnswerRequest]) (*connect.Response[v1.CreateAnswerResponse], error) {
	return c.createAnswer.CallUnary(ctx, req)
}

// UpdateAnswer calls backend.v1.ContestService.UpdateAnswer.
func (c *contestServiceClient) UpdateAnswer(ctx context.Context, req *connect.Request[v1.UpdateAnswerRequest]) (*connect.Response[v1.UpdateAnswerResponse], error) {
	return c.updateAnswer.CallUnary(ctx, req)
}

// DeleteAnswer calls backend.v1.ContestService.DeleteAnswer.
func (c *contestServiceClient) DeleteAnswer(ctx context.Context, req *connect.Request[v1.DeleteAnswerRequest]) (*connect.Response[v1.DeleteAnswerResponse], error) {
	return c.deleteAnswer.CallUnary(ctx, req)
}

// ContestServiceHandler is an implementation of the backend.v1.ContestService service.
type ContestServiceHandler interface {
	// コンテストを作成する
	CreateContest(context.Context, *connect.Request[v1.CreateContestRequest]) (*connect.Response[v1.CreateContestResponse], error)
	// コンテスト情報を取得する
	GetContest(context.Context, *connect.Request[v1.GetContestRequest]) (*connect.Response[v1.GetContestResponse], error)
	// コンテスト情報を更新する
	UpdateContest(context.Context, *connect.Request[v1.UpdateContestRequest]) (*connect.Response[v1.UpdateContestResponse], error)
	// コンテストの一覧を取得する
	ListContests(context.Context, *connect.Request[v1.ListContestsRequest]) (*connect.Response[v1.ListContestsResponse], error)
	// コンテストに紐づく問題の一覧を取得する
	ListContestTasks(context.Context, *connect.Request[v1.ListContestTasksRequest]) (*connect.Response[v1.ListContestTasksResponse], error)
	// コンテストに紐づく問題を取得する
	GetContestTask(context.Context, *connect.Request[v1.GetContestTaskRequest]) (*connect.Response[v1.GetContestTaskResponse], error)
	// 問題をコンテストに紐づかせる
	SyncContestTasks(context.Context, *connect.Request[v1.SyncContestTasksRequest]) (*connect.Response[v1.SyncContestTasksResponse], error)
	// 自分の問題ごとの結果情報を取得する
	GetMySubmissionStatuses(context.Context, *connect.Request[v1.GetMySubmissionStatusesRequest]) (*connect.Response[v1.GetMySubmissionStatusesResponse], error)
	// 順位表を取得する
	GetStandings(context.Context, *connect.Request[v1.GetStandingsRequest]) (*connect.Response[v1.GetStandingsResponse], error)
	// 参加登録をする
	RegisterMe(context.Context, *connect.Request[v1.RegisterMeRequest]) (*connect.Response[v1.RegisterMeResponse], error)
	// Clarification を作成する
	CreateClarification(context.Context, *connect.Request[v1.CreateClarificationRequest]) (*connect.Response[v1.CreateClarificationResponse], error)
	// ClarificationListを取得する
	ListClarifications(context.Context, *connect.Request[v1.ListClarificationsRequest]) (*connect.Response[v1.ListClarificationsResponse], error)
	// Clarification を削除する
	DeleteClarification(context.Context, *connect.Request[v1.DeleteClarificationRequest]) (*connect.Response[v1.DeleteClarificationResponse], error)
	// Answerを追加する
	CreateAnswer(context.Context, *connect.Request[v1.CreateAnswerRequest]) (*connect.Response[v1.CreateAnswerResponse], error)
	// Answerを更新する
	UpdateAnswer(context.Context, *connect.Request[v1.UpdateAnswerRequest]) (*connect.Response[v1.UpdateAnswerResponse], error)
	// Answerを削除する
	DeleteAnswer(context.Context, *connect.Request[v1.DeleteAnswerRequest]) (*connect.Response[v1.DeleteAnswerResponse], error)
}

// NewContestServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewContestServiceHandler(svc ContestServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	contestServiceCreateContestHandler := connect.NewUnaryHandler(
		ContestServiceCreateContestProcedure,
		svc.CreateContest,
		opts...,
	)
	contestServiceGetContestHandler := connect.NewUnaryHandler(
		ContestServiceGetContestProcedure,
		svc.GetContest,
		opts...,
	)
	contestServiceUpdateContestHandler := connect.NewUnaryHandler(
		ContestServiceUpdateContestProcedure,
		svc.UpdateContest,
		opts...,
	)
	contestServiceListContestsHandler := connect.NewUnaryHandler(
		ContestServiceListContestsProcedure,
		svc.ListContests,
		opts...,
	)
	contestServiceListContestTasksHandler := connect.NewUnaryHandler(
		ContestServiceListContestTasksProcedure,
		svc.ListContestTasks,
		opts...,
	)
	contestServiceGetContestTaskHandler := connect.NewUnaryHandler(
		ContestServiceGetContestTaskProcedure,
		svc.GetContestTask,
		opts...,
	)
	contestServiceSyncContestTasksHandler := connect.NewUnaryHandler(
		ContestServiceSyncContestTasksProcedure,
		svc.SyncContestTasks,
		opts...,
	)
	contestServiceGetMySubmissionStatusesHandler := connect.NewUnaryHandler(
		ContestServiceGetMySubmissionStatusesProcedure,
		svc.GetMySubmissionStatuses,
		opts...,
	)
	contestServiceGetStandingsHandler := connect.NewUnaryHandler(
		ContestServiceGetStandingsProcedure,
		svc.GetStandings,
		opts...,
	)
	contestServiceRegisterMeHandler := connect.NewUnaryHandler(
		ContestServiceRegisterMeProcedure,
		svc.RegisterMe,
		opts...,
	)
	contestServiceCreateClarificationHandler := connect.NewUnaryHandler(
		ContestServiceCreateClarificationProcedure,
		svc.CreateClarification,
		opts...,
	)
	contestServiceListClarificationsHandler := connect.NewUnaryHandler(
		ContestServiceListClarificationsProcedure,
		svc.ListClarifications,
		opts...,
	)
	contestServiceDeleteClarificationHandler := connect.NewUnaryHandler(
		ContestServiceDeleteClarificationProcedure,
		svc.DeleteClarification,
		opts...,
	)
	contestServiceCreateAnswerHandler := connect.NewUnaryHandler(
		ContestServiceCreateAnswerProcedure,
		svc.CreateAnswer,
		opts...,
	)
	contestServiceUpdateAnswerHandler := connect.NewUnaryHandler(
		ContestServiceUpdateAnswerProcedure,
		svc.UpdateAnswer,
		opts...,
	)
	contestServiceDeleteAnswerHandler := connect.NewUnaryHandler(
		ContestServiceDeleteAnswerProcedure,
		svc.DeleteAnswer,
		opts...,
	)
	return "/backend.v1.ContestService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ContestServiceCreateContestProcedure:
			contestServiceCreateContestHandler.ServeHTTP(w, r)
		case ContestServiceGetContestProcedure:
			contestServiceGetContestHandler.ServeHTTP(w, r)
		case ContestServiceUpdateContestProcedure:
			contestServiceUpdateContestHandler.ServeHTTP(w, r)
		case ContestServiceListContestsProcedure:
			contestServiceListContestsHandler.ServeHTTP(w, r)
		case ContestServiceListContestTasksProcedure:
			contestServiceListContestTasksHandler.ServeHTTP(w, r)
		case ContestServiceGetContestTaskProcedure:
			contestServiceGetContestTaskHandler.ServeHTTP(w, r)
		case ContestServiceSyncContestTasksProcedure:
			contestServiceSyncContestTasksHandler.ServeHTTP(w, r)
		case ContestServiceGetMySubmissionStatusesProcedure:
			contestServiceGetMySubmissionStatusesHandler.ServeHTTP(w, r)
		case ContestServiceGetStandingsProcedure:
			contestServiceGetStandingsHandler.ServeHTTP(w, r)
		case ContestServiceRegisterMeProcedure:
			contestServiceRegisterMeHandler.ServeHTTP(w, r)
		case ContestServiceCreateClarificationProcedure:
			contestServiceCreateClarificationHandler.ServeHTTP(w, r)
		case ContestServiceListClarificationsProcedure:
			contestServiceListClarificationsHandler.ServeHTTP(w, r)
		case ContestServiceDeleteClarificationProcedure:
			contestServiceDeleteClarificationHandler.ServeHTTP(w, r)
		case ContestServiceCreateAnswerProcedure:
			contestServiceCreateAnswerHandler.ServeHTTP(w, r)
		case ContestServiceUpdateAnswerProcedure:
			contestServiceUpdateAnswerHandler.ServeHTTP(w, r)
		case ContestServiceDeleteAnswerProcedure:
			contestServiceDeleteAnswerHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedContestServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedContestServiceHandler struct{}

func (UnimplementedContestServiceHandler) CreateContest(context.Context, *connect.Request[v1.CreateContestRequest]) (*connect.Response[v1.CreateContestResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.CreateContest is not implemented"))
}

func (UnimplementedContestServiceHandler) GetContest(context.Context, *connect.Request[v1.GetContestRequest]) (*connect.Response[v1.GetContestResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.GetContest is not implemented"))
}

func (UnimplementedContestServiceHandler) UpdateContest(context.Context, *connect.Request[v1.UpdateContestRequest]) (*connect.Response[v1.UpdateContestResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.UpdateContest is not implemented"))
}

func (UnimplementedContestServiceHandler) ListContests(context.Context, *connect.Request[v1.ListContestsRequest]) (*connect.Response[v1.ListContestsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.ListContests is not implemented"))
}

func (UnimplementedContestServiceHandler) ListContestTasks(context.Context, *connect.Request[v1.ListContestTasksRequest]) (*connect.Response[v1.ListContestTasksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.ListContestTasks is not implemented"))
}

func (UnimplementedContestServiceHandler) GetContestTask(context.Context, *connect.Request[v1.GetContestTaskRequest]) (*connect.Response[v1.GetContestTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.GetContestTask is not implemented"))
}

func (UnimplementedContestServiceHandler) SyncContestTasks(context.Context, *connect.Request[v1.SyncContestTasksRequest]) (*connect.Response[v1.SyncContestTasksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.SyncContestTasks is not implemented"))
}

func (UnimplementedContestServiceHandler) GetMySubmissionStatuses(context.Context, *connect.Request[v1.GetMySubmissionStatusesRequest]) (*connect.Response[v1.GetMySubmissionStatusesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.GetMySubmissionStatuses is not implemented"))
}

func (UnimplementedContestServiceHandler) GetStandings(context.Context, *connect.Request[v1.GetStandingsRequest]) (*connect.Response[v1.GetStandingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.GetStandings is not implemented"))
}

func (UnimplementedContestServiceHandler) RegisterMe(context.Context, *connect.Request[v1.RegisterMeRequest]) (*connect.Response[v1.RegisterMeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.RegisterMe is not implemented"))
}

func (UnimplementedContestServiceHandler) CreateClarification(context.Context, *connect.Request[v1.CreateClarificationRequest]) (*connect.Response[v1.CreateClarificationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.CreateClarification is not implemented"))
}

func (UnimplementedContestServiceHandler) ListClarifications(context.Context, *connect.Request[v1.ListClarificationsRequest]) (*connect.Response[v1.ListClarificationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.ListClarifications is not implemented"))
}

func (UnimplementedContestServiceHandler) DeleteClarification(context.Context, *connect.Request[v1.DeleteClarificationRequest]) (*connect.Response[v1.DeleteClarificationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.DeleteClarification is not implemented"))
}

func (UnimplementedContestServiceHandler) CreateAnswer(context.Context, *connect.Request[v1.CreateAnswerRequest]) (*connect.Response[v1.CreateAnswerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.CreateAnswer is not implemented"))
}

func (UnimplementedContestServiceHandler) UpdateAnswer(context.Context, *connect.Request[v1.UpdateAnswerRequest]) (*connect.Response[v1.UpdateAnswerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.UpdateAnswer is not implemented"))
}

func (UnimplementedContestServiceHandler) DeleteAnswer(context.Context, *connect.Request[v1.DeleteAnswerRequest]) (*connect.Response[v1.DeleteAnswerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("backend.v1.ContestService.DeleteAnswer is not implemented"))
}