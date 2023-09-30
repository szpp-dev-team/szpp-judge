package contests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/test/fixture"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

func Test_GetStandings(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.GetStandingsRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.GetStandingsRequest, resp *backendv1.GetStandingsResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *backendv1.GetTaskRequest) {
				contest_id := 
			},
			assert: func(ctx context.Context, t *testing.T, req *backendv1.GetTaskRequest, resp *backendv1.GetTaskResponse) {
				task := entClient.Task.GetX(context.Background(), int(resp.Task.Id))
				assert.Equal(t, "a", task.Title)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer utils.TruncateDB(t, entClient)

			ctx := context.Background()
			req := &backendv1.GetTaskRequest{}
			if test.prepare != nil {
				test.prepare(t, req)
			}
			resp, err := interactor.GetStandings(ctx, req)
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
