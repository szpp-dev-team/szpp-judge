package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	entuser "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

func Test_GetUser(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)
	now := timejst.Now()

	tests := map[string]struct {
		prepare func(t *testing.T, req *pb.GetUserRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *pb.GetUserRequest, resp *pb.GetUserResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *pb.GetUserRequest) {
				q := entClient.User.Create().
					SetUsername("hogege").
					SetHashedPassword(HashPassword("fugafuga")).
					SetEmail("szppi1@szpp.com").
					SetRole("ADMIN").
					SetCreatedAt(now).
					SetUpdatedAt(now)
				_, err := q.Save(context.Background())
				require.NoError(t, err)
			},
			assert: func(ctx context.Context, t *testing.T, req *pb.GetUserRequest, resp *pb.GetUserResponse) {
				user, err := entClient.User.Query().Where(entuser.ID(int(resp.User.Id))).Only(ctx)
				require.NoError(t, err)
				assert.Equal(t, "hogege", user.Username)
				err = VerifyPassword(user.HashedPassword, "fugafuga")
				require.NoError(t, err)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			req := &pb.GetUserRequest{Username: "hogege"}
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

func Test_CreateUser(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)

	tests := map[string]struct {
		prepare func(t *testing.T, req *pb.CreateUserRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *pb.CreateUserRequest, resp *pb.CreateUserResponse)
	}{
		"success": {
			assert: func(ctx context.Context, t *testing.T, req *pb.CreateUserRequest, resp *pb.CreateUserResponse) {
				user, err := entClient.User.Query().Where(entuser.ID(int(resp.User.Id))).Only(ctx)
				require.NoError(t, err)
				assert.Equal(t, req.Username, user.Username)
				err = VerifyPassword(user.HashedPassword, req.Password)
				require.NoError(t, err)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			req := &pb.CreateUserRequest{
				Username: "szppi",
				Password: "szppi_tensai",
				Email:    "szppi@szpp.com",
			}
			resp, err := interactor.CreateUser(ctx, req)
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

func Test_ExistsUsername(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)

	tests := map[string]struct {
		prepare func(t *testing.T, req *pb.ExistsUsernameRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *pb.ExistsUsernameRequest, resp *pb.ExistsUsernameResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *pb.ExistsUsernameRequest) {
				q := entClient.User.Create().
					SetUsername(req.Username).
					SetHashedPassword(HashPassword("fugafuga")).
					SetEmail("szppi2@szpp.com").
					SetRole("ADMIN").
					SetCreatedAt(timejst.Now()).
					SetUpdatedAt(timejst.Now())
				_, err := q.Save(context.Background())
				require.NoError(t, err)
			},
			assert: func(ctx context.Context, t *testing.T, req *pb.ExistsUsernameRequest, resp *pb.ExistsUsernameResponse) {
				assert.True(t, resp.Exists)
			},
		},
		"not found": {
			assert: func(ctx context.Context, t *testing.T, req *pb.ExistsUsernameRequest, resp *pb.ExistsUsernameResponse) {
				assert.False(t, resp.Exists)
			},
		},
	}

	i := 0
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			req := &pb.ExistsUsernameRequest{Username: fmt.Sprintf("hogegege%d", i)}
			if test.prepare != nil {
				test.prepare(t, req)
			}
			resp, err := interactor.ExistsUsername(ctx, req)
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
