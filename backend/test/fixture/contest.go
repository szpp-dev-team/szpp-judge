package fixture

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

func CreateContest(
	t *testing.T,
	client *ent.Client,
	name, slug string,
	penaltySeconds int,
	isPublic bool,
	startAt, endAt time.Time,
	contestType string,
) *ent.Contest {
	contest, err := client.Contest.Create().
		SetName(name).
		SetSlug(slug).
		SetDescription("This is the first SZPP Programming Contest.").
		SetPenaltySeconds(penaltySeconds).
		SetContestType(backendv1.ContestType_OFFICIAL.String()).
		SetIsPublic(isPublic).
		SetStartAt(startAt).
		SetEndAt(endAt).
		Save(context.Background())
	require.NoError(t, err)
	return contest
}
