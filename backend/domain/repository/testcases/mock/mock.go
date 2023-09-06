package mock

import (
	"context"
	"fmt"

	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockImpl struct {
	testcaseMap map[string]*testcases.Testcase
}

func NewMock() testcases.Repository {
	return &mockImpl{
		testcaseMap: make(map[string]*testcases.Testcase),
	}
}

func (i *mockImpl) DownloadTestcase(ctx context.Context, taskID int, name string) (*testcases.Testcase, error) {
	testcase, ok := i.testcaseMap[generateKey(taskID, name)]
	if !ok {
		return nil, status.Error(codes.NotFound, "")
	}
	return testcase, nil
}

func (i *mockImpl) UploadTestcase(ctx context.Context, taskID int, testcase *testcases.Testcase) error {
	key := generateKey(taskID, testcase.Name)
	if _, ok := i.testcaseMap[key]; ok {
		return status.Error(codes.AlreadyExists, "")
	}
	i.testcaseMap[key] = testcase
	return nil
}

func generateKey(taskID int, name string) string {
	return fmt.Sprintf("%d_%s", taskID, name)
}
