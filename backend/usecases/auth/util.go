package auth

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	enttoken "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/refreshtoken"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func GenerateTokens(ctx context.Context, entClient *ent.Client, secret []byte, username string) (Tokens, error) {
	accessToken, err := generateAccessToken(secret, username)
	if err != nil {
		return Tokens{}, connect.NewError(connect.CodeInternal, err)
	}
	refreshToken, err := generateRefreshToken(ctx, entClient)
	if err != nil {
		return Tokens{}, connect.NewError(connect.CodeInternal, err)
	}

	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateAccessToken(secret []byte, username string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timejst.Now().Add(time.Hour)),
		},
	})
	return accessToken.SignedString(secret)
}

func generateRefreshToken(ctx context.Context, entClient *ent.Client) (string, error) {
	uuidObj, _ := uuid.NewRandom()
	token := uuidObj.String()
	expiresAt := timejst.Now().Add(time.Hour * 24 * 30)

	q := entClient.RefreshToken.Create().
		SetToken(token).
		SetExpiresAt(expiresAt).
		SetIsDead(false)

	if _, err := q.Save(ctx); err != nil {
		return "", connect.NewError(connect.CodeInternal, err)
	}

	return token, nil
}

func killRefreshToken(ctx context.Context, entClient *ent.Client, token string) error {
	q := entClient.RefreshToken.Update().
		Where(enttoken.Token(token)).
		SetIsDead(true)
	_, err := q.Save(ctx)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	return nil
}

func verifyRefreshToken(ctx context.Context, entClient *ent.Client, token string) (bool, error) {
	now := timejst.Now()
	// isdeadではなく、有効期限が切れていないトークンであるかどうかを確認する
	q := entClient.RefreshToken.Query().
		Where(enttoken.Token(token)).
		Where(enttoken.IsDead(false)).
		Where(enttoken.ExpiresAtGT(now))
	exist, err := q.Exist(ctx)
	if err != nil {
		return false, connect.NewError(connect.CodeInternal, err)
	}
	return exist, nil
}
