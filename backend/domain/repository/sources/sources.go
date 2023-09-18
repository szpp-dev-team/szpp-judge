package sources

import "context"

type Repository interface {
	UploadSource(ctx context.Context, submitID int, source []byte) error
	DownloadSource(ctx context.Context, submitID int) ([]byte, error)
}
