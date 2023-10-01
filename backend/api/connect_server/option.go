package connect_server

import (
	"log/slog"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
)

type option struct {
	Logger              *slog.Logger
	EntClient           *ent.Client
	CloudtasksClient    *cloudtasks.Client
	UseReflection       bool
	TestcasesRepository testcases.Repository
	JudgeClient         judgev1.JudgeServiceClient
	Secret              string
}

func defaultOption() *option {
	return &option{
		Logger: slog.Default(),
	}
}

type optionFunc func(*option)

func WithLogger(logger *slog.Logger) optionFunc {
	return func(o *option) {
		o.Logger = logger
	}
}

// publish grpc server information(method, service, etc.)
func WithReflection(b bool) optionFunc {
	return func(o *option) {
		o.UseReflection = b
	}
}

func WithEntClient(c *ent.Client) optionFunc {
	return func(o *option) {
		o.EntClient = c
	}
}

func WithCloudtasksClient(c *cloudtasks.Client) optionFunc {
	return func(o *option) {
		o.CloudtasksClient = c
	}
}

func WithTestcasesRepository(r testcases.Repository) optionFunc {
	return func(o *option) {
		o.TestcasesRepository = r
	}
}

func WithJudgeClient(c judgev1.JudgeServiceClient) optionFunc {
	return func(o *option) {
		o.JudgeClient = c
	}
}

func WithSecret(secret string) optionFunc {
	return func(o *option) {
		o.Secret = secret
	}
}
