package auth

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	intercepter "github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server/intercepter"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	enttoken "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/refreshtoken"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/user"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

func Test_Login(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	secret := "JWT_SECRET"
	interactor := NewInteractor(entClient, secret)
	now := timejst.Now()

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.LoginRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.LoginRequest, resp *backendv1.LoginResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *backendv1.LoginRequest) {
				q := entClient.User.Create().
					SetUsername("authTestUser0").
					SetHashedPassword(user.HashPassword("tooreeee")).
					SetEmail("szppi_authtestuser0@szpp.com").
					SetRole("USER").
					SetCreatedAt(now).
					SetUpdatedAt(now)
				_, err := q.Save(context.Background())
				require.NoError(t, err)
			},
			assert: func(ctx context.Context, t *testing.T, req *backendv1.LoginRequest, resp *backendv1.LoginResponse) {
				assert.Equal(t, req.Username, resp.User.Username)
				username := intercepter.GetClaimsFromContext(ctx).Username
				assert.Equal(t, req.Username, username)
			},
		},
		"wrong password": {
			prepare: func(t *testing.T, req *backendv1.LoginRequest) {
				q := entClient.User.Create().
					SetUsername("authTestUser1").
					SetHashedPassword(user.HashPassword("toosanai")).
					SetEmail("szppi_authtestuser1@szpp.com").
					SetRole("USER").
					SetCreatedAt(now).
					SetUpdatedAt(now)
				_, err := q.Save(context.Background())
				require.NoError(t, err)
			},
			wantErr: true,
		},
	}

	i := 0
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			ctx = intercepter.SetClaimsToContext(ctx, &intercepter.Claims{Username: fmt.Sprintf("authTestUser%d", i)})
			req := &backendv1.LoginRequest{
				Username: fmt.Sprintf("authTestUser%d", i),
				Password: "tooreeee",
			}
			if test.prepare != nil {
				test.prepare(t, req)
			}
			resp, err := interactor.Login(ctx, req, secret)
			if test.wantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			if test.assert != nil {
				test.assert(ctx, t, req, resp)
			}
		})
		i++
	}

}

func Test_Logout(t *testing.T) {

}

func Test_RefreshAccessToken(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	secret := "JWT_SECRET"
	interactor := NewInteractor(entClient, secret)
	tests := map[string]struct {
		prepare func(t *testing.T) string
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.RefreshAccessTokenRequest, resp *backendv1.RefreshAccessTokenResponse)
	}{
		"success": {
			// refreshTokenの期限が有効であるときに正しくaccessTokenが発行される
			prepare: func(t *testing.T) string {
				username := "tokenTestUser0"
				q := entClient.User.Create().
					SetUsername(username).
					SetHashedPassword(user.HashPassword("tooreeee")).
					SetEmail("tokenTestUser0.szpp.com").
					SetRole("USER").
					SetCreatedAt(timejst.Now()).
					SetUpdatedAt(timejst.Now())
				_, err := q.Save(context.Background())
				ctx := context.Background()
				ctx = intercepter.SetClaimsToContext(ctx, &intercepter.Claims{Username: username})
				loginReq := &backendv1.LoginRequest{
					Username: username,
					Password: "tooreeee",
				}
				loginResp, err := interactor.Login(ctx, loginReq, secret)
				refreshToken := loginResp.RefreshToken
				require.NoError(t, err)
				return refreshToken
			},
			wantErr: false,
			assert: func(ctx context.Context, t *testing.T, req *backendv1.RefreshAccessTokenRequest, resp *backendv1.RefreshAccessTokenResponse) {
				token, err := entClient.RefreshToken.Query().Where(enttoken.Token(req.RefreshToken)).Only(ctx)
				require.NoError(t, err)
				assert.False(t, token.IsDead)
			},
		},
		"invalid refresh token": {
			// refreshTokenが死んでいるときにエラーが返る killRefreshToken()のテスト
			prepare: func(t *testing.T) string {
				username := "tokenTestUser1"
				q := entClient.User.Create().
					SetUsername(username).
					SetHashedPassword(user.HashPassword("tooreeee")).
					SetEmail("tokenTestUser1.szpp.com").
					SetRole("USER").
					SetCreatedAt(timejst.Now()).
					SetUpdatedAt(timejst.Now())
				_, err := q.Save(context.Background())
				ctx := context.Background()
				ctx = intercepter.SetClaimsToContext(ctx, &intercepter.Claims{Username: username})
				loginReq := &backendv1.LoginRequest{
					Username: username,
					Password: "tooreeee",
				}
				loginResp, err := interactor.Login(ctx, loginReq, secret)
				refreshToken := loginResp.RefreshToken
				killRefreshToken(ctx, entClient, refreshToken)
				require.NoError(t, err)
				return refreshToken
			},
			wantErr: true,
			assert: func(ctx context.Context, t *testing.T, req *backendv1.RefreshAccessTokenRequest, resp *backendv1.RefreshAccessTokenResponse) {
				token, err := entClient.RefreshToken.Query().Where(enttoken.Token(req.RefreshToken)).Only(ctx)
				require.NoError(t, err)
				assert.True(t, token.IsDead)
			},
		},
	}

	i := 0
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			ctx = intercepter.SetClaimsToContext(ctx, &intercepter.Claims{Username: fmt.Sprintf("authTestUser%d", i)})

			refreshToken := ""
			if test.prepare != nil {
				refreshToken = test.prepare(t)
			}
			req := &backendv1.RefreshAccessTokenRequest{
				RefreshToken: refreshToken,
			}
			resp, err := interactor.RefreshAccessToken(ctx, req)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			if test.assert != nil {
				test.assert(ctx, t, req, resp)
			}
		})
	}
}
