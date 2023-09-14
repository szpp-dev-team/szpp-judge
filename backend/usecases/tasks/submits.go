package tasks

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/samber/lo"
	"github.com/szpp-dev-team/szpp-judge/backend/core/entutil"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_language "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/language"
	ent_submit "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	ent_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	ent_user "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/sources"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Interactor) PostSubmit(ctx context.Context, req *backendv1.SubmitRequest) (*backendv1.SubmitResponse, error) {
	task, err := i.entClient.Task.Query().WithTestcases().Where(ent_task.ID(int(req.TaskId))).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the task was not found")
		}
		i.logger.Error("failed to get the task", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	langID, err := i.entClient.Language.Query().Where(ent_language.SlugEQ(req.LangId)).OnlyID(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the task was not found")
		}
		i.logger.Error("failed to get the language", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	var submit *ent.Submit
	if err := entutil.WithTx(ctx, i.entClient, func(tx *ent.Tx) error {
		now := timejst.Now()
		submit, err = tx.Submit.Create().
			SetSubmittedAt(now).
			SetCreatedAt(now).
			SetTask(task).
			SetLanguageID(langID).
			Save(ctx)
		if err != nil {
			i.logger.Error("failed to get create a submit", slog.Any("error", err))
			return status.Error(codes.Internal, err.Error())
		}

		if err := i.sourcesRepo.UploadSource(ctx, submit.ID, []byte(req.SourceCode)); err != nil {
			i.logger.Error("failed to upload source code to the storage", slog.Any("error", err))
			return status.Error(codes.Internal, err.Error())
		}

		req, err := buildJudgeRequest(submit.ID, req.LangId, task)
		if err != nil {
			i.logger.Error("[CRETICAL ERROR] inconsistency was found through building JudgeRequest", slog.Any("error", err), slog.Int("taskID", task.ID))
			return status.Error(codes.Internal, err.Error())
		}
		if err := i.judgeQueue.Push(ctx, submit.ID, req); err != nil {
			i.logger.Error("failed to push a judge task to the queue", slog.Any("error", err))
			return status.Error(codes.Internal, err.Error())
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
		WithTestcaseResults(func(trq *ent.TestcaseResultQuery) {
			trq.WithTestcase()
		}).
		Where(ent_submit.ID(int(req.Id))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the submission was not found")
		}
		i.logger.Error("failed to get the submission", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	source, err := i.sourcesRepo.DownloadSource(ctx, submit.ID)
	if err != nil {
		i.logger.Error("failed to download source code from the storage", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	var submissionDetail *backendv1.SubmissionDetail
	if submit.Status == nil {
		submissionDetail = &backendv1.SubmissionDetail{
			Id:         int32(submit.ID),
			UserId:     0,   // TODO: set
			Username:   "",  // TODO: set
			ContestId:  nil, // TODO: set
			TaskId:     int32(submit.Edges.Task.ID),
			TaskTitle:  submit.Edges.Task.Title,
			LangId:     submit.Edges.Language.Slug,
			SourceCode: string(source),
		}
	} else {
		testcaseResults := make([]*backendv1.TestcaseResult, 0, len(submit.Edges.TestcaseResults))
		for _, tcResult := range submit.Edges.TestcaseResults {
			testcaseResults = append(testcaseResults, &backendv1.TestcaseResult{
				TestcaseName:  tcResult.Edges.Testcase.Name,
				JudgeStatus:   judgev1.JudgeStatus(judgev1.JudgeStatus_value[tcResult.Status]),
				ExecTimeMs:    uint32(tcResult.ExecTime),
				ExecMemoryKib: uint32(tcResult.ExecMemory),
			})
		}
		submissionDetail = &backendv1.SubmissionDetail{
			Id:              int32(submit.ID),
			UserId:          0,   // TODO: set
			Username:        "",  // TODO: set
			ContestId:       nil, // TODO: set
			TaskId:          int32(submit.Edges.Task.ID),
			TaskTitle:       submit.Edges.Task.Title,
			Score:           int32(submit.Score),
			LangId:          submit.Edges.Language.Slug,
			SourceCode:      string(source),
			Status:          lo.ToPtr(judgev1.JudgeStatus(judgev1.JudgeStatus_value[*submit.Status])),
			ExecTimeMs:      lo.ToPtr(uint32(submit.ExecTime)),
			ExecMemoryKib:   lo.ToPtr(uint32(submit.ExecMemory)),
			TestcaseResults: testcaseResults,
		}
	}

	return &backendv1.GetSubmissionDetailResponse{
		SubmissionDetail: submissionDetail,
	}, nil
}

func (i *Interactor) ListSubmissions(ctx context.Context, req *backendv1.ListSubmissionsRequest) (*backendv1.ListSubmissionsResponse, error) {
	q := i.entClient.Submit.Query().WithLanguage().WithTask()
	if req.ContestId != nil {
		// TODO: HasContest
	}
	if req.UserId != nil {
		q = q.Where(ent_submit.HasUserWith(ent_user.ID(int(*req.UserId))))
	}
	submits, err := q.All(ctx)
	if err != nil {
		i.logger.Error("failed to get the submissions", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	submissionList := make([]*backendv1.SubmissionSummary, 0, len(submits))
	for _, submit := range submits {
		var judgeStatus *judgev1.JudgeStatus
		if submit.Status != nil {
			judgeStatus = lo.ToPtr(judgev1.JudgeStatus(judgev1.JudgeStatus_value[*submit.Status]))
		}
		submissionList = append(submissionList, &backendv1.SubmissionSummary{
			Id:            int32(submit.ID),
			UserId:        0,   // TODO: set
			Username:      "",  // TODO: set
			ContestId:     nil, // TODO: set
			TaskId:        int32(submit.Edges.Task.ID),
			TaskTitle:     submit.Edges.Task.Title,
			Score:         int32(submit.Score),
			LangId:        submit.Edges.Language.Slug,
			JudgeStatus:   judgeStatus,
			ExecTimeMs:    lo.ToPtr(uint32(submit.ExecTime)),
			ExecMemoryKib: lo.ToPtr(uint32(submit.ExecMemory)),
			SubmittedAt:   timestamppb.New(submit.SubmittedAt),
		})
	}

	return &backendv1.ListSubmissionsResponse{
		Submissions: submissionList,
	}, nil
}

func buildJudgeRequest(submitID int, langID string, task *ent.Task) (*judgev1.JudgeRequest, error) {
	var judgeType *judgev1.JudgeType
	switch task.JudgeType {
	case ent_task.JudgeTypeNormal:
		judgeType = &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Normal{
				Normal: &judgev1.JudgeTypeNormal{
					CaseInsensitive: task.CaseInsensitive,
				},
			},
		}
	case ent_task.JudgeTypeEps:
		if task.Ndigits == nil {
			return nil, fmt.Errorf("the type of task is JudgeTypeEps, but the field ndigits is not set")
		}
		judgeType = &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Eps{
				Eps: &judgev1.JudgeTypeEPS{
					Ndigits: uint32(*task.Ndigits),
				},
			},
		}
	case ent_task.JudgeTypeCustom:
		if task.JudgeCodePath == nil {
			return nil, fmt.Errorf("the type of task is JudgeTypeCustom, but the field judge_code_path is not set")
		}
		judgeType = &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Custom{
				Custom: &judgev1.JudgeTypeCustom{
					JudgeCodePath: *task.JudgeCodePath,
				},
			},
		}
	case ent_task.JudgeTypeInteractive:
		if task.JudgeCodePath == nil {
			return nil, fmt.Errorf("the type of task is JudgeTypeInteractive, but the field judge_code_path is not set")
		}
		judgeType = &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Interactive{
				Interactive: &judgev1.JudgeTypeInteractive{
					JudgeCodePath: *task.JudgeCodePath,
				},
			},
		}
	default:
		return nil, fmt.Errorf("unknown judge type: %s", task.JudgeType)
	}

	testcaseList := make([]*judgev1.Testcase, 0, len(task.Edges.Testcases))
	for _, tc := range task.Edges.Testcases {
		testcaseList = append(testcaseList, &judgev1.Testcase{
			InputPath:    testcases.BuildTestcaseInPath(task.ID, tc.Name),
			ExpectedPath: testcases.BuildTestcaseOutPath(task.ID, tc.Name),
		})
	}

	return &judgev1.JudgeRequest{
		SourceCodePath:     sources.BuildSourceCodePath(submitID),
		LangId:             langID,
		JudgeType:          judgeType,
		ExecTimeLimitMs:    uint32(task.ExecMemoryLimit),
		ExecMemoryLimitMib: uint32(task.ExecMemoryLimit),
		Testcases:          testcaseList,
		WantResultDetail:   lo.ToPtr(true), // TODO: これが何か確認する
	}, nil
}
