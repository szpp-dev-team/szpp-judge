package auth

import (
	"context"
	"crypto/rand"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	enttoken "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/refreshtoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return Tokens{}, status.Error(codes.Internal, err.Error())
	}
	refreshToken, err := generateRefreshToken(ctx, entClient)
	if err != nil {
		return Tokens{}, status.Error(codes.Internal, err.Error())
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
	//todo: 重複の検証
	token, _ := MakeRandomStr(32)
	expiresAt := timejst.Now().Add(time.Hour * 24 * 30)

	q := entClient.RefreshToken.Create().
		SetToken(token).
		SetExpiresAt(expiresAt).
		SetIsDead(false)
	_, err := q.Save(ctx)
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	return token, nil
}

func killRefreshToken(ctx context.Context, entClient *ent.Client, token string) error {
	q := entClient.RefreshToken.Update().
		Where(enttoken.Token(token)).
		SetIsDead(true)
	_, err := q.Save(ctx)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
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
		return false, status.Error(codes.Internal, err.Error())
	}
	return exist, nil
}

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
