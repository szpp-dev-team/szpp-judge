package grpc_server

import (
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	"golang.org/x/exp/slog"
)

type option struct {
	logger              *slog.Logger
	entClient           *ent.Client
	useReflection       bool
	testcasesRepository testcases.Repository
}

func defaultOption() *option {
	return &option{
		logger: slog.Default(),
	}
}

type optionFunc func(*option)

func WithLogger(logger *slog.Logger) optionFunc {
	return func(o *option) {
		o.logger = logger
	}
}

// publish grpc server information(method, service, etc.)
func WithReflection(b bool) optionFunc {
	return func(o *option) {
		o.useReflection = b
	}
}

func WithEntClient(c *ent.Client) optionFunc {
	return func(o *option) {
		o.entClient = c
	}
}

func WithTestcasesRepository(r testcases.Repository) optionFunc {
	return func(o *option) {
		o.testcasesRepository = r
	}
}
