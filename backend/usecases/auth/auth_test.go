package auth

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	intercepter "github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server/intercepter"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	user "github.com/szpp-dev-team/szpp-judge/backend/usecases/user"
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

}
