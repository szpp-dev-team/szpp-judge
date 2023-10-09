package checkers

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/szpp-dev-team/szpp-judge/backend/core/gcs"
)

type gcsImpl struct {
	client     *storage.Client
	bucketName string
}

func NewRepository(client *storage.Client, bucketName string) Repository {
	return &gcsImpl{client, bucketName}
}

func (i *gcsImpl) UploadChecker(ctx context.Context, taskID int, checker []byte) error {
	bucket := i.client.Bucket(i.bucketName)
	return gcs.UploadFile(ctx, bucket, BuildCheckerCodePath(taskID), checker)
}

func (i *gcsImpl) DownloadChecker(ctx context.Context, taskID int) ([]byte, error) {
	bucket := i.client.Bucket(i.bucketName)
	return gcs.DownloadFile(ctx, bucket, BuildCheckerCodePath(taskID))
}

func BuildCheckerCodePath(taskID int) string {
	return fmt.Sprintf("tasks/%d/checker.cpp", taskID)
}
