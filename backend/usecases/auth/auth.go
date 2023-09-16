package auth

import (
	"context"

	intercepter "github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server/intercepter"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	entuser "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
	Secret    string
}

func NewInteractor(entClient *ent.Client, secret string) *Interactor {
	logger := slog.Default().With(slog.String("usecase", "auth"))
	return &Interactor{entClient, logger, secret}
}

func (i *Interactor) Login(ctx context.Context, req *pb.LoginRequest, secret string) (*pb.LoginResponse, error) {
	username := intercepter.GetClaimsFromContext(ctx).Username

	q := i.entClient.User.Query()
	user, err := q.Where(entuser.Username(username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = VerifyPassword(user.HashedPassword, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	tokens, err := GenerateTokens(ctx, i.entClient, []byte(secret), username)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		User:         toPbUser(user),
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (i *Interactor) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := KillRefreshToken(ctx, i.entClient, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutResponse{}, nil
}

func (i *Interactor) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenRequest, secret string) (*pb.RefreshAccessTokenResponse, error) {
	isTokenValid, err := VerifyRefreshToken(ctx, i.entClient, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	if !isTokenValid {
		return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
	} else {
		username := intercepter.GetClaimsFromContext(ctx).Username
		accessToken, _ := GenerateAccessToken([]byte(secret), username)
		return &pb.RefreshAccessTokenResponse{
			AccessToken: accessToken,
		}, nil
	}
}

func toPbUser(t *ent.User) *pb.User {
	return &pb.User{
		Id:        int32(t.ID),
		Username:  t.Username,
		IsAdmin:   pb.Role_value[t.Role] == pb.Role_value["ADMIN"],
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}
