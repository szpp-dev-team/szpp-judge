package testcases

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/szpp-dev-team/szpp-judge/backend/core/gcs"
)

type gcsImpl struct {
	client *storage.Client
}

func NewRepository(client *storage.Client) Repository {
	return &gcsImpl{client}
}

func (i *gcsImpl) DownloadTestcase(ctx context.Context, taskID int, name string) (*Testcase, error) {
	bucket := i.client.Bucket("szpp-judge")
	inBytes, err := gcs.DownloadFile(ctx, bucket, BuildTestcaseInPath(taskID, name))
	if err != nil {
		return nil, err
	}
	outBytes, err := gcs.DownloadFile(ctx, bucket, BuildTestcaseOutPath(taskID, name))
	if err != nil {
		return nil, err
	}
	return &Testcase{
		Name: name,
		In:   inBytes,
		Out:  outBytes,
	}, nil
}

func (i *gcsImpl) UpsertTestcase(ctx context.Context, taskID int, testcase *Testcase) error {
	bucket := i.client.Bucket("szpp-judge")
	if err := gcs.UploadFile(ctx, bucket, BuildTestcaseInPath(taskID, testcase.Name), testcase.In); err != nil {
		return err
	}
	if err := gcs.UploadFile(ctx, bucket, BuildTestcaseOutPath(taskID, testcase.Name), testcase.Out); err != nil {
		return err
	}
	return nil
}

func BuildTestcaseInPath(taskID int, name string) string {
	return fmt.Sprintf("testcases/%d/in/%s", taskID, name)
}

func BuildTestcaseOutPath(taskID int, name string) string {
	return fmt.Sprintf("testcases/%d/out/%s", taskID, name)
}

var _ Repository = (*gcsImpl)(nil)
