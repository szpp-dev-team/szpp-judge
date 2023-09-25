package grpc_server

import (
	"context"

	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/szpp-dev-team/szpp-judge/backend/api"
	grpc_interfaces "github.com/szpp-dev-team/szpp-judge/backend/interfaces/grpc"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/contests"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/judge"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/tasks"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/user"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(opts ...api.OptionFunc) *grpc.Server {
	opt := api.DefaultOption()
	for _, f := range opts {
		f(opt)
	}

	serverOptions := make([]grpc.ServerOption, 0)
	// set logging interceptor
	serverOptions = append(serverOptions,
		grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(interceptorLogger(opt.Logger))),
		grpc.ChainStreamInterceptor(logging.StreamServerInterceptor(interceptorLogger(opt.Logger))),
	)

	srv := grpc.NewServer(serverOptions...)
	if opt.UseReflection {
		reflection.Register(srv)
	}
	healthcheckSrv := grpc_interfaces.NewHealthcheckServiceServer()
	pb.RegisterHealthcheckServiceServer(srv, healthcheckSrv)
	userSrv := grpc_interfaces.NewUserServiceServer(user.NewInteractor(opt.EntClient))
	pb.RegisterUserServiceServer(srv, userSrv)
	taskInteractor := tasks.NewInteractor(opt.EntClient, opt.TestcasesRepository)
	taskSrv := grpc_interfaces.NewTaskServiceServer(taskInteractor)
	pb.RegisterTaskServiceServer(srv, taskSrv)
	judgeInteractor := judge.NewInteractor(opt.JudgeClient, opt.EntClient)
	judgeSrv := grpc_interfaces.NewJudgeServiceServer(judgeInteractor)
	pb.RegisterJudgeServiceServer(srv, judgeSrv)
	contestInteractor := contests.NewInteractor(opt.EntClient)
	contetSrv := grpc_interfaces.NewContestServiceServer(contestInteractor)
	pb.RegisterContestServiceServer(srv, contetSrv)

	return srv
}

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
