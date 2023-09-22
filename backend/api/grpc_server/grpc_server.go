package grpc_server

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_interfaces "github.com/szpp-dev-team/szpp-judge/backend/interfaces/grpc"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/tasks"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/user"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(opts ...optionFunc) *grpc.Server {
	opt := defaultOption()
	for _, f := range opts {
		f(opt)
	}

	serverOptions := make([]grpc.ServerOption, 0)
	// set logging interceptor
	serverOptions = append(serverOptions,
		grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(interceptorLogger(opt.logger))),
		grpc.ChainStreamInterceptor(logging.StreamServerInterceptor(interceptorLogger(opt.logger))),
	)

	srv := grpc.NewServer(serverOptions...)
	if opt.useReflection {
		reflection.Register(srv)
	}
	healthcheckSrv := grpc_interfaces.NewHealthcheckServiceServer()
	pb.RegisterHealthcheckServiceServer(srv, healthcheckSrv)
	userSrv := grpc_interfaces.NewUserServiceServer(user.NewInteractor(opt.entClient))
	pb.RegisterUserServiceServer(srv, userSrv)
	taskInteractor := tasks.NewInteractor(opt.entClient, opt.testcasesRepository)
	taskSrv := grpc_interfaces.NewTaskServiceServer(taskInteractor)
	pb.RegisterTaskServiceServer(srv, taskSrv)

	return srv
}

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
