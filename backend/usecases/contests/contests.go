package contests

import (
	"context"
	"log/slog"

	"github.com/samber/lo"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_contest "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client) *Interactor {
	logger := slog.Default().With(slog.String("usecase", "tasks"))
	return &Interactor{entClient, logger}
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
			return nil, status.Error(codes.AlreadyExists, "the contest has already existed")
		}
		return nil, status.Error(codes.Internal, "failed to create contest")
	}
	return &backendv1.CreateContestResponse{
		Contest: toPbContest(contest),
	}, nil
}

func (i *Interactor) GetContest(ctx context.Context, req *backendv1.GetContestRequest) (*backendv1.GetContestResponse, error) {
	contest, err := i.entClient.Contest.Query().Where(ent_contest.Slug(req.Slug)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the contest was not found")
		}
		return nil, status.Error(codes.Internal, "failed to get contest")
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
			return nil, status.Error(codes.AlreadyExists, "the slug has been already used")
		}
		return nil, status.Error(codes.Internal, "failed to update contest")
	}
	return &backendv1.UpdateContestResponse{
		Contest: toPbContest(contest),
	}, nil
}

func (i *Interactor) ListContests(ctx context.Context, req *backendv1.ListContestsRequest) (*backendv1.ListContestsResponse, error) {
	contests, err := i.entClient.Contest.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the contest was not found")
		}
		return nil, status.Error(codes.Internal, "failed to get contest")
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
			return nil, status.Error(codes.NotFound, "the contest was not found")
		}
		return nil, status.Error(codes.Internal, "failed to get contest")
	}
	return &backendv1.ListContestTasksResponse{
		Tasks: lo.Map(contest.Edges.ContestTask, func(t *ent.ContestTask, _ int) *backendv1.ContestTask {
			return toPbContestTask(t)
		}),
	}, nil
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
		Id:              int32(ct.ID),
		Title:           ct.Edges.Task.Title,
		ExecTimeLimit:   int32(ct.Edges.Task.ExecTimeLimit),
		ExecMemoryLimit: int32(ct.Edges.Task.ExecMemoryLimit),
		Difficulty:      backendv1.Difficulty(backendv1.Difficulty_value[ct.Edges.Task.Difficulty]),
		Score:           int32(ct.Score),
	}
}
