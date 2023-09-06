package testcases

import "context"

type Repository interface {
	DownloadTestcase(ctx context.Context, taskID int, name string) (*Testcase, error)
	UploadTestcase(ctx context.Context, taskID int, testcase *Testcase) error
}

type Testcase struct {
	Name string
	In   []byte
	Out  []byte
}
