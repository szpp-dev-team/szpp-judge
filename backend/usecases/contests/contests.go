package contests

import (
	"context"
	"log/slog"
	"time"

	"github.com/samber/lo"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_contest "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	ent_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/tasks"
	ent_user "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	intercepter "github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server/intercepter"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/clarification"
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

func (i *Interactor) GetContestTask(ctx context.Context, req *backendv1.GetContestTaskRequest) (*backendv1.GetContestTaskResponse, error) {
	task, err := i.entClient.Task.Query().
		Where(
			ent_task.HasContestsWith(ent_contest.Slug(req.ContestSlug)),
			ent_task.ID(int(req.TaskId)),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "the contest was not found")
		}
		return nil, status.Error(codes.Internal, "failed to get contest")
	}
	return &backendv1.GetContestTaskResponse{
		Task: tasks.ToPbTask(task),
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

func (i *Interactor) CreateClarification(ctx context.Context, req *backendv1.CreateClarificationRequest) (*backendv1.CreateClarificationResponse, error) {
	contest, err := i.entClient.Contest.Query().Where(ent_contest.Slug(req.ContestSlug)).Only(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get contest")
	}

	task, err := i.entClient.Task.Query().Where(ent_task.ID(int(*req.TaskId))).Only(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	username := intercepter.GetClaimsFromContext(ctx).Username
	user, err := i.entClient.User.Query().Where(ent_user.Username(username)).Only(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}
	

	now := time.Now()
	clarification, err := i.entClient.Clarification.Create().
		SetContent(req.Content).
		SetIsPublic(false).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		AddContest(contest).
		AddTask(task).
		AddUser(user).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create clarification")
	}
	return &backendv1.CreateClarificationResponse{
		Clarification: &backendv1.Clarification {
			Id: int32(clarification.ID),
			Content: clarification.Content,
			IsPublic: clarification.IsPublic,
			IsMine: true,
			CreatedAt: timestamppb.New(clarification.CreatedAt),
			UpdatedAt: timestamppb.New(clarification.UpdatedAt),
		},
	}, nil	
}

func (i *Interactor) ListClarifications(ctx context.Context, req *backendv1.ListClarificationsRequest) (*backendv1.ListClarificationsResponse, error) {
	// Eager-load the related Users using WithUsers()
	clarifications, err := i.entClient.Clarification.Query().
		WithUser().
		Where(clarification.HasContestWith(ent_contest.Slug(req.ContestSlug))).
		All(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get clarifications")
	}

	username := intercepter.GetClaimsFromContext(ctx).Username
	user, err := i.entClient.User.Query().Where(ent_user.Username(username)).Only(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	var pbClarifications []*backendv1.Clarification
	for _, clarification := range clarifications {
		isMine := false
		for _, u := range clarification.Edges.User {  // Adjusted this line (assuming the edge name is Users)
			if u.ID == user.ID {
				isMine = true
				break
			}
		}

		if isMine || clarification.IsPublic {
			pbClarifications = append(pbClarifications, &backendv1.Clarification{
				Id:       int32(clarification.ID),
				Content:  clarification.Content,
				IsPublic: clarification.IsPublic,
				CreatedAt: timestamppb.New(clarification.CreatedAt),
				UpdatedAt: timestamppb.New(clarification.UpdatedAt),
			})
		}
	}
	return &backendv1.ListClarificationsResponse{
		Clarifications: pbClarifications,
	}, nil
}

func (i *Interactor) DeleteClarification(ctx context.Context, req *backendv1.DeleteClarificationRequest) (*backendv1.DeleteClarificationResponse, error) {
    // 指定されたIDでClarificationエンティティを検索
    clarification, err := i.entClient.Clarification.Query().Where(clarification.ID(int(req.Id))).Only(ctx)
    if err != nil {
        return nil, status.Error(codes.Internal, "failed to get clarification")
    }

	//自分で作成したClarificationのみ削除可能
	username := intercepter.GetClaimsFromContext(ctx).Username
	user, err := i.entClient.User.Query().Where(ent_user.Username(username)).Only(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}
	if user.ID != clarification.Edges.User[0].ID {
		return nil, status.Error(codes.Internal, "failed to delete clarification")
	}

    // 該当するClarificationエンティティを削除
    err = i.entClient.Clarification.DeleteOne(clarification).Exec(ctx)
    if err != nil {
        return nil, status.Error(codes.Internal, "failed to delete clarification")
    }

    // 成功のレスポンスを返す
    return &backendv1.DeleteClarificationResponse{}, nil
}

func (i *Interactor) CreateAnswer(ctx context.Context, req *backendv1.CreateAnswerRequest) (*backendv1.CreateAnswerResponse, error) {
	// 指定されたIDでClarificationエンティティを検索
	clarification, err := i.entClient.Clarification.Query().Where(clarification.ID(int(req.ClarificationId))).Only(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get clarification")
	}

	username := intercepter.GetClaimsFromContext(ctx).Username
	user, err := i.entClient.User.Query().Where(ent_user.Username(username)).Only(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	//userののRoleがadminがWRITERではない場合はエラー
	if user.Role != "admin" && user.Role != "WRITER" {
		return nil, status.Error(codes.Internal, "failed to create clarification")
	}

	now := time.Now()
	// 回答内容を設定
	clarification, err = clarification.Update().
		SetIsPublic(req.IsPublic).
		SetAnswerContent(req.Content).
		SetUpdatedAt(now).
		SetAnswerCreatedAt(now).
		SetAnswerUpdatedAt(now).
		AddAnswerUser(user). 
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update clarification")
	}

    // nil チェック
    var answerContent string
    if clarification.AnswerContent != nil {
        answerContent = *clarification.AnswerContent
    }
    var answerCreatedAt, answerUpdatedAt *timestamppb.Timestamp
    if clarification.AnswerCreatedAt != nil {
        answerCreatedAt = timestamppb.New(*clarification.AnswerCreatedAt)
    }
    if clarification.AnswerUpdatedAt != nil {
        answerUpdatedAt = timestamppb.New(*clarification.AnswerUpdatedAt)
    }

    // Answerに関する情報とIsPublicを返す
    return &backendv1.CreateAnswerResponse{
        Answer: &backendv1.Clarification_Answer {
            Id: int32(clarification.ID),
            Content: answerContent,
            IsMine: true,
            CreatedAt: answerCreatedAt,
            UpdatedAt: answerUpdatedAt,
        }, // ここでIsPublicを返す
		IsPublic: clarification.IsPublic,
	}, nil
}

func (i *Interactor) UpdateAnswer (ctx context.Context, req *backendv1.UpdateAnswerRequest) (*backendv1.UpdateAnswerResponse, error) {
	// 指定されたIDでClarificationエンティティを検索
	clarification, err := i.entClient.Clarification.Query().Where(clarification.ID(int(req.AnswerId))).Only(ctx)//TODO: AnswerIDはないので、Clarification IDにする必要がある。
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get clarification")
	}

	now := time.Now()
	// 回答内容を設定
	clarification, err = clarification.Update().
		SetAnswerContent(req.Content).
		SetUpdatedAt(now).
		SetAnswerUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update clarification")
	}

	// nil チェック
	var answerContent string
	if clarification.AnswerContent != nil {
		answerContent = *clarification.AnswerContent
	}
	var answerUpdatedAt *timestamppb.Timestamp
	if clarification.AnswerUpdatedAt != nil {
		answerUpdatedAt = timestamppb.New(*clarification.AnswerUpdatedAt)
	}

	return &backendv1.UpdateAnswerResponse{
		Answer: &backendv1.Clarification_Answer {
			Id: int32(clarification.ID),
			Content: answerContent,
			IsMine: true,
			UpdatedAt: answerUpdatedAt,
		},
	}, nil
}

func (i *Interactor) DeleteAsnwer (ctx context.Context, req *backendv1.DeleteAnswerRequest) (*backendv1.DeleteAnswerResponse, error) {
	// 指定されたIDでClarificationエンティティを検索
	clarification, err := i.entClient.Clarification.Query().Where(clarification.ID(int(req.AnswerId))).Only(ctx)//TODO: AnswerIDはないので、Clarification IDにする必要がある。
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get clarification")
	}

	// 回答内容を設定
	clarification, err = clarification.Update().
		ClearAnswerContent().
		ClearAnswerCreatedAt().
		ClearAnswerUpdatedAt().
		ClearAnswerUser().
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete clarification")
	}

    return &backendv1.DeleteAnswerResponse{}, nil
}
