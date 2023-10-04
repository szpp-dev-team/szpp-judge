package user

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	entuser "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client) *Interactor {
	logger := slog.Default().With(slog.String("usecase", "user"))
	return &Interactor{entClient, logger}
}

func (i *Interactor) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	q := i.entClient.User.Query()
	user, err := q.Where(entuser.Username(req.Username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &pb.GetUserResponse{
		User: ToPbUser(user),
	}, nil
}

func (i *Interactor) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	now := timejst.Now()
	q := i.entClient.User.Create().
		SetUsername(req.Username).
		SetHashedPassword(HashPassword(req.Password)).
		SetEmail(req.Email).
		SetRole(pb.Role_name[int32(pb.Role_USER)]).
		SetCreatedAt(now).
		SetUpdatedAt(now)
	user, err := q.Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("username or email already exists"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &pb.CreateUserResponse{
		User: ToPbUser(user),
	}, nil
}

func (i *Interactor) ExistsUsername(ctx context.Context, req *pb.ExistsUsernameRequest) (*pb.ExistsUsernameResponse, error) {
	q := i.entClient.User.Query()
	exist, err := q.Where(entuser.Username(req.Username)).Exist(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &pb.ExistsUsernameResponse{
		Exists: exist,
	}, nil
}

func (i *Interactor) ExistsEmail(ctx context.Context, req *pb.ExistsEmailRequest) (*pb.ExistsEmailResponse, error) {
	q := i.entClient.User.Query()
	exist, err := q.Where(entuser.Email(req.Email)).Exist(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &pb.ExistsEmailResponse{
		Exists: exist,
	}, nil
}

func ToPbUser(t *ent.User) *pb.User {
	return &pb.User{
		Id:        int32(t.ID),
		Username:  t.Username,
		IsAdmin:   pb.Role_value[t.Role] == int32(pb.Role_ADMIN),
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}
