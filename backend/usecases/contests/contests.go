package contests

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/samber/lo"
	"github.com/szpp-dev-team/szpp-judge/backend/api/connect_server/interceptor"
	"github.com/szpp-dev-team/szpp-judge/backend/core/entutil"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_contest "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	ent_contesttask "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contesttask"
	ent_contestuser "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contestuser"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	ent_submit "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	ent_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	ent_testcaseset "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseset"
	ent_user "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	testcases_repo "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/tasks"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient     *ent.Client
	logger        *slog.Logger
	testcasesRepo testcases_repo.Repository
}

func NewInteractor(entClient *ent.Client, testcasesRepo testcases_repo.Repository) *Interactor {
	logger := slog.Default().With(slog.String("usecase", "contests"))
	return &Interactor{entClient, logger, testcasesRepo}
}

func (i *Interactor) CreateContest(ctx context.Context, req *backendv1.CreateContestRequest) (*backendv1.CreateContestResponse, error) {
	contest, err := i.entClient.Contest.Create().
		SetName(req.Contest.Name).
		SetSlug(req.Contest.Slug).
		SetDescription(req.Contest.Description).
		SetPenaltySeconds(int(req.Contest.PenaltySeconds)).
		SetContestType(req.Contest.ContestType.String()).
		SetIsPublic(req.Contest.IsPublic).
		SetStartAt(req.Contest.StartAt.AsTime()).
		SetEndAt(req.Contest.EndAt.AsTime()).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("the contest has already existed"))
		}
		i.logger.Error("failed to create contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to create contest"))
	}
	return &backendv1.CreateContestResponse{
		Contest: toPbContest(contest),
	}, nil
}

func (i *Interactor) GetContest(ctx context.Context, req *backendv1.GetContestRequest) (*backendv1.GetContestResponse, error) {
	contest, err := i.entClient.Contest.Query().Where(ent_contest.Slug(req.Slug)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the contest was not found"))
		}
		i.logger.Error("failed to get contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to get contest"))
	}
	return &backendv1.GetContestResponse{
		Contest: toPbContest(contest),
	}, nil
}

func (i *Interactor) UpdateContest(ctx context.Context, req *backendv1.UpdateContestRequest) (*backendv1.UpdateContestResponse, error) {
	contest, err := i.entClient.Contest.UpdateOneID(int(req.Id)).
		SetName(req.Contest.Name).
		SetSlug(req.Contest.Slug).
		SetDescription(req.Contest.Description).
		SetPenaltySeconds(int(req.Contest.PenaltySeconds)).
		SetContestType(req.Contest.ContestType.String()).
		SetIsPublic(req.Contest.IsPublic).
		SetStartAt(req.Contest.StartAt.AsTime()).
		SetEndAt(req.Contest.EndAt.AsTime()).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("the slug has been already used"))
		}
		i.logger.Error("failed to update contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to update contest"))
	}
	return &backendv1.UpdateContestResponse{
		Contest: toPbContest(contest),
	}, nil
}

func (i *Interactor) ListContests(ctx context.Context, req *backendv1.ListContestsRequest) (*backendv1.ListContestsResponse, error) {
	contests, err := i.entClient.Contest.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the contest was not found"))
		}
		i.logger.Error("failed to get contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to get contest"))
	}
	return &backendv1.ListContestsResponse{
		Contests: lo.Map(contests, func(c *ent.Contest, _ int) *backendv1.Contest {
			return toPbContest(c)
		}),
	}, nil
}

func (i *Interactor) ListContestTasks(ctx context.Context, req *backendv1.ListContestTasksRequest) (*backendv1.ListContestTasksResponse, error) {
	contest, err := i.entClient.Contest.Query().
		WithContestTask(func(ctq *ent.ContestTaskQuery) {
			ctq.WithTask()
		}).
		Where(ent_contest.Slug(req.ContestSlug)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the contest was not found"))
		}
		i.logger.Error("failed to get contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to get contest"))
	}
	return &backendv1.ListContestTasksResponse{
		Tasks: lo.Map(contest.Edges.ContestTask, func(t *ent.ContestTask, _ int) *backendv1.ContestTask {
			return toPbContestTask(t)
		}),
	}, nil
}

func (i *Interactor) GetContestTask(ctx context.Context, req *backendv1.GetContestTaskRequest) (*backendv1.GetContestTaskResponse, error) {
	task, err := i.entClient.Task.Query().
		WithTestcaseSets(func(tsq *ent.TestcaseSetQuery) {
			tsq.Where(ent_testcaseset.IsSample(true))
			tsq.WithTestcases(func(tq *ent.TestcaseQuery) {
				tq.WithTask()
			})
		}).
		Where(
			ent_task.HasContestsWith(ent_contest.Slug(req.ContestSlug)),
			ent_task.ID(int(req.TaskId)),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the contest was not found"))
		}
		i.logger.Error("failed to get contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to get contest"))
	}

	i.logger.Info("testcase sets", slog.Any("testcase sets", task.Edges.TestcaseSets))
	for _, tcs := range task.Edges.TestcaseSets {
		i.logger.Info("testcases", slog.Any("testcases", tcs.Edges.Testcases))
	}

	testcases := make([]*testcases_repo.Testcase, 0, len(task.Edges.Testcases))
	for _, tcs := range task.Edges.TestcaseSets {
		for _, tc := range tcs.Edges.Testcases {
			i.logger.Info("task id", slog.Any("task id", tc.Edges.Task.ID))
			i.logger.Info("testcase name", slog.Any("testcase name", tc.Name))
			testcase, err := i.testcasesRepo.DownloadTestcase(ctx, tc.Edges.Task.ID, tc.Name)
			if err != nil {
				i.logger.Error("failed to download testcase", slog.Any("error", err))
				return nil, connect.NewError(connect.CodeInternal, errors.New("failed to download testcase"))
			}
			testcases = append(testcases, testcase)
		}
	}

	return &backendv1.GetContestTaskResponse{
		Task: tasks.ToPbTask(task),
		Samples: lo.Map(task.Edges.Testcases, func(tc *ent.Testcase, i int) *backendv1.Testcase {
			return tasks.ToPbTestcase(tc, testcases[i].In, testcases[i].Out)
		}),
	}, nil
}

func (i *Interactor) SyncContestTasks(ctx context.Context, req *backendv1.SyncContestTasksRequest) (*backendv1.SyncContestTasksResponse, error) {
	contestTaskIDs := make([]int, 0, len(req.Tasks))
	if err := entutil.WithTx(ctx, i.entClient, func(tx *ent.Tx) error {
		contest, err := i.entClient.Contest.Query().
			Where(ent_contest.Slug(req.ContestSlug)).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return connect.NewError(connect.CodeNotFound, err)
			}
			i.logger.Error("failed to get contest", slog.Any("error", err), slog.String("slug", req.ContestSlug))
			return connect.NewError(connect.CodeInternal, err)
		}
		if _, err := i.entClient.Contest.UpdateOne(contest).ClearTasks().Save(ctx); err != nil {
			i.logger.Error("failed to clear tasks which are related with the contest", slog.Any("error", err))
			return connect.NewError(connect.CodeInternal, err)
		}

		for order, task := range req.Tasks {
			contestTask, err := tx.ContestTask.Create().
				SetContest(contest).
				SetTaskID(int(task.Id)).
				SetScore(int(task.Score)).
				SetOrder(order + 1).
				Save(ctx)
			if err != nil {
				i.logger.Error("failed to create contest task", slog.Any("error", err))
				return connect.NewError(connect.CodeInternal, errors.New("failed to create contest task"))
			}
			contestTaskIDs = append(contestTaskIDs, contestTask.ID)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	predicates := make([]predicate.ContestTask, 0, len(contestTaskIDs))
	for _, id := range contestTaskIDs {
		predicates = append(predicates, ent_contesttask.ID(id))
	}
	contestTasks, err := i.entClient.ContestTask.Query().
		WithTask().
		Where(ent_contesttask.Or(predicates...)).
		All(ctx)
	if err != nil {
		i.logger.Error("failed to get contest tasks", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to get contest tasks"))
	}

	slices.SortFunc(contestTasks, func(a, b *ent.ContestTask) int {
		return a.Order - b.Order
	})
	return &backendv1.SyncContestTasksResponse{
		Tasks: lo.Map(contestTasks, func(ct *ent.ContestTask, _ int) *backendv1.Task {
			return tasks.ToPbTask(ct.Edges.Task)
		}),
	}, nil
}

func (i *Interactor) GetMySubmissionStatuses(ctx context.Context, req *backendv1.GetMySubmissionStatusesRequest) (*backendv1.GetMySubmissionStatusesResponse, error) {
	claims := interceptor.GetClaimsFromContext(ctx)

	contest, err := i.entClient.Contest.Query().
		WithSubmits(func(sq *ent.SubmitQuery) {
			sq.WithTask().
				Where(ent_submit.HasUserWith(ent_user.Username(claims.Username)))
		}).
		Where(ent_contest.Slug(req.ContestSlug)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the contest was not found"))
		}
		i.logger.Error("failed to get contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to get contest"))
	}

	submissionStatusByTaskID := make(map[int]*backendv1.SubmissionStatus)
	for _, submit := range contest.Edges.Submits {
		if _, ok := submissionStatusByTaskID[submit.Edges.Task.ID]; !ok {
			submissionStatusByTaskID[submit.Edges.Task.ID] = &backendv1.SubmissionStatus{
				TaskId: int32(submit.Edges.Task.ID),
				Score:  lo.ToPtr(int32(submit.Score)),
			}
		}
		if int32(submit.Score) > *submissionStatusByTaskID[submit.Edges.Task.ID].Score {
			submissionStatusByTaskID[submit.Edges.Task.ID].Score = lo.ToPtr(int32(submit.Score))
		}
	}

	submissionStatuses := make([]*backendv1.SubmissionStatus, 0, len(submissionStatusByTaskID))
	for _, submissionStatus := range submissionStatusByTaskID {
		submissionStatuses = append(submissionStatuses, submissionStatus)
	}

	return &backendv1.GetMySubmissionStatusesResponse{
		SubmissionStatuses: submissionStatuses,
	}, nil
}

func (i *Interactor) GetMyRegistrationStatus(ctx context.Context, req *backendv1.GetMyRegistrationStatusRequest) (*backendv1.GetMyRegistrationStatusResponse, error) {
	claims := interceptor.GetClaimsFromContext(ctx)

	contest, err := i.entClient.Contest.Query().
		WithContestUser(func(cuq *ent.ContestUserQuery) {
			cuq.Where(ent_contestuser.UserID(claims.UserID))
		}).
		Where(ent_contest.Slug(req.ContestSlug)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "contest not found")
		}
		i.logger.Error("failed to get contest", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &backendv1.GetMyRegistrationStatusResponse{
		Registered: contest.Edges.ContestUser != nil,
	}, nil
}

func (i *Interactor) RegisterMe(ctx context.Context, req *backendv1.RegisterMeRequest) (*backendv1.RegisterMeResponse, error) {
	contest, err := i.entClient.Contest.Query().Where(ent_contest.Slug(req.ContestSlug)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("the contest was not found"))
		}
		i.logger.Error("failed to get contest", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to get contest"))
	}

	if _, err := i.entClient.ContestUser.Create().
		// SetRole().
		SetContest(contest).
		// SetUserID().
		Save(ctx); err != nil {
		if ent.IsConstraintError(err) {
			return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("the user has already registered"))
		}
		i.logger.Error("failed to register user", slog.Any("error", err))
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to register user"))
	}

	return &backendv1.RegisterMeResponse{}, nil
}

func toPbContest(contest *ent.Contest) *backendv1.Contest {
	return &backendv1.Contest{
		Id:             int32(contest.ID),
		Name:           contest.Name,
		Slug:           contest.Slug,
		Description:    contest.Description,
		PenaltySeconds: int32(contest.PenaltySeconds),
		ContestType:    backendv1.ContestType(backendv1.ContestType_value[contest.ContestType]),
		StartAt:        timestamppb.New(contest.StartAt),
		EndAt:          timestamppb.New(contest.EndAt),
	}
}

func toPbContestTask(ct *ent.ContestTask) *backendv1.ContestTask {
	return &backendv1.ContestTask{
		Id:              int32(ct.Edges.Task.ID),
		Title:           ct.Edges.Task.Title,
		ExecTimeLimit:   int32(ct.Edges.Task.ExecTimeLimit),
		ExecMemoryLimit: int32(ct.Edges.Task.ExecMemoryLimit),
		Difficulty:      backendv1.Difficulty(backendv1.Difficulty_value[ct.Edges.Task.Difficulty]),
		Score:           int32(ct.Score),
	}
}
