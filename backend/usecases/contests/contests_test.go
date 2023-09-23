package contests

import (
	"context"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_contest "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	"github.com/szpp-dev-team/szpp-judge/backend/test/fixture"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_CreateContest(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)

	tests := map[string]struct {
		assert func(t *testing.T)
	}{
		"success": {
			assert: func(t *testing.T) {
				contest, err := entClient.Contest.Query().
					Where(ent_contest.Slug("szpc001")).
					Only(context.Background())
				require.NoError(t, err)
				require.Equal(t, "SZPP Programming Contest 001", contest.Name)
				require.Equal(t, "szpc001", contest.Slug)
				require.Equal(t, "This is the first SZPP Programming Contest.", contest.Description)
				require.Equal(t, 300, contest.PenaltySeconds)
				require.Equal(t, true, contest.IsPublic)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer utils.TruncateDB(t, entClient)
			req := &backendv1.CreateContestRequest{
				Contest: &backendv1.MutationContest{
					Name:           "SZPP Programming Contest 001",
					Slug:           "szpc001",
					Description:    "This is the first SZPP Programming Contest.",
					PenaltySeconds: 300,
					StartAt:        timestamppb.New(time.Now()),
					EndAt:          timestamppb.New(time.Now().Add(1 * time.Hour)),
					IsPublic:       true,
					ContestType:    backendv1.ContestType_OFFICIAL,
				},
			}
			_, err := interactor.CreateContest(context.Background(), req)
			require.NoError(t, err)
			if test.assert != nil {
				test.assert(t)
			}
		})
	}
}

func Test_GetContest(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)

	tests := map[string]struct {
		assert func(t *testing.T, contest *backendv1.GetContestResponse)
	}{
		"success": {
			assert: func(t *testing.T, contest *backendv1.GetContestResponse) {
				require.Equal(t, "SZPP Programming Contest", contest.Contest.Name)
				assert.Equal(t, "szpc001", contest.Contest.Slug)
				assert.Equal(t, "This is the first SZPP Programming Contest.", contest.Contest.Description)
				assert.Equal(t, int32(300), contest.Contest.PenaltySeconds)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer utils.TruncateDB(t, entClient)
			contest := fixture.CreateContest(t, entClient,
				"SZPP Programming Contest", "szpc001",
				300, true,
				time.Now(), time.Now().Add(time.Hour),
				backendv1.ContestType_OFFICIAL.String())
			req := &backendv1.GetContestRequest{
				Slug: contest.Slug,
			}
			resp, err := interactor.GetContest(context.Background(), req)
			require.NoError(t, err)
			if test.assert != nil {
				test.assert(t, resp)
			}
		})
	}
}

func Test_UpdateContest(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)

	tests := map[string]struct {
		modify func(t *testing.T, req *backendv1.UpdateContestRequest)
		assert func(t *testing.T, resp *backendv1.UpdateContestResponse)
	}{
		"success": {
			modify: func(t *testing.T, req *backendv1.UpdateContestRequest) {
				req.Contest.Name = "SZPP Programming Contest 002"
				req.Contest.Slug = "szpc002"
				req.Contest.PenaltySeconds = 600
			},
			assert: func(t *testing.T, resp *backendv1.UpdateContestResponse) {
				contest, err := entClient.Contest.Get(context.Background(), int(resp.Contest.Id))
				require.NoError(t, err)
				assert.Equal(t, "SZPP Programming Contest 002", contest.Name)
				assert.Equal(t, "szpc002", contest.Slug)
				assert.Equal(t, 600, contest.PenaltySeconds)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer utils.TruncateDB(t, entClient)
			contest := fixture.CreateContest(t, entClient,
				"SZPP Programming Contest", "szpc001",
				300, true,
				time.Now(), time.Now().Add(time.Hour),
				backendv1.ContestType_OFFICIAL.String())
			req := &backendv1.UpdateContestRequest{
				Id: int32(contest.ID),
				Contest: &backendv1.MutationContest{
					Name:           contest.Name,
					Slug:           contest.Slug,
					Description:    contest.Description,
					PenaltySeconds: int32(contest.PenaltySeconds),
					StartAt:        timestamppb.New(contest.StartAt),
					EndAt:          timestamppb.New(contest.EndAt),
					IsPublic:       contest.IsPublic,
					ContestType:    backendv1.ContestType_OFFICIAL,
				},
			}
			if test.modify != nil {
				test.modify(t, req)
			}
			resp, err := interactor.UpdateContest(context.Background(), req)
			require.NoError(t, err)
			if test.assert != nil {
				test.assert(t, resp)
			}
		})
	}
}

func Test_SyncContestTasks(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient)

	contest := fixture.CreateContest(t, entClient,
		"SZPP Programming Contest", "szpc001",
		300, true,
		time.Now(), time.Now().Add(time.Hour),
		backendv1.ContestType_OFFICIAL.String())
	normalJudge := judgev1.JudgeType_Normal{Normal: &judgev1.JudgeTypeNormal{CaseInsensitive: lo.ToPtr(false)}}
	user := fixture.CreateUser(t, entClient, "hoge", "ADMIN")
	task1 := fixture.CreateTask(t, entClient, "a", normalJudge, user.ID, &contest.ID)
	task2 := fixture.CreateTask(t, entClient, "b", normalJudge, user.ID, &contest.ID)
	task3 := fixture.CreateTask(t, entClient, "c", normalJudge, user.ID, &contest.ID)

	tests := map[string]struct {
		prepare func(t *testing.T, interactor *Interactor, contest *ent.Contest, task1, task2, task3 *ent.Task)
		modify  func(t *testing.T, req *backendv1.SyncContestTasksRequest)
		assert  func(t *testing.T, entClient *ent.Client, resp *backendv1.SyncContestTasksResponse)
	}{
		"success": {
			assert: func(t *testing.T, entClient *ent.Client, resp *backendv1.SyncContestTasksResponse) {
				contest, err := entClient.Contest.Query().
					WithContestTask(func(ctq *ent.ContestTaskQuery) {
						ctq.WithTask()
					}).
					Where(ent_contest.Slug("szpc001")).
					Only(context.Background())
				require.NoError(t, err)
				require.Len(t, contest.Edges.ContestTask, 3)
				assert.Equal(t, "a", contest.Edges.ContestTask[0].Edges.Task.Title)
				assert.Equal(t, "b", contest.Edges.ContestTask[1].Edges.Task.Title)
				assert.Equal(t, "c", contest.Edges.ContestTask[2].Edges.Task.Title)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := &backendv1.SyncContestTasksRequest{
				ContestSlug: contest.Slug,
				Tasks: []*backendv1.SyncContestTasksRequest_Task{
					{Id: int32(task1.ID), Score: 100},
					{Id: int32(task2.ID), Score: 200},
					{Id: int32(task3.ID), Score: 300},
				},
			}
			if test.prepare != nil {
				test.prepare(t, interactor, contest, task1, task2, task3)
			}
			if test.modify != nil {
				test.modify(t, req)
			}
			resp, err := interactor.SyncContestTasks(context.Background(), req)
			require.NoError(t, err)
			if test.assert != nil {
				test.assert(t, entClient, resp)
			}
		})
	}
}
