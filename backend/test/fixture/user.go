package fixture

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	ent_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
)

func CreateUser(t *testing.T, client *ent.Client, username, role string) *ent.User {
	t.Helper()
	user, err := client.User.Create().
		SetUsername(username).
		SetEmail(username + "@example.com").
		SetHashedPassword([]byte("9A8B7C6D")).
		SetRole(role).
		SetCreatedAt(timejst.Now()).
		Save(context.Background())
	require.NoError(t, err)
	return user
}

func CreateTask(t *testing.T, client *ent.Client, jt any, userID int, contestID *int) *ent.Task {
	t.Helper()

	q := client.Task.Create().
		SetTitle("hoge").
		SetStatement("fuga").
		SetDifficulty(backendv1.Difficulty_BEGINNER.String()).
		SetExecTimeLimit(2000).
		SetExecMemoryLimit(1024).
		SetJudgeType(ent_task.JudgeTypeNormal).
		SetUserID(userID).
		SetCreatedAt(timejst.Now())
	// if contestID != nil {
	// 	q = q.SetContestID(*contestID)
	// }

	switch jt := jt.(type) {
	case *judgev1.JudgeType_Normal:
		q = q.SetJudgeType(ent_task.JudgeTypeNormal).
			SetCaseInsensitive(jt.Normal.GetCaseInsensitive())
	case *judgev1.JudgeType_Interactive:
		q = q.SetJudgeType(ent_task.JudgeTypeInteractive).
			SetJudgeCodePath(jt.Interactive.GetJudgeCodePath())
	case *judgev1.JudgeType_Eps:
		q = q.SetJudgeType(ent_task.JudgeTypeEps).
			SetNdigits(uint(jt.Eps.GetNdigits()))
	case *judgev1.JudgeType_Custom:
		q = q.SetJudgeType(ent_task.JudgeTypeCustom).
			SetJudgeCodePath(jt.Custom.GetJudgeCodePath())
	}

	task, err := q.Save(context.Background())
	require.NoError(t, err)
	return task
}

func CreateTestcase(t *testing.T, client *ent.Client, slug string, taskID int) *ent.Testcase {
	t.Helper()
	testcase, err := client.Testcase.Create().
		SetName(slug).
		SetTaskID(taskID).
		SetCreatedAt(timejst.Now()).
		Save(context.Background())
	require.NoError(t, err)
	return testcase
}

func CreateTestcaseSet(t *testing.T, client *ent.Client, slug string, taskID int, testcaseIDs []int) *ent.TestcaseSet {
	t.Helper()
	testcaseSet, err := client.TestcaseSet.Create().
		SetName(slug).
		SetScore(100).
		SetIsSample(false).
		SetTaskID(taskID).
		AddTestcaseIDs(testcaseIDs...).
		SetCreatedAt(timejst.Now()).
		Save(context.Background())
	require.NoError(t, err)
	return testcaseSet
}
