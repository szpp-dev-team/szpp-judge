package user

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	entuser "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"golang.org/x/exp/slog"
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

func (i *Interactor) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	q := i.entClient.User.Query()
	user, err := q.Where(entuser.Username(req.Username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetUserResponse{
		User: toPbUser(user),
	}, nil
}

func (i *Interactor) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	now := timejst.Now()
	q := i.entClient.User.Create().
		SetUsername(req.Username).
		SetHashedPassword(HashPassword(req.Password)).
		SetRole("USER").
		SetCreatedAt(now).
		SetUpdatedAt(now)
	user, err := q.Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateUserResponse{
		User: toPbUser(user),
	}, nil
}

func (i *Interactor) ExistsUsername(ctx context.Context, req *pb.ExistsUsernameRequest) (*pb.ExistsUsernameResponse, error) {
	q := i.entClient.User.Query()
	_, err := q.Where(entuser.Username(req.Username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &pb.ExistsUsernameResponse{
				Exists: false,
			}, nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ExistsUsernameResponse{
		Exists: true,
	}, nil
}

func (i *Interactor) ExistsEmail(ctx context.Context, req *pb.ExistsEmailRequest) (*pb.ExistsEmailResponse, error) {
	q := i.entClient.User.Query()
	_, err := q.Where(entuser.Email(req.Email)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &pb.ExistsEmailResponse{
				Exists: false,
			}, nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ExistsEmailResponse{
		Exists: true,
	}, nil
}

func toPbUser(t *ent.User) *backendv1.User {
	return &backendv1.User{
		Id:        int32(t.ID),
		Username:  t.Username,
		IsAdmin:   backendv1.Role_value[t.Role] == backendv1.Role_value["ADMIN"],
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}
