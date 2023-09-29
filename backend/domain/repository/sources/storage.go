package sources

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

func (i *gcsImpl) UploadSource(ctx context.Context, submitID int, source []byte) error {
	bucket := i.client.Bucket("szpp-judge")
	return gcs.UploadFile(ctx, bucket, BuildSourceCodePath(submitID), source)
}

func (i *gcsImpl) DownloadSource(ctx context.Context, submitID int) ([]byte, error) {
	bucket := i.client.Bucket("szpp-judge")
	return gcs.DownloadFile(ctx, bucket, BuildSourceCodePath(submitID))
}

func BuildSourceCodePath(submitID int) string {
	return fmt.Sprintf("/submits/%d", submitID)
}
