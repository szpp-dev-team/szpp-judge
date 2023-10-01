package connect_server

import (
	"log/slog"

	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/judge_queue"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/sources"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
)

type option struct {
	Logger              *slog.Logger
	EntClient           *ent.Client
	UseReflection       bool
	TestcasesRepository testcases.Repository
	SourcesRepository   sources.Repository
	judgeQueue          judge_queue.JudgeQueue
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

func WithEntClient(c *ent.Client) optionFunc {
	return func(o *option) {
		o.EntClient = c
	}
}

func WithSourcesRepository(r sources.Repository) optionFunc {
	return func(o *option) {
		o.SourcesRepository = r
	}
}

func WithTestcasesRepository(r testcases.Repository) optionFunc {
	return func(o *option) {
		o.TestcasesRepository = r
	}
}

func WithJudgeQueue(q judge_queue.JudgeQueue) optionFunc {
	return func(o *option) {
		o.judgeQueue = q
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
