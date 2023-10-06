package interceptor

import (
	"context"
	"errors"
	"log"
	"strings"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var claimsKey struct{}
var excludeMethodSet = map[string]struct{}{
	// Task
	backendv1.TaskService_GetTask_FullMethodName:         {}, // 問題閲覧はゲストでもできるようにする
	backendv1.TaskService_GetTestcaseSets_FullMethodName: {}, // 入力例閲覧はゲストでもできるようにする
	// Auth
	backendv1.AuthService_Login_FullMethodName:              {},
	backendv1.AuthService_RefreshAccessToken_FullMethodName: {},
	// User
	backendv1.UserService_ExistsEmail_FullMethodName:    {},
	backendv1.UserService_ExistsUsername_FullMethodName: {},
	backendv1.UserService_GetUser_FullMethodName:        {},
	backendv1.UserService_CreateUser_FullMethodName:     {},
	// Contest
	backendv1.ContestService_GetContest_FullMethodName:   {},
	backendv1.ContestService_ListContests_FullMethodName: {},
	backendv1.ContestService_GetStandings_FullMethodName: {},
	// Judge
	backendv1.JudgeService_GetSubmissionDetail_FullMethodName: {},
	backendv1.JudgeService_ListSubmissions_FullMethodName:     {},
}

func Auth(secret []byte) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// skip authentication
			if _, ok := excludeMethodSet[req.Spec().Procedure]; ok {
				return next(ctx, req)
			}

			accessToken := req.Header().Get("Authorization")
			log.Println("AccessToken: " + accessToken)
			claims, err := GetClaimsFromToken(strings.TrimPrefix(accessToken, "Bearer "), secret)
			if err != nil {
				log.Println(err)
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("failed to parse jwt"))
			}
			ctx = context.WithValue(ctx, claimsKey, claims)
			return next(ctx, req)
		}
	}
}

func GetClaimsFromContext(ctx context.Context) *Claims {
	claims, _ := ctx.Value(claimsKey).(*Claims)
	return claims
}

func SetClaimsToContext(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func GetClaimsFromToken(accessToken string, secret []byte) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	claims, _ := jwtToken.Claims.(*Claims)
	if claims == nil {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
