package judge

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/samber/lo"
	"github.com/szpp-dev-team/szpp-judge/backend/api/connect_server/interceptor"
	"github.com/szpp-dev-team/szpp-judge/backend/core/entutil"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_contest "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	ent_language "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/language"
	ent_submit "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	ent_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	ent_user "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/sources"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Interactor) PostSubmit(ctx context.Context, req *backendv1.SubmitRequest) (*backendv1.SubmitResponse, error) {
	claims := interceptor.GetClaimsFromContext(ctx)

	task, err := i.entClient.Task.Query().WithTestcases().Where(ent_task.ID(int(req.TaskId))).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the task was not found"))
		}
		i.logger.Error("failed to get the task", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	langID, err := i.entClient.Language.Query().Where(ent_language.SlugEQ(req.LangId)).OnlyID(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the task was not found"))
		}
		i.logger.Error("failed to get the language", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var submit *ent.Submit
	if err := entutil.WithTx(ctx, i.entClient, func(tx *ent.Tx) error {
		now := timejst.Now()
		q := tx.Submit.Create().
			SetSubmittedAt(now).
			SetCreatedAt(now).
			SetTask(task).
			SetLanguageID(langID).
			SetUserID(claims.UserID)
		if req.ContestId != nil {
			q.SetContestID(int(*req.ContestId))
		}

		submit, err = q.Save(ctx)
		if err != nil {
			i.logger.Error("failed to get create a submit", slog.Any("error", err))
			return connect.NewError(connect.CodeInternal, err)
		}

		if err := i.sourcesRepo.UploadSource(ctx, submit.ID, []byte(req.SourceCode)); err != nil {
			i.logger.Error("failed to upload source code to the storage", slog.Any("error", err))
			return connect.NewError(connect.CodeInternal, err)
		}

		req, err := buildJudgeRequest(submit.ID, req.LangId, task)
		if err != nil {
			i.logger.Error("[CRETICAL ERROR] inconsistency was found through building JudgeRequest", slog.Any("error", err), slog.Int("taskID", task.ID))
			return connect.NewError(connect.CodeInternal, err)
		}
		if err := i.judgeQueue.Push(ctx, req); err != nil {
			i.logger.Error("failed to push a judge task to the queue", slog.Any("error", err))
			return connect.NewError(connect.CodeInternal, err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &backendv1.SubmitResponse{
		SubmissionId: int32(submit.ID),
	}, nil
}

func (i *Interactor) GetSubmissionDetail(ctx context.Context, req *backendv1.GetSubmissionDetailRequest) (*backendv1.GetSubmissionDetailResponse, error) {
	submit, err := i.entClient.Submit.Query().
		WithLanguage().
		WithTask().
		WithUser().
		WithContest().
		WithTestcaseResults(func(trq *ent.TestcaseResultQuery) {
			trq.WithTestcase()
		}).
		Where(ent_submit.ID(int(req.Id))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the submission was not found"))
		}
		i.logger.Error("failed to get the submission", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	source, err := i.sourcesRepo.DownloadSource(ctx, submit.ID)
	if err != nil {
		i.logger.Error("failed to download source code from the storage", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var submissionDetail *backendv1.SubmissionDetail
	var contestID *int32
	if submit.Edges.Contest != nil {
		contestID = lo.ToPtr(int32(submit.Edges.Contest.ID))
	}
	var updatedAt *timestamppb.Timestamp
	if submit.UpdatedAt != nil {
		updatedAt = timestamppb.New(*submit.UpdatedAt)
	}
	submissionDetail = &backendv1.SubmissionDetail{
		Id:          int32(submit.ID),
		UserId:      int32(submit.Edges.User.ID),
		Username:    submit.Edges.User.Username,
		ContestId:   contestID,
		TaskId:      int32(submit.Edges.Task.ID),
		TaskTitle:   submit.Edges.Task.Title,
		LangId:      submit.Edges.Language.Slug,
		SourceCode:  string(source),
		SubmittedAt: timestamppb.New(submit.SubmittedAt),
		CreatedAt:   timestamppb.New(submit.CreatedAt),
		UpdatedAt:   updatedAt,
	}
	if submit.Status != nil {
		testcaseResults := make([]*backendv1.TestcaseResult, 0, len(submit.Edges.TestcaseResults))
		for _, tcResult := range submit.Edges.TestcaseResults {
			testcaseResults = append(testcaseResults, &backendv1.TestcaseResult{
				TestcaseName:  tcResult.Edges.Testcase.Name,
				JudgeStatus:   judgev1.JudgeStatus(judgev1.JudgeStatus_value[tcResult.Status]),
				ExecTimeMs:    uint32(tcResult.ExecTime),
				ExecMemoryKib: uint32(tcResult.ExecMemory),
			})
		}
		submissionDetail.Score = int32(submit.Score)
		submissionDetail.Status = judgev1.JudgeStatus(judgev1.JudgeStatus_value[*submit.Status])
		submissionDetail.ExecTimeMs = lo.ToPtr(uint32(submit.ExecTime))
		submissionDetail.ExecMemoryKib = lo.ToPtr(uint32(submit.ExecMemory))
		submissionDetail.TestcaseResults = testcaseResults
	}

	return &backendv1.GetSubmissionDetailResponse{
		SubmissionDetail: submissionDetail,
	}, nil
}

func (i *Interactor) ListSubmissions(ctx context.Context, req *backendv1.ListSubmissionsRequest) (*backendv1.ListSubmissionsResponse, error) {
	now := timejst.Now()

	claims := interceptor.GetClaimsFromContext(ctx)
	isAdmin := false
	if claims != nil {
		user, err := i.entClient.User.Query().Where(ent_user.Username(claims.Username)).Only(ctx)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		isAdmin = user.Role == backendv1.Role_ADMIN.String()
	}

	contest, err := i.entClient.Contest.Query().
		WithSubmits(func(sq *ent.SubmitQuery) {
			sq.WithUser()
		}).
		Where(ent_contest.ID(int(*req.ContestId))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the contest was not found"))
		}
		i.logger.Error("failed to get the contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	// admin もしくはコンテスト終了後は全ての提出の閲覧が可能
	if isAdmin || now.After(contest.EndAt) {
		return &backendv1.ListSubmissionsResponse{
			Submissions: lo.Map(contest.Edges.Submits, func(s *ent.Submit, _ int) *backendv1.SubmissionSummary {
				return toPbSubmissionSummary(s)
			}),
		}, nil
	}
	// コンテスト開始前は permissionDenied として扱う
	if now.Before(contest.StartAt) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("contest is not held"))
	}

	// 未ログインの場合は空を返す
	if claims == nil {
		return &backendv1.ListSubmissionsResponse{}, nil
	}

	// コンテスト開催中は自分の提出のみ返す
	return &backendv1.ListSubmissionsResponse{
		Submissions: lo.FilterMap(contest.Edges.Submits, func(s *ent.Submit, _ int) (*backendv1.SubmissionSummary, bool) {
			if s.Edges.User.Username != claims.Username {
				return nil, false
			}
			return toPbSubmissionSummary(s), true
		}),
	}, nil
}

func toPbSubmissionSummary(submit *ent.Submit) *backendv1.SubmissionSummary {
	var judgeStatus judgev1.JudgeStatus
	if submit.Status != nil {
		judgeStatus = judgev1.JudgeStatus(judgev1.JudgeStatus_value[*submit.Status])
	}
	var contestID *int32
	if submit.Edges.Contest != nil {
		contestID = lo.ToPtr(int32(submit.Edges.Contest.ID))
	}
	return &backendv1.SubmissionSummary{
		Id:            int32(submit.ID),
		UserId:        int32(submit.Edges.User.ID),
		Username:      submit.Edges.User.Username,
		ContestId:     contestID,
		TaskId:        int32(submit.Edges.Task.ID),
		TaskTitle:     submit.Edges.Task.Title,
		Score:         int32(submit.Score),
		LangId:        submit.Edges.Language.Slug,
		JudgeStatus:   judgeStatus,
		ExecTimeMs:    lo.ToPtr(uint32(submit.ExecTime)),
		ExecMemoryKib: lo.ToPtr(uint32(submit.ExecMemory)),
		SubmittedAt:   timestamppb.New(submit.SubmittedAt),
	}
}

func buildJudgeRequest(submitID int, langID string, task *ent.Task) (*judgev1.JudgeRequest, error) {
	testcaseList := make([]*judgev1.Testcase, 0, len(task.Edges.Testcases))
	for _, tc := range task.Edges.Testcases {
		testcaseList = append(testcaseList, &judgev1.Testcase{
			Id:           uint32(tc.ID),
			InputPath:    testcases.BuildTestcaseInPath(task.ID, tc.Name),
			ExpectedPath: testcases.BuildTestcaseOutPath(task.ID, tc.Name),
		})
	}

	return &judgev1.JudgeRequest{
		SourceCodePath:     sources.BuildSourceCodePath(submitID),
		LangId:             langID,
		JudgeType:          nil,
		ExecTimeLimitMs:    uint32(task.ExecMemoryLimit),
		ExecMemoryLimitMib: uint32(task.ExecMemoryLimit),
		Testcases:          testcaseList,
		WantResultDetail:   lo.ToPtr(true), // TODO: これが何か確認する
		StdoutLimitKib:     lo.ToPtr(uint32(512)),
		StderrLimitKib:     lo.ToPtr(uint32(512)),
		SubmissionId:       int32(submitID),
		CheckerCodePath:    fmt.Sprintf("tasks/%d/checker.cpp", task.ID),
	}, nil
}
