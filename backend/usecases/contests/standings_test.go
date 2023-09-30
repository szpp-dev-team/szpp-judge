package contests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

func Test_GetStandings(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)

	ctx := context.Background()
	resp, err := interactor.GetStandings(ctx, &backendv1.GetStandingsRequest{
		ContestId: 1,
	})
	require.NoError(t, err)

	t.Log(resp)
}
