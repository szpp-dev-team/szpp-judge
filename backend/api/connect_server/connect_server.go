package connect_server

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/rs/cors"
	"github.com/szpp-dev-team/szpp-judge/backend/api/connect_server/interceptor"
	connect_interfaces "github.com/szpp-dev-team/szpp-judge/backend/interfaces/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/auth"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/contests"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/judge"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/tasks"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/user"
	"github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1/backendv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func New(addr string, opts ...optionFunc) *http.Server {
	opt := defaultOption()
	for _, f := range opts {
		f(opt)
	}

	interceptors := connect.WithInterceptors(
		interceptor.Auth([]byte(opt.Secret)),
		interceptor.Logger(),
	)

	mux := http.NewServeMux()

	healthcheckSrv := connect_interfaces.NewHealthcheckServiceServer()
	mux.Handle(backendv1connect.NewHealthcheckServiceHandler(healthcheckSrv))
	userSrv := connect_interfaces.NewUserServiceServer(user.NewInteractor(opt.EntClient))
	mux.Handle(backendv1connect.NewUserServiceHandler(userSrv, interceptors))
	taskInteractor := tasks.NewInteractor(opt.EntClient, opt.TestcasesRepository)
	taskSrv := connect_interfaces.NewTaskServiceServer(taskInteractor)
	mux.Handle(backendv1connect.NewTaskServiceHandler(taskSrv, interceptors))
	judgeInteractor := judge.NewInteractor(opt.JudgeClient, opt.EntClient, opt.SourcesRepository, opt.judgeQueue)
	judgeSrv := connect_interfaces.NewJudgeServiceServer(judgeInteractor)
	mux.Handle(backendv1connect.NewJudgeServiceHandler(judgeSrv, interceptors))
	contestInteractor := contests.NewInteractor(opt.EntClient)
	contestSrv := connect_interfaces.NewContestServiceServer(contestInteractor)
	mux.Handle(backendv1connect.NewContestServiceHandler(contestSrv, interceptors))
	authSrv := connect_interfaces.NewAuthServiceServer(auth.NewInteractor(opt.EntClient, opt.Secret))
	mux.Handle(backendv1connect.NewAuthServiceHandler(authSrv, interceptors))

	// https://connectrpc.com/docs/go/deployment/#cors
	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{"http://localhost:5000", "http://localhost:3000", opt.FrontendURL},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Connect-Accept-Encoding",  // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Timeout",             // Used for gRPC-web
			"X-Grpc-Web",               // Used for gRPC-web
			"X-User-Agent",             // Used for gRPC-web
			"Authorization",
		},
		ExposedHeaders: []string{
			"Content-Encoding",         // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Status",              // Required for gRPC-web
			"Grpc-Message",             // Required for gRPC-web
		},
	})

	return &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(
			corsHandler.Handler(mux),
			&http2.Server{},
		),
	}
}
