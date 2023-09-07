package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	entuser "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

func Test_GetUser(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)
	now := timejst.Now()

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.GetUserRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.GetUserRequest, resp *backendv1.GetUserResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *backendv1.GetUserRequest) {
				q := entClient.User.Create().
					SetName("hogege").
					SetEncryptedPassword(HashPassword("fugafuga")).
					SetRole("ADMIN").
					SetCreatedAt(now).
					SetUpdatedAt(now)
				_, err := q.Save(context.Background())
				require.NoError(t, err)
			},
			assert: func(ctx context.Context, t *testing.T, req *backendv1.GetUserRequest, resp *backendv1.GetUserResponse) {
				user, err := entClient.User.Query().Where(entuser.ID(int(resp.User.Id))).Only(ctx)
				require.NoError(t, err)
				assert.Equal(t, "hogege", user.Name)
				err = VerifyPassword(user.EncryptedPassword, "fugafuga")
				require.NoError(t, err)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			req := &backendv1.GetUserRequest{Username: "hogege"}
			if test.prepare != nil {
				test.prepare(t, req)
			}
			resp, err := interactor.GetUser(ctx, req)
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
	}
}
