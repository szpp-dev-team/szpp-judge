package judge

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/samber/lo"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_submit "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
)

type Interactor struct {
	logger      *slog.Logger
	judgeClient judgev1.JudgeServiceClient
	entClient   *ent.Client
}

func NewInteractor(judgeClient judgev1.JudgeServiceClient, entClient *ent.Client) *Interactor {
	logger := slog.Default().With(slog.String("usecase", "judge"))
	return &Interactor{
		logger:      logger,
		judgeClient: judgeClient,
		entClient:   entClient,
	}
}

// Cloud Tasks からのみ呼び出される
func (i *Interactor) PostJudgeRequest(ctx context.Context, req *judgev1.JudgeRequest) error {
	stream, err := i.judgeClient.Judge(ctx, req)
	if err != nil {
		if err := i.updateSubmitStatusIE(ctx, int(req.SubmissionId)); err != nil {
			i.logger.Error("failed to update submit status", slog.Int("submissionID", int(req.SubmissionId)), slog.Any("error", err))
			return err
		}
		i.logger.Error("failed to execute judge rpc", slog.Any("error", err))
		return err
	}

	testcaseResults := make([]*ent.TestcaseResult, 0)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			if err := stream.CloseSend(); err != nil {
				if err := i.updateSubmitStatusIE(ctx, int(req.SubmissionId)); err != nil {
					i.logger.Error("failed to update submit status", slog.Int("submissionID", int(req.SubmissionId)), slog.Any("error", err))
					return err
				}
				i.logger.Error("failed to close stream", slog.Int("submissionID", int(req.SubmissionId)), slog.Any("error", err))
				return err
			}
			break
		}

		res, err := i.entClient.TestcaseResult.Create().
			SetStatus(resp.Status.String()).
			SetExecTime(int(resp.ExecTimeMs)).
			SetExecMemory(int(resp.ExecMemoryKib)).
			SetSubmitID(int(req.SubmissionId)).
			Save(ctx)
		if err != nil {
			if err := i.updateSubmitStatusIE(ctx, int(req.SubmissionId)); err != nil {
				i.logger.Error("failed to update submit status", slog.Int("submissionID", int(req.SubmissionId)), slog.Any("error", err))
				return err
			}
			i.logger.Error("failed to create testcase result", slog.Int("submissionID", int(req.SubmissionId)), slog.Any("error", err))
			return err
		}
		testcaseResults = append(testcaseResults, res)
	}

	if err := i.updateSubmitResult(ctx, int(req.SubmissionId), testcaseResults); err != nil {
		if err := i.updateSubmitStatusIE(ctx, int(req.SubmissionId)); err != nil {
			i.logger.Error("failed to update submit status", slog.Int("submissionID", int(req.SubmissionId)), slog.Any("error", err))
			return err
		}
		i.logger.Error("failed to set submit result", slog.Int("submissionID", int(req.SubmissionId)), slog.Any("error", err))
		return err
	}

	return nil
}

func (i *Interactor) updateSubmitResult(ctx context.Context, submitID int, testcaseResults []*ent.TestcaseResult) error {
	submit, err := i.entClient.Submit.Query().
		WithTask(func(tq *ent.TaskQuery) {
			tq.WithTestcaseSets(func(tsq *ent.TestcaseSetQuery) {
				tsq.WithTestcases()
			})
		}).
		Where(ent_submit.ID(submitID)).Only(ctx)
	if err != nil {
		i.logger.Error("failed to get submit", slog.Any("error", err))
		return err
	}

	// Status, ExecMemory, ExecTime を更新
	for _, tr := range testcaseResults {
		judgeStatusPriority := priorityByJudgeStatus[judgev1.JudgeStatus(judgev1.JudgeStatus_value[*submit.Status])]
		trStatusPriority := priorityByJudgeStatus[judgev1.JudgeStatus(judgev1.JudgeStatus_value[tr.Status])]
		if submit.Status == nil || judgeStatusPriority < trStatusPriority {
			submit.Status = &tr.Status
			submit.ExecMemory = tr.ExecMemory
			submit.ExecTime = tr.ExecTime
		}
	}

	// 得点計算
	score, err := i.calculateScore(ctx, submit.Edges.Task.Edges.TestcaseSets, testcaseResults)
	if err != nil {
		i.logger.Error("[!!!inconsistency!!!]", slog.Int("submitID", submitID), slog.Int("taskID", submit.Edges.Task.ID), slog.Any("error", err))
		return err
	}
	submit.Score = score

	submit.UpdatedAt = lo.ToPtr(timejst.Now())
	if _, err := submit.Update().Save(ctx); err != nil {
		i.logger.Error("failed to update submit", slog.Any("error", err))
		return err
	}
	return nil
}

// testcaseSets has edges to Testcase
func (i *Interactor) calculateScore(ctx context.Context, testcaseSets []*ent.TestcaseSet, testcaseResults []*ent.TestcaseResult) (int, error) {
	isACByTestcaseName := make(map[string]bool)
	for _, testcaseResult := range testcaseResults {
		isACByTestcaseName[testcaseResult.Edges.Testcase.Name] = testcaseResult.Status == judgev1.JudgeStatus_AC.String()
	}

	score := 0
	for _, testcaseSet := range testcaseSets {
		isAllAC := true
		for _, testcase := range testcaseSet.Edges.Testcases {
			isAC, ok := isACByTestcaseName[testcase.Name]
			if !ok {
				return 0, fmt.Errorf("testcase \"%s\" is not found in the testcaseSet \"%s\"", testcase.Name, testcaseSet.Name)
			}
			if isAC {
				isAllAC = false
				break
			}
		}
		if isAllAC {
			score += testcaseSet.Score
		}
	}

	return score, nil
}

func (i *Interactor) updateSubmitStatusIE(ctx context.Context, submitID int) error {
	submit, err := i.entClient.Submit.UpdateOneID(submitID).SetStatus(judgev1.JudgeStatus_IE.String()).Save(ctx)
	if err != nil {
		return err
	}
	i.logger.Info("[IE] updated submit status", slog.Any("submit", submit))
	return nil
}

var priorityByJudgeStatus = map[judgev1.JudgeStatus]int{
	judgev1.JudgeStatus_AC:  0,
	judgev1.JudgeStatus_TLE: 1,
	judgev1.JudgeStatus_MLE: 2,
	judgev1.JudgeStatus_OLE: 3,
	judgev1.JudgeStatus_WA:  4,
	judgev1.JudgeStatus_RE:  5,
	judgev1.JudgeStatus_CE:  6,
	judgev1.JudgeStatus_IE:  7,
}
