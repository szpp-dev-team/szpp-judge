package checkers

import "context"

type Repository interface {
	UploadChecker(ctx context.Context, taskID int, source []byte) error
	DownloadChecker(ctx context.Context, taskID int) ([]byte, error)
}
