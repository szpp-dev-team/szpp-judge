package tasks

import (
	"context"
	"fmt"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	ent_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases/mock"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
)

func Test_CreateTask(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	testcasesRepo := mock.NewMock()
	interactor := NewInteractor(entClient, testcasesRepo)
	now := timejst.Now()

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.CreateTaskRequest)
		modify  func(req *backendv1.CreateTaskRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.CreateTaskRequest, resp *backendv1.CreateTaskResponse)
	}{
		"success": {
			assert: func(ctx context.Context, t *testing.T, req *backendv1.CreateTaskRequest, resp *backendv1.CreateTaskResponse) {
				task, err := entClient.Task.Query().
					WithTestcases().
					WithTestcaseSets().
					Where(ent_task.ID(int(resp.Task.Id))).Only(ctx)
				require.NoError(t, err)

				assert.Equal(t, req.Task.Title, task.Title)

				assert.Equal(t, len(req.Testcases), len(task.Edges.Testcases))
				for _, testcase := range req.Testcases {
					testcase2, err := testcasesRepo.DownloadTestcase(ctx, task.ID, testcase.Slug)
					require.NoError(t, err)
					assert.Equal(t, testcase.Input, string(testcase2.In))
					assert.Equal(t, testcase.Output, string(testcase2.Out))
				}

				assert.Equal(t, len(req.TestcaseSets), len(task.Edges.TestcaseSets))
			},
		},
		"DB constraints": {
			prepare: func(t *testing.T, req *backendv1.CreateTaskRequest) {
				task1, err := entClient.Task.Create().SetTitle("hoge").SetStatement("fuga").SetDifficulty(backendv1.Difficulty_BEGINNER.String()).
					SetExecTimeLimit(2000).SetExecMemoryLimit(1024).SetJudgeType(ent_task.JudgeTypeNormal).SetCaseInsensitive(false).
					SetCreatedAt(now).Save(context.Background())
				require.NoError(t, err)
				task2, err := entClient.Task.Create().SetTitle("hoge").SetStatement("fuga").SetDifficulty(backendv1.Difficulty_BEGINNER.String()).
					SetExecTimeLimit(2000).SetExecMemoryLimit(1024).SetJudgeType(ent_task.JudgeTypeNormal).SetCaseInsensitive(false).
					SetCreatedAt(now).Save(context.Background())
				require.NoError(t, err)

				_, err = entClient.Testcase.Create().SetName("foo").SetTask(task1).SetCreatedAt(now).Save(context.Background())
				assert.NoError(t, err)
				_, err = entClient.Testcase.Create().SetName("foo").SetTask(task1).SetCreatedAt(now).Save(context.Background())
				assert.Error(t, err) // 同じ問題の中に同じ Testcase は作れない
				_, err = entClient.Testcase.Create().SetName("foo").SetTask(task2).SetCreatedAt(now).Save(context.Background())
				assert.NoError(t, err)

				_, err = entClient.TestcaseSet.Create().SetName("foo").SetScore(100).SetIsSample(false).SetTask(task1).SetCreatedAt(now).Save(context.Background())
				assert.NoError(t, err)
				_, err = entClient.TestcaseSet.Create().SetName("foo").SetScore(100).SetIsSample(false).SetTask(task1).SetCreatedAt(now).Save(context.Background())
				assert.Error(t, err) // 同じ問題の中に同じ TestcaseSet は作れない
				_, err = entClient.TestcaseSet.Create().SetName("foo").SetScore(100).SetIsSample(false).SetTask(task2).SetCreatedAt(now).Save(context.Background())
				assert.NoError(t, err)

			},
		},
	}

	for name, test := range tests {
		testcaseSets, testcases := seedMutationTestcasePairs()
		req := &backendv1.CreateTaskRequest{
			Task:         seedMutationTask(),
			TestcaseSets: testcaseSets,
			Testcases:    testcases,
		}

		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			if test.modify != nil {
				test.modify(req)
			}
			if test.prepare != nil {
				test.prepare(t, req)
			}

			resp, err := interactor.CreateTask(ctx, req)
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

func Test_GetTask(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	testcasesRepo := mock.NewMock()
	interactor := NewInteractor(entClient, testcasesRepo)

	now := timejst.Now()

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.GetTaskRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.GetTaskRequest, resp *backendv1.GetTaskResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *backendv1.GetTaskRequest) {
				q := entClient.Task.Create().
					SetTitle("hoge").
					SetStatement("fuga").
					SetDifficulty(backendv1.Difficulty_BEGINNER.String()).
					SetExecTimeLimit(2000).
					SetExecMemoryLimit(1024).
					SetJudgeType(ent_task.JudgeTypeNormal).
					SetCaseInsensitive(false).
					SetCreatedAt(now)
				testcaseIDs := []int{}
				for i := range make([]struct{}, 4) {
					name := fmt.Sprintf("%02d", i)
					testcase, err := entClient.Testcase.Create().
						SetName(name).
						SetCreatedAt(now).
						Save(context.Background())
					require.NoError(t, err)
					q.AddTestcaseIDs(testcase.ID)
					testcaseIDs = append(testcaseIDs, testcase.ID)
				}
				ts1, err := entClient.TestcaseSet.Create().SetName("all").SetScore(100).SetIsSample(false).SetCreatedAt(now).AddTestcaseIDs(testcaseIDs...).Save(context.Background())
				require.NoError(t, err)
				ts2, err := entClient.TestcaseSet.Create().SetName("sample").SetScore(0).SetIsSample(true).SetCreatedAt(now).AddTestcaseIDs(testcaseIDs[:2]...).Save(context.Background())
				require.NoError(t, err)
				q.AddTestcaseSetIDs(ts1.ID, ts2.ID)
				task, err := q.Save(context.Background())
				require.NoError(t, err)
				req.Id = int32(task.ID)

				for i := range testcaseIDs {
					err := testcasesRepo.UploadTestcase(context.Background(), task.ID, &testcases.Testcase{
						Name: fmt.Sprintf("%02d", i),
						In:   []byte("a"),
						Out:  []byte("b"),
					})
					require.NoError(t, err)
				}
			},
			assert: func(ctx context.Context, t *testing.T, req *backendv1.GetTaskRequest, resp *backendv1.GetTaskResponse) {
				task, err := entClient.Task.Query().
					WithTestcases().
					WithTestcaseSets().
					Where(ent_task.ID(int(resp.Task.Id))).Only(ctx)
				require.NoError(t, err)
				assert.Equal(t, "hoge", task.Title)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			req := &backendv1.GetTaskRequest{Id: 1}
			if test.prepare != nil {
				test.prepare(t, req)
			}
			resp, err := interactor.GetTask(ctx, req)
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

func makeMutationTestcase(slug, input, output string) *backendv1.MutationTestcase {
	return &backendv1.MutationTestcase{
		Slug:   slug,
		Input:  input,
		Output: output,
	}
}

func seedMutationTask() *backendv1.MutationTask {
	return &backendv1.MutationTask{
		Title:           "hoge",
		Statement:       "fuga",
		ExecTimeLimit:   2000,
		ExecMemoryLimit: 1024,
		JudgeType: &judgev1.JudgeType{
			JudgeType: &judgev1.JudgeType_Normal{
				Normal: &judgev1.JudgeTypeNormal{
					CaseInsensitive: lo.ToPtr(false),
				},
			},
		},
		Difficulty: backendv1.Difficulty_EASY,
	}
}

func seedMutationTestcasePairs() ([]*backendv1.MutationTestcaseSet, []*backendv1.MutationTestcase) {
	testcaseSets := []*backendv1.MutationTestcaseSet{
		{
			Slug:          "sample",
			Score:         0,
			IsSample:      true,
			TestcaseSlugs: []string{"example01", "example02", "example03"},
		},
		{
			Slug:  "all",
			Score: 100,
			TestcaseSlugs: []string{
				"example01", "example02", "example03",
				"random01", "random02", "random03", "random04", "random05",
			},
		},
	}
	testcases := []*backendv1.MutationTestcase{
		{
			Slug:   "example01",
			Input:  "1",
			Output: "yes",
		},
		{
			Slug:   "example02",
			Input:  "2",
			Output: "no",
		},
		{
			Slug:   "example03",
			Input:  "3",
			Output: "yes",
		},
		{
			Slug:   "random01",
			Input:  "4",
			Output: "no",
		},
		{
			Slug:   "random02",
			Input:  "5",
			Output: "yes",
		},
		{
			Slug:   "random03",
			Input:  "6",
			Output: "no",
		},
		{
			Slug:   "random04",
			Input:  "7",
			Output: "yes",
		},
		{
			Slug:   "random05",
			Input:  "8",
			Output: "no",
		},
	}
	return testcaseSets, testcases
}
