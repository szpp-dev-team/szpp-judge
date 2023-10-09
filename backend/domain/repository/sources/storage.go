package sources

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

func (i *gcsImpl) UploadSource(ctx context.Context, submitID int, source []byte) error {
	bucket := i.client.Bucket(i.bucketName)
	return gcs.UploadFile(ctx, bucket, BuildSourceCodePath(submitID), source)
}

func (i *gcsImpl) DownloadSource(ctx context.Context, submitID int) ([]byte, error) {
	bucket := i.client.Bucket(i.bucketName)
	return gcs.DownloadFile(ctx, bucket, BuildSourceCodePath(submitID))
}

func BuildSourceCodePath(submitID int) string {
	return fmt.Sprintf("submits/%d", submitID)
}
