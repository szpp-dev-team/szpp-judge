package contests

import (
	"context"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

func Test_GetStandings(t *testing.T) {
	println("check01111111111111111111111111111111111111")
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	println("check0222222222222222222222222222222222222")
	interactor := NewInteractor(entClient)
	println("check03")
	var req *backendv1.GetStandingsRequest
	req.ContestId = 1
	println("check04")
	standings, err := interactor.GetStandings(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	println("check05")
	for _, row := range standings.StandingsList {
		print(strconv.Itoa(int(row.Rank)) + " " + row.LatestAcAt.String() + " " + row.Username)
	}

	test := map[string]struct {
		prepare func(t *testing.T, req *backendv1.GetStandingsRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.GetStandingsRequest, resp *backendv1.GetStandingsResponse)
	}{
		"success": {
			assert: func(ctx context.Context, t *testing.T, req *backendv1.GetStandingsRequest, resp *backendv1.GetStandingsResponse) {
				assert.Equal(t, "hoge", "hoge")
			},
		},
	}

	test = test
}
