package api

import (
	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"golang.org/x/exp/slog"
)

type option struct {
	Logger           *slog.Logger
	EntClient        *ent.Client
	CloudtasksClient *cloudtasks.Client
	UseReflection    bool
	JudgeClient      judgev1.JudgeServiceClient
}

func DefaultOption() *option {
	return &option{
		Logger: slog.Default(),
	}
}

type OptionFunc func(*option)

func WithLogger(logger *slog.Logger) OptionFunc {
	return func(o *option) {
		o.Logger = logger
	}
}

// publish grpc server information(method, service, etc.)
func WithReflection(b bool) OptionFunc {
	return func(o *option) {
		o.UseReflection = b
	}
}

func WithEntClient(c *ent.Client) OptionFunc {
	return func(o *option) {
		o.EntClient = c
	}
}

func WithCloudtasksClient(c *cloudtasks.Client) OptionFunc {
	return func(o *option) {
		o.CloudtasksClient = c
	}
}

func WithJudgeClient(judgeClient judgev1.JudgeServiceClient) OptionFunc {
	return func(o *option) {
		o.JudgeClient = judgeClient
	}
}
