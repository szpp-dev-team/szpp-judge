package tasks

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/samber/lo"
	"github.com/szpp-dev-team/szpp-judge/backend/api/connect_server/interceptor"
	"github.com/szpp-dev-team/szpp-judge/backend/core/entutil"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	ent_testcase "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcase"
	ent_testcaseset "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseset"
	ent_user "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	testcases_repo "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient     *ent.Client
	testcasesRepo testcases_repo.Repository
	logger        *slog.Logger
}

func NewInteractor(entClient *ent.Client, testcasesRepo testcases_repo.Repository) *Interactor {
	logger := slog.Default().With(slog.String("usecase", "tasks"))
	return &Interactor{entClient, testcasesRepo, logger}
}

func (i *Interactor) CreateTask(ctx context.Context, req *backendv1.CreateTaskRequest) (*backendv1.CreateTaskResponse, error) {
	claims := interceptor.GetClaimsFromContext(ctx)

	var task *ent.Task
	if err := entutil.WithTx(ctx, i.entClient, func(tx *ent.Tx) (err error) {
		userID, err := tx.User.Query().Where(ent_user.Username(claims.Username)).OnlyID(ctx)
		if err != nil {
			i.logger.Error("[!!!inconsistency!!!] failed to get user", slog.Any("error", err), slog.String("username", claims.Username))
			return connect.NewError(connect.CodeInternal, err)
		}

		q := tx.Task.Create().
			SetTitle(req.Task.Title).
			SetStatement(req.Task.Statement).
			SetDifficulty(req.Task.Difficulty.String()).
			SetExecTimeLimit(uint(req.Task.ExecTimeLimit)).
			SetExecMemoryLimit(uint(req.Task.ExecMemoryLimit)).
			SetCreatedAt(timejst.Now()).
			SetUserID(userID)
		switch ty := req.Task.JudgeType.JudgeType.(type) {
		case *judgev1.JudgeType_Normal:
			q.SetJudgeType(ent_task.JudgeTypeNormal)
			q.SetCaseInsensitive(*ty.Normal.CaseInsensitive)
		case *judgev1.JudgeType_Eps:
			q.SetJudgeType(ent_task.JudgeTypeEps)
			q.SetNdigits(uint(ty.Eps.Ndigits))
		case *judgev1.JudgeType_Interactive:
			q.SetJudgeType(ent_task.JudgeTypeInteractive)
			q.SetJudgeCodePath(ty.Interactive.JudgeCodePath)
		case *judgev1.JudgeType_Custom:
			q.SetJudgeType(ent_task.JudgeTypeCustom)
			q.SetJudgeCodePath(ty.Custom.JudgeCodePath)
		default:
			return fmt.Errorf("unrecognized JudgeType: %T", ty)
		}
		task, err = q.Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		i.logger.Error("failed to register task information", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	i.logger.Info("succeeded to register task information", slog.Int("task_id", task.ID))

	return &backendv1.CreateTaskResponse{
		Task: ToPbTask(task),
	}, nil
}

func (i *Interactor) GetTask(ctx context.Context, req *backendv1.GetTaskRequest) (*backendv1.GetTaskResponse, error) {
	task, err := i.entClient.Task.Get(ctx, int(req.TaskId))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("the task(id: %d) is not found", req.TaskId))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &backendv1.GetTaskResponse{
		Task: ToPbTask(task),
	}, nil
}

func (i *Interactor) UpdateTask(ctx context.Context, req *backendv1.UpdateTaskRequest) (*backendv1.UpdateTaskResponse, error) {
	claims := interceptor.GetClaimsFromContext(ctx)

	var task *ent.Task
	if err := entutil.WithTx(ctx, i.entClient, func(tx *ent.Tx) (err error) {
		task, err = tx.Task.Query().WithUser().Where(ent_task.ID(int(req.TaskId))).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return connect.NewError(connect.CodeNotFound, fmt.Errorf("the task(id: %d) is not found", req.TaskId))
			}
			i.logger.Error("failed to get task", slog.Any("error", err), slog.Int("task_id", int(req.TaskId)))
			return connect.NewError(connect.CodeInternal, err)
		}
		if claims.Username != task.Edges.User.Username {
			return connect.NewError(connect.CodePermissionDenied, fmt.Errorf("the task(id: %d) is not yours", req.TaskId))
		}

		q := tx.Task.UpdateOneID(int(req.TaskId)).
			SetTitle(req.Task.Title).
			SetStatement(req.Task.Statement).
			SetDifficulty(req.Task.Difficulty.String()).
			SetExecTimeLimit(uint(req.Task.ExecTimeLimit)).
			SetExecMemoryLimit(uint(req.Task.ExecMemoryLimit)).
			SetUpdatedAt(timejst.Now())
		switch ty := req.Task.JudgeType.JudgeType.(type) {
		case *judgev1.JudgeType_Normal:
			q.SetJudgeType(ent_task.JudgeTypeNormal)
			q.SetCaseInsensitive(*ty.Normal.CaseInsensitive)
		case *judgev1.JudgeType_Eps:
			q.SetJudgeType(ent_task.JudgeTypeEps)
			q.SetNdigits(uint(ty.Eps.Ndigits))
		case *judgev1.JudgeType_Interactive:
			q.SetJudgeType(ent_task.JudgeTypeInteractive)
			q.SetJudgeCodePath(ty.Interactive.JudgeCodePath)
		case *judgev1.JudgeType_Custom:
			q.SetJudgeType(ent_task.JudgeTypeCustom)
			q.SetJudgeCodePath(ty.Custom.JudgeCodePath)
		default:
			return fmt.Errorf("unrecognized JudgeType: %T", ty)
		}
		task, err = q.Save(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return connect.NewError(connect.CodeNotFound, fmt.Errorf("the task(id: %d) is not found", req.TaskId))
			}
			return connect.NewError(connect.CodeInternal, err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	i.logger.Info("succeeded to update task information", slog.Int("task_id", task.ID))

	return &backendv1.UpdateTaskResponse{
		Task: ToPbTask(task),
	}, nil
}

func (i *Interactor) GetTestcaseSets(ctx context.Context, req *backendv1.GetTestcaseSetsRequest) (*backendv1.GetTestcaseSetsResponse, error) {
	testcaseSets, err := i.entClient.TestcaseSet.Query().
		WithTestcases().
		Where(ent_testcaseset.HasTaskWith(ent_task.ID(int(req.TaskId)))).
		All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	testcaseByName := make(map[string]*backendv1.Testcase)
	for _, ts := range testcaseSets {
		for _, t := range ts.Edges.Testcases {
			if _, ok := testcaseByName[t.Name]; ok {
				continue
			}
			testcase, err := i.testcasesRepo.DownloadTestcase(ctx, int(req.TaskId), t.Name)
			if err != nil {
				i.logger.Error("failed to download testcase", slog.Int("task_id", int(req.TaskId)), slog.String("testcase_name", t.Name))
				return nil, connect.NewError(connect.CodeInternal, err)
			}
			testcaseByName[t.Name] = toPbTestcase(t, testcase.In, testcase.Out)
		}
	}

	return &backendv1.GetTestcaseSetsResponse{
		Testcases: lo.MapToSlice(testcaseByName, func(name string, t *backendv1.Testcase) *backendv1.Testcase {
			return t
		}),
		TestcaseSets: lo.Map(testcaseSets, func(ts *ent.TestcaseSet, _ int) *backendv1.TestcaseSet {
			return toPbTestcaseSet(ts)
		}),
	}, nil

}

func (i *Interactor) SyncTestcaseSets(ctx context.Context, req *backendv1.SyncTestcaseSetsRequest) (*backendv1.SyncTestcaseSetsResponse, error) {
	claims := interceptor.GetClaimsFromContext(ctx)

	var (
		testcaseSets []*ent.TestcaseSet
		testcases    []*ent.Testcase
	)
	if err := entutil.WithTx(ctx, i.entClient, func(tx *ent.Tx) (err error) {
		task, err := tx.Task.Query().WithUser().Where(ent_task.ID(int(req.TaskId))).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return connect.NewError(connect.CodeNotFound, fmt.Errorf("the task(id: %d) is not found", req.TaskId))
			}
			i.logger.Error("failed to get task", slog.Any("error", err), slog.Int("task_id", int(req.TaskId)))
			return connect.NewError(connect.CodeInternal, err)
		}
		if claims.Username != task.Edges.User.Username {
			return connect.NewError(connect.CodePermissionDenied, fmt.Errorf("the task(id: %d) is not yours", req.TaskId))
		}

		// Testcase
		testcases, err = syncTestcases(ctx, tx, int(req.TaskId), req.Testcases)
		if err != nil {
			i.logger.Error("failed to upsert testcases", slog.Any("error", err))
			return err
		}

		// TestcaseSet
		testcaseSets, err = syncTestcaseSets(ctx, tx, int(req.TaskId), req.TestcaseSets, testcases)
		if err != nil {
			i.logger.Error("failed to upsert testcase_sets", slog.Any("error", err))
			return err
		}

		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	for _, testcase := range req.Testcases {
		i.logger.Info("uploading testcase", slog.String("testcase", testcase.Slug), slog.Int("task_id", int(req.TaskId)))
		if err := i.testcasesRepo.UpsertTestcase(ctx, int(req.TaskId), &testcases_repo.Testcase{
			Name: testcase.Slug,
			In:   []byte(testcase.Input),
			Out:  []byte(testcase.Output),
		}); err != nil {
			i.logger.Error("failed to upload a testcase", slog.Any("error", err))
			return nil, err
		}
	}
	return &backendv1.SyncTestcaseSetsResponse{
		TestcaseSets: lo.Map(testcaseSets, func(ts *ent.TestcaseSet, _ int) *backendv1.TestcaseSet {
			return toPbTestcaseSet(ts)
		}),
		Testcases: lo.Map(testcases, func(t *ent.Testcase, _ int) *backendv1.Testcase {
			return toPbTestcase(t, nil, nil)
		}),
	}, nil
}

func syncTestcases(ctx context.Context, tx *ent.Tx, taskID int, list []*backendv1.MutationTestcase) ([]*ent.Testcase, error) {
	var (
		oldTestcaseIDs []int
		err            error
	)

	oldTestcaseIDs, err = tx.Testcase.Query().Where(ent_testcase.HasTaskWith(ent_task.ID(taskID))).IDs(ctx)
	if err != nil {
		return nil, err
	}

	newTestcaseIDs := make([]int, 0, len(list))
	now := timejst.Now()
	for _, mt := range list {
		testcaseID, err := tx.Testcase.Create().
			SetName(mt.Slug).
			SetNillableDescription(mt.Description).
			SetTaskID(taskID).
			SetCreatedAt(now).
			OnConflict().
			UpdateNewValues().
			Update(func(tu *ent.TestcaseUpsert) {
				tu.SetUpdatedAt(now)
			}).ID(ctx)
		if err != nil {
			return nil, err
		}
		newTestcaseIDs = append(newTestcaseIDs, testcaseID)
	}

	// 使われなくなった Testcase を削除
	_, removeIDs := diffSlices(oldTestcaseIDs, newTestcaseIDs)
	if _, err := tx.Testcase.Delete().Where(ent_testcase.IDIn(removeIDs...)).Exec(ctx); err != nil {
		return nil, err
	}

	return tx.Testcase.Query().Where(ent_testcase.IDIn(newTestcaseIDs...)).All(ctx)
}

func syncTestcaseSets(
	ctx context.Context,
	tx *ent.Tx,
	taskID int,
	testcaseSetList []*backendv1.MutationTestcaseSet,
	testcaseList []*ent.Testcase,
) ([]*ent.TestcaseSet, error) {
	now := timejst.Now()

	var (
		testcaseSetAllIDs []int
		err               error
	)
	testcaseSetAllIDs, err = tx.TestcaseSet.Query().Where(ent_testcaseset.HasTaskWith(ent_task.ID(taskID))).IDs(ctx)
	if err != nil {
		return nil, err
	}
	testcaseByName := lo.Associate(testcaseList, func(t *ent.Testcase) (string, *ent.Testcase) { return t.Name, t })

	testcaseSetIDs := make([]int, 0, len(testcaseSetList))
	for _, testcaseSet := range testcaseSetList {
		oldTestcaseIDs, err := tx.Testcase.Query().Where(ent_testcase.HasTestcaseSetsWith(ent_testcaseset.Name(testcaseSet.Slug))).IDs(ctx)
		if err != nil {
			return nil, err
		}

		newTestcaseIDs := lo.Map(testcaseSet.TestcaseSlugs, func(s string, _ int) int { return testcaseByName[s].ID })
		testcaseSetID, err := tx.TestcaseSet.Create().
			SetName(testcaseSet.Slug).
			SetScoreRatio(int(testcaseSet.ScoreRatio)).
			SetIsSample(testcaseSet.IsSample).
			SetTaskID(taskID).
			AddTestcaseIDs(newTestcaseIDs...).
			SetCreatedAt(now).
			OnConflict().
			UpdateNewValues().
			Update(func(tsu *ent.TestcaseSetUpsert) {
				tsu.SetUpdatedAt(now)
			}).ID(ctx)
		if err != nil {
			return nil, err
		}

		// 使われなくなった Testcase を TestcaseSet から削除
		_, removeIDs := diffSlices(oldTestcaseIDs, newTestcaseIDs)
		if err := tx.TestcaseSet.UpdateOneID(testcaseSetID).RemoveTestcaseIDs(removeIDs...).Exec(ctx); err != nil {
			return nil, err
		}

		testcaseSetIDs = append(testcaseSetIDs, testcaseSetID)
	}

	// 使われなくなった TestcaseSet を削除
	_, removeIDs := diffSlices(testcaseSetAllIDs, testcaseSetIDs)
	if _, err := tx.Testcase.Delete().Where(ent_testcase.IDIn(removeIDs...)).Exec(ctx); err != nil {
		return nil, err
	}

	return tx.TestcaseSet.Query().Where(ent_testcaseset.IDIn(testcaseSetIDs...)).All(ctx)
}

func diffSlices[T comparable](prevList, nextList []T) (addList, removeList []T) {
	prevSet := lo.Associate(prevList, func(t T) (T, struct{}) { return t, struct{}{} })
	nextSet := lo.Associate(nextList, func(t T) (T, struct{}) { return t, struct{}{} })
	for prev := range prevSet {
		if _, ok := nextSet[prev]; !ok {
			removeList = append(removeList, prev)
		}
	}
	for next := range nextSet {
		addList = append(addList, next)
	}
	return addList, removeList
}

func ToPbTask(t *ent.Task) *backendv1.Task {
	var updatedAt *timestamppb.Timestamp
	if t.UpdatedAt != nil {
		updatedAt = timestamppb.New(*t.UpdatedAt)
	}
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
	case ent_task.JudgeTypeNormal:
		return &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Normal{
				Normal: &judgev1.JudgeTypeNormal{
					CaseInsensitive: t.CaseInsensitive,
				},
			},
		}
	case ent_task.JudgeTypeEps:
		return &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Eps{
				Eps: &judgev1.JudgeTypeEPS{
					Ndigits: uint32(*t.Ndigits),
				},
			},
		}
	case ent_task.JudgeTypeInteractive:
		return &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Interactive{
				Interactive: &judgev1.JudgeTypeInteractive{
					JudgeCodePath: *t.JudgeCodePath,
				},
			},
		}
	case ent_task.JudgeTypeCustom:
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
		ScoreRatio:    int32(t.ScoreRatio),
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
