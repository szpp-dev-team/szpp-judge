package intercepter

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var claimsKey struct{}
var excludeMethodSet = map[string]struct{}{
	backendv1.TaskService_GetTask_FullMethodName:         {}, // 問題閲覧はゲストでもできるようにする
	backendv1.TaskService_GetTestcaseSets_FullMethodName: {}, // 入力例閲覧はゲストでもできるようにする
}

func Auth(secret []byte) grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
		method, ok := grpc.Method(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "failed to get method name from context")
		}
		// skip authentication
		if _, ok := excludeMethodSet[method]; ok {
			return ctx, nil
		}

		accessToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		claims, err := GetClaimsFromToken(accessToken, secret)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "failed to parse jwt")
		}
		return context.WithValue(ctx, claimsKey, claims), nil
	})
}

func GetClaimsFromContext(ctx context.Context) *Claims {
	claims, _ := ctx.Value(claimsKey).(*Claims)
	return claims
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
