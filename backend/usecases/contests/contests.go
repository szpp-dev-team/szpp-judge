package contests

import (
	"context"
	"log/slog"

	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
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
			return nil, status.Error(codes.AlreadyExists, "the contest already exists")
		}
		return nil, status.Error(codes.Internal, "failed to create contest")
	}
	return &backendv1.CreateContestResponse{
		Contest: toPbContest(contest),
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
