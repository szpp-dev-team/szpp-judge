package fixture

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
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

func CreateTask(t *testing.T, client *ent.Client, title string, jt any, userID int, contestID *int) *ent.Task {
	t.Helper()

	q := client.Task.Create().
		SetTitle(title).
		SetStatement("fuga").
		SetDifficulty(backendv1.Difficulty_BEGINNER.String()).
		SetExecTimeLimit(2000).
		SetExecMemoryLimit(1024).
		SetUserID(userID).
		SetCreatedAt(timejst.Now())

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
		SetScoreRatio(100).
		SetIsSample(false).
		SetTaskID(taskID).
		AddTestcaseIDs(testcaseIDs...).
		SetCreatedAt(timejst.Now()).
		Save(context.Background())
	require.NoError(t, err)
	return testcaseSet
}
