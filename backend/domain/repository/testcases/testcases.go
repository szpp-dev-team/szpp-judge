package testcases

import "context"

type Repository interface {
	DownloadTestcases(ctx context.Context, taskID int) ([]*Testcase, error)
}

type Testcase struct {
	TaskID int
	Name   string
	In     []byte
	Out    []byte
}
