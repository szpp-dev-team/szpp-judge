package tasks

import (
	"context"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server/intercepter"
	"github.com/szpp-dev-team/szpp-judge/backend/core/timejst"
	ent_task "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	ent_testcase "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcase"
	testcases_repo "github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases/mock"
	"github.com/szpp-dev-team/szpp-judge/backend/test/fixture"
	"github.com/szpp-dev-team/szpp-judge/backend/test/utils"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
)

func Test_CreateTask(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient, nil, nil, nil)

	tests := map[string]struct {
		modify  func(ctx context.Context, req *backendv1.CreateTaskRequest) context.Context
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.CreateTaskRequest, resp *backendv1.CreateTaskResponse)
	}{
		"success": {
			assert: func(ctx context.Context, t *testing.T, req *backendv1.CreateTaskRequest, resp *backendv1.CreateTaskResponse) {
				task := entClient.Task.GetX(context.Background(), int(resp.Task.Id))
				assert.Equal(t, req.Task.Title, task.Title)
			},
		},
	}

	for name, test := range tests {
		req := &backendv1.CreateTaskRequest{
			Task: seedMutationTask(),
		}

		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			user := fixture.CreateUser(t, entClient, "hoge", "ADMIN")
			ctx = intercepter.SetClaimsToContext(ctx, &intercepter.Claims{Username: user.Username})

			if test.modify != nil {
				ctx = test.modify(ctx, req)
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

			utils.TruncateDB(t, entClient)
		})
	}
}

func Test_UpdateTask(t *testing.T) {
	entClient := utils.NewTestClient(t)
	interactor := NewInteractor(entClient, nil, nil, nil)

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.UpdateTaskRequest)
		modify  func(req *backendv1.UpdateTaskRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.UpdateTaskRequest, resp *backendv1.UpdateTaskResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *backendv1.UpdateTaskRequest) {
				task := fixture.CreateTask(t, entClient, judgev1.JudgeType_Normal{Normal: &judgev1.JudgeTypeNormal{CaseInsensitive: lo.ToPtr(false)}}, 1, nil)
				req.TaskId = int32(task.ID)
			},
			modify: func(req *backendv1.UpdateTaskRequest) {
				req.Task.Title = "foo"
			},
			assert: func(ctx context.Context, t *testing.T, req *backendv1.UpdateTaskRequest, resp *backendv1.UpdateTaskResponse) {
				task := entClient.Task.GetX(context.Background(), int(resp.Task.Id))
				assert.Equal(t, req.Task.Title, task.Title)
			},
		},
	}

	for name, test := range tests {
		req := &backendv1.UpdateTaskRequest{
			Task: seedMutationTask(),
		}

		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			user := fixture.CreateUser(t, entClient, "hoge", "ADMIN")
			ctx = intercepter.SetClaimsToContext(ctx, &intercepter.Claims{Username: user.Username})

			if test.modify != nil {
				test.modify(req)
			}
			if test.prepare != nil {
				test.prepare(t, req)
			}

			resp, err := interactor.UpdateTask(ctx, req)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if test.assert != nil {
				test.assert(ctx, t, req, resp)
			}

			utils.TruncateDB(t, entClient)
		})
	}
}

func Test_GetTask(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	interactor := NewInteractor(entClient, nil, nil, nil)

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.GetTaskRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.GetTaskRequest, resp *backendv1.GetTaskResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *backendv1.GetTaskRequest) {
				user := fixture.CreateUser(t, entClient, "hoge", "ADMIN")
				task := fixture.CreateTask(t, entClient, newJudgeTypeNormal(false), user.ID, nil)
				req.TaskId = int32(task.ID)
			},
			assert: func(ctx context.Context, t *testing.T, req *backendv1.GetTaskRequest, resp *backendv1.GetTaskResponse) {
				task := entClient.Task.GetX(context.Background(), int(resp.Task.Id))
				assert.Equal(t, "hoge", task.Title)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			req := &backendv1.GetTaskRequest{}
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

			utils.TruncateDB(t, entClient)
		})
	}
}

func Test_GetTestcaseSets(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	testcasesRepo := mock.NewMock()
	interactor := NewInteractor(entClient, testcasesRepo, nil, nil)

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.GetTestcaseSetsRequest)
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.GetTestcaseSetsRequest, resp *backendv1.GetTestcaseSetsResponse)
	}{
		"success": {
			prepare: func(t *testing.T, req *backendv1.GetTestcaseSetsRequest) {
				testcaseSets, testcases := seedMutationTestcasePairs()
				testcaseIDByName := map[string]int{}
				for _, testcase := range testcases {
					testcasesRepo.UpsertTestcase(context.Background(), int(req.TaskId), &testcases_repo.Testcase{
						Name: testcase.Slug,
						In:   []byte("a"),
						Out:  []byte("b"),
					})
					testcase := fixture.CreateTestcase(t, entClient, testcase.Slug, int(req.TaskId))
					testcaseIDByName[testcase.Name] = testcase.ID
				}

				for _, testcaseSet := range testcaseSets {
					testcaseIDs := []int{}
					for _, testcaseSlug := range testcaseSet.TestcaseSlugs {
						testcaseIDs = append(testcaseIDs, testcaseIDByName[testcaseSlug])
					}
					_ = fixture.CreateTestcaseSet(t, entClient, testcaseSet.Slug, int(req.TaskId), testcaseIDs)
				}
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			user := fixture.CreateUser(t, entClient, "hoge", "ADMIN")
			task := fixture.CreateTask(t, entClient, newJudgeTypeNormal(false), user.ID, nil)

			req := &backendv1.GetTestcaseSetsRequest{TaskId: int32(task.ID)}
			if test.prepare != nil {
				test.prepare(t, req)
			}
			resp, err := interactor.GetTestcaseSets(ctx, req)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			if test.assert != nil {
				test.assert(ctx, t, req, resp)
			}

			utils.TruncateDB(t, entClient)
		})
	}
}

func Test_SyncTestcaseSets(t *testing.T) {
	entClient := utils.NewTestClient(t)
	defer entClient.Close()
	testcasesRepo := mock.NewMock()
	interactor := NewInteractor(entClient, testcasesRepo, nil, nil)

	tests := map[string]struct {
		prepare func(t *testing.T, req *backendv1.SyncTestcaseSetsRequest)
		modify  func(ctx context.Context, req *backendv1.SyncTestcaseSetsRequest) context.Context
		wantErr bool
		assert  func(ctx context.Context, t *testing.T, req *backendv1.SyncTestcaseSetsRequest, resp *backendv1.SyncTestcaseSetsResponse)
	}{
		"create": {
			assert: func(ctx context.Context, t *testing.T, req *backendv1.SyncTestcaseSetsRequest, resp *backendv1.SyncTestcaseSetsResponse) {
				testcaseSets := entClient.TestcaseSet.Query().WithTestcases().AllX(context.Background())
				assert.Equal(t, 2, len(testcaseSets))
				testcases := entClient.Testcase.Query().AllX(context.Background())
				assert.Equal(t, 8, len(testcases))
				assert.Equal(t, 3, len(testcaseSets[0].Edges.Testcases))
				assert.Equal(t, 8, len(testcaseSets[1].Edges.Testcases))
				for _, testcase := range testcases {
					_, err := testcasesRepo.DownloadTestcase(context.Background(), int(req.TaskId), testcase.Name)
					require.NoError(t, err)
				}
			},
		},
		// testcase: killer01, killer02 will be created, and random04, random05 will be deleted.
		"replace": {
			prepare: func(t *testing.T, req *backendv1.SyncTestcaseSetsRequest) {
				testcaseSets, testcases := seedMutationTestcasePairs()
				testcaseIDByName := map[string]int{}
				for _, testcase := range testcases {
					testcasesRepo.UpsertTestcase(context.Background(), int(req.TaskId), &testcases_repo.Testcase{
						Name: testcase.Slug,
						In:   []byte("a"),
						Out:  []byte("b"),
					})
					id, err := entClient.Testcase.Create().
						SetName(testcase.Slug).
						SetCreatedAt(timejst.Now()).
						SetTaskID(int(req.TaskId)).
						OnConflict().
						UpdateNewValues().
						ID(context.Background())
					require.NoError(t, err, "failed to create testcase %s", testcase.Slug)
					testcaseIDByName[testcase.Slug] = id
				}

				for _, testcaseSet := range testcaseSets {
					testcaseIDs := []int{}
					for _, testcaseSlug := range testcaseSet.TestcaseSlugs {
						testcaseIDs = append(testcaseIDs, testcaseIDByName[testcaseSlug])
					}
					_ = fixture.CreateTestcaseSet(t, entClient, testcaseSet.Slug, int(req.TaskId), testcaseIDs)
				}
			},
			modify: func(ctx context.Context, req *backendv1.SyncTestcaseSetsRequest) context.Context {
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
							"random01", "random02", "random03", "killer01", "killer02", "killer03",
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
						Slug:   "killer01",
						Input:  "7",
						Output: "yes",
					},
					{
						Slug:   "killer02",
						Input:  "8",
						Output: "no",
					},
					{
						Slug:   "killer03",
						Input:  "8",
						Output: "no",
					},
				}
				req.TestcaseSets = testcaseSets
				req.Testcases = testcases
				return ctx
			},
			assert: func(ctx context.Context, t *testing.T, req *backendv1.SyncTestcaseSetsRequest, resp *backendv1.SyncTestcaseSetsResponse) {
				testcaseSets := entClient.TestcaseSet.Query().WithTestcases().AllX(context.Background())
				assert.Equal(t, 2, len(testcaseSets))
				testcases := entClient.Testcase.Query().AllX(context.Background())
				assert.Equal(t, 9, len(testcases))
				assert.Equal(t, 3, len(testcaseSets[0].Edges.Testcases))
				assert.Equal(t, 9, len(testcaseSets[1].Edges.Testcases))
				for _, testcase := range testcases {
					_, err := testcasesRepo.DownloadTestcase(context.Background(), int(req.TaskId), testcase.Name)
					require.NoError(t, err)
				}
				assert.False(t, entClient.Testcase.Query().Where(ent_testcase.Name("random04")).ExistX(context.Background()))
				assert.False(t, entClient.Testcase.Query().Where(ent_testcase.Name("random05")).ExistX(context.Background()))
				assert.True(t, entClient.Testcase.Query().Where(ent_testcase.Name("killer01")).ExistX(context.Background()))
				assert.True(t, entClient.Testcase.Query().Where(ent_testcase.Name("killer02")).ExistX(context.Background()))
				assert.True(t, entClient.Testcase.Query().Where(ent_testcase.Name("killer03")).ExistX(context.Background()))
			},
		},
		"permission denied": {
			modify: func(ctx context.Context, req *backendv1.SyncTestcaseSetsRequest) context.Context {
				user := fixture.CreateUser(t, entClient, "fuga", "ADMIN")
				return intercepter.SetClaimsToContext(ctx, &intercepter.Claims{Username: user.Username})
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			user := fixture.CreateUser(t, entClient, "hoge", "ADMIN")
			ctx = intercepter.SetClaimsToContext(ctx, &intercepter.Claims{Username: user.Username})

			task := fixture.CreateTask(t, entClient, ent_task.JudgeTypeNormal, int(user.ID), nil)

			testcaseSets, testcases := seedMutationTestcasePairs()
			req := &backendv1.SyncTestcaseSetsRequest{
				TaskId:       int32(task.ID),
				TestcaseSets: testcaseSets,
				Testcases:    testcases,
			}
			if test.prepare != nil {
				test.prepare(t, req)
			}
			if test.modify != nil {
				ctx = test.modify(ctx, req)
			}
			resp, err := interactor.SyncTestcaseSets(ctx, req)
			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			if test.assert != nil {
				test.assert(ctx, t, req, resp)
			}

			utils.TruncateDB(t, entClient)
		})
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

func newJudgeTypeNormal(b bool) *judgev1.JudgeType_Normal {
	return &judgev1.JudgeType_Normal{
		Normal: &judgev1.JudgeTypeNormal{
			CaseInsensitive: lo.ToPtr(b),
		},
	}
}

//lint:ignore U1000 this function will be used in the future
func newJudgeTypeInteractive(p string) *judgev1.JudgeType_Interactive {
	return &judgev1.JudgeType_Interactive{
		Interactive: &judgev1.JudgeTypeInteractive{
			JudgeCodePath: p,
		},
	}
}

//lint:ignore U1000 this function will be used in the future
func newJudgeTypeEps(n uint32) *judgev1.JudgeType_Eps {
	return &judgev1.JudgeType_Eps{
		Eps: &judgev1.JudgeTypeEPS{
			Ndigits: n,
		},
	}
}

//lint:ignore U1000 this function will be used in the future
func newJudgeTypeCustom(p string) *judgev1.JudgeType_Custom {
	return &judgev1.JudgeType_Custom{
		Custom: &judgev1.JudgeTypeCustom{
			JudgeCodePath: p,
		},
	}
}
