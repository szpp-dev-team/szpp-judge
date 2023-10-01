package auth

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/szpp-dev-team/szpp-judge/backend/api/connect_server/interceptor"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	entuser "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	u_user "github.com/szpp-dev-team/szpp-judge/backend/usecases/user"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
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

func (i *Interactor) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := interceptor.GetClaimsFromContext(ctx).Username

	q := i.entClient.User.Query()
	user, err := q.Where(entuser.Username(username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err = u_user.VerifyPassword(user.HashedPassword, req.Password); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	secret := i.Secret
	tokens, err := GenerateTokens(ctx, i.entClient, []byte(secret), username)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		User:         u_user.ToPbUser(user),
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (i *Interactor) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	if err := killRefreshToken(ctx, i.entClient, req.RefreshToken); err != nil {
		return nil, err
	}
	return &pb.LogoutResponse{}, nil
}

func (i *Interactor) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	isTokenValid, err := verifyRefreshToken(ctx, i.entClient, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	if !isTokenValid {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid refresh token"))
	} else {
		secret := i.Secret
		username := interceptor.GetClaimsFromContext(ctx).Username
		accessToken, err := generateAccessToken([]byte(secret), username)
		if err != nil {
			return nil, err
		}
		return &pb.RefreshAccessTokenResponse{
			AccessToken: accessToken,
		}, nil
	}
}
