package tasks

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/samber/lo"
	"github.com/szpp-dev-team/szpp-judge/backend/core/entutil"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	enttask "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	testcases_repo "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient  *ent.Client
	repository testcases.Repository
	logger     *slog.Logger
}

func NewInteractor(entClient *ent.Client, repository testcases.Repository) *Interactor {
	logger := slog.Default().With(slog.String("usecase", "tasks"))
	return &Interactor{entClient, repository, logger}
}

func (i *Interactor) CreateTask(ctx context.Context, req *backendv1.CreateTaskRequest) (*backendv1.CreateTaskResponse, error) {
	var (
		task         *ent.Task
		testcaseSets []*ent.TestcaseSet
		testcases    []*ent.Testcase
	)
	if err := entutil.WithTx(ctx, i.entClient, func(tx *ent.Tx) error {
		now := timejst.Now()
		var err error

		// Testcase
		testcaseBuilders := make([]*ent.TestcaseCreate, 0, len(req.Testcases))
		for _, testcase := range req.Testcases {
			q := tx.Testcase.Create().
				SetName(testcase.Slug).
				SetNillableDescription(testcase.Description).
				SetCreatedAt(now)
			testcaseBuilders = append(testcaseBuilders, q)
		}
		testcases, err = tx.Testcase.CreateBulk(testcaseBuilders...).Save(ctx)
		if err != nil {
			i.logger.Error("failed to create testcases", slog.Any("error", err))
			return err
		}
		testcaseByName := lo.Associate(testcases, func(t *ent.Testcase) (string, *ent.Testcase) {
			return t.Name, t
		})

		// TestcaseSet
		testcaseSetBuilders := make([]*ent.TestcaseSetCreate, 0, len(req.TestcaseSets))
		for _, testcaseSet := range req.TestcaseSets {
			testcaseIDs := make([]int, 0, len(testcaseSet.TestcaseSlugs))
			for _, name := range testcaseSet.TestcaseSlugs {
				testcaseIDs = append(testcaseIDs, testcaseByName[name].ID)
			}
			q := tx.TestcaseSet.Create().
				SetName(testcaseSet.Slug).
				SetScore(int(testcaseSet.Score)).
				SetIsSample(testcaseSet.IsSample).
				AddTestcaseIDs(testcaseIDs...).
				SetCreatedAt(now)
			testcaseSetBuilders = append(testcaseSetBuilders, q)
		}
		testcaseSets, err = tx.TestcaseSet.CreateBulk(testcaseSetBuilders...).Save(ctx)
		if err != nil {
			i.logger.Error("failed to create testcase_sets", slog.Any("error", err))
			return err
		}

		// Task
		q := tx.Task.Create().
			SetTitle(req.Task.Title).
			SetStatement(req.Task.Statement).
			SetDifficulty(req.Task.Difficulty.String()).
			SetExecTimeLimit(uint(req.Task.ExecTimeLimit)).
			SetExecMemoryLimit(uint(req.Task.ExecMemoryLimit)).
			AddTestcases(testcases...).
			AddTestcaseSets(testcaseSets...).
			SetCreatedAt(now)
			// TODO: SetUser
		switch ty := req.Task.JudgeType.JudgeType.(type) {
		case *judgev1.JudgeType_Normal:
			q.SetJudgeType(enttask.JudgeTypeNormal)
			q.SetCaseInsensitive(*ty.Normal.CaseInsensitive)
		case *judgev1.JudgeType_Eps:
			q.SetJudgeType(enttask.JudgeTypeEps)
			q.SetNdigits(uint(ty.Eps.Ndigits))
		case *judgev1.JudgeType_Interactive:
			q.SetJudgeType(enttask.JudgeTypeInteractive)
			q.SetJudgeCodePath(ty.Interactive.JudgeCodePath)
		case *judgev1.JudgeType_Custom:
			q.SetJudgeType(enttask.JudgeTypeCustom)
			q.SetJudgeCodePath(ty.Custom.JudgeCodePath)
		default:
			return fmt.Errorf("unrecognized JudgeType: %T", ty)
		}
		task, err = q.Save(ctx)
		if err != nil {
			i.logger.Error("failed to create task", slog.Any("error", err))
			return err
		}

		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	i.logger.Info("succeeded to register task information", slog.Int("task_id", task.ID))

	for _, testcase := range req.Testcases {
		if err := i.repository.UploadTestcase(ctx, task.ID, &testcases_repo.Testcase{
			Name: testcase.Slug,
			In:   []byte(testcase.Input),
			Out:  []byte(testcase.Output),
		}); err != nil {
			i.logger.Error("failed to upload a testcase", slog.Any("error", err))
			return nil, err
		}
	}

	i.logger.Info("succeeded to upload testcases", slog.Int("task_id", task.ID))

	return &backendv1.CreateTaskResponse{
		Task: toPbTask(task),
		Testcases: lo.Map(testcases, func(t *ent.Testcase, i int) *backendv1.Testcase {
			return toPbTestcase(t, []byte(req.Testcases[i].Input), []byte(req.Testcases[i].Output))
		}),
		TestcaseSets: lo.Map(testcaseSets, func(t *ent.TestcaseSet, _ int) *backendv1.TestcaseSet {
			return toPbTestcaseSet(t)
		}),
	}, nil
}

func (i *Interactor) GetTask(ctx context.Context, req *backendv1.GetTaskRequest) (*backendv1.GetTaskResponse, error) {
	q := i.entClient.Task.Query()
	task, err := q.WithTestcaseSets().
		WithTestcases().
		Where(enttask.ID(int(req.Id))).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "the user(id: %d) is not found", req.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	testcases := make([]*testcases_repo.Testcase, 0, len(task.Edges.Testcases))
	for _, t := range task.Edges.Testcases {
		testcase, err := i.repository.DownloadTestcase(ctx, int(req.Id), t.Name)
		if err != nil {
			i.logger.Error("failed to download testcase", slog.Int("task_id", int(req.Id)), slog.String("testcase_name", t.Name))
			return nil, status.Error(codes.Internal, err.Error())
		}
		testcases = append(testcases, testcase)
	}

	return &backendv1.GetTaskResponse{
		Task: toPbTask(task),
		TestcaseSets: lo.Map(task.Edges.TestcaseSets, func(t *ent.TestcaseSet, _ int) *backendv1.TestcaseSet {
			return toPbTestcaseSet(t)
		}),
		Testcases: lo.Map(task.Edges.Testcases, func(t *ent.Testcase, i int) *backendv1.Testcase {
			return toPbTestcase(t, testcases[i].In, testcases[i].Out)
		}),
	}, nil
}

func toPbTask(t *ent.Task) *backendv1.Task {
	var updatedAt *timestamppb.Timestamp
	if t.UpdatedAt != nil {
		updatedAt = timestamppb.New(*t.UpdatedAt)
	}
	// TODO: set contest_ids
	return &backendv1.Task{
		Id:              int32(t.ID),
		Title:           t.Title,
		Statement:       t.Statement,
		ExecTimeLimit:   int32(t.ExecTimeLimit),
		ExecMemoryLimit: int32(t.ExecMemoryLimit),
		JudgeType:       toPbJudgeType(t),
		Difficulty:      backendv1.Difficulty(backendv1.Difficulty_value[t.Difficulty]),
		CreatedAt:       timestamppb.New(t.CreatedAt),
		UpdatedAt:       updatedAt,
	}
}

func toPbJudgeType(t *ent.Task) *judgev1.JudgeType {
	switch t.JudgeType {
	case enttask.JudgeTypeNormal:
		return &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Normal{
				Normal: &judgev1.JudgeTypeNormal{
					CaseInsensitive: t.CaseInsensitive,
				},
			},
		}
	case enttask.JudgeTypeEps:
		return &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Eps{
				Eps: &judgev1.JudgeTypeEPS{
					Ndigits: uint32(*t.Ndigits),
				},
			},
		}
	case enttask.JudgeTypeInteractive:
		return &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Interactive{
				Interactive: &judgev1.JudgeTypeInteractive{
					JudgeCodePath: *t.JudgeCodePath,
				},
			},
		}
	case enttask.JudgeTypeCustom:
		return &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Custom{
				Custom: &judgev1.JudgeTypeCustom{
					JudgeCodePath: *t.JudgeCodePath,
				},
			},
		}
	default:
		return nil
	}
}

func toPbTestcaseSet(t *ent.TestcaseSet) *backendv1.TestcaseSet {
	var updatedAt *timestamppb.Timestamp
	if t.UpdatedAt != nil {
		updatedAt = timestamppb.New(*t.UpdatedAt)
	}
	return &backendv1.TestcaseSet{
		Id:            int32(t.ID),
		Slug:          t.Name,
		Score:         int32(t.Score),
		IsSample:      t.IsSample,
		TestcaseSlugs: lo.Map(t.Edges.Testcases, func(t *ent.Testcase, _ int) string { return t.Name }),
		CreatedAt:     timestamppb.New(t.CreatedAt),
		UpdatedAt:     updatedAt,
	}
}

func toPbTestcase(t *ent.Testcase, input, output []byte) *backendv1.Testcase {
	var updatedAt *timestamppb.Timestamp
	if t.UpdatedAt != nil {
		updatedAt = timestamppb.New(*t.UpdatedAt)
	}
	return &backendv1.Testcase{
		Id:          int32(t.ID),
		Slug:        t.Name,
		Description: t.Description,
		Input:       string(input),
		Output:      string(output),
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   updatedAt,
	}
}
