package testcases

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

type gcsImpl struct {
	client *storage.Client
}

func NewRepository(client *storage.Client) Repository {
	return &gcsImpl{client}
}

func (i *gcsImpl) DownloadTestcase(ctx context.Context, taskID int, name string) (*Testcase, error) {
	bucket := i.client.Bucket("szpp-judge")
	inBytes, err := download(ctx, bucket, fmt.Sprintf("/testcases/%d/in/%s", taskID, name))
	if err != nil {
		return nil, err
	}
	outBytes, err := download(ctx, bucket, fmt.Sprintf("/testcases/%d/out/%s", taskID, name))
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
	if err := upload(ctx, bucket, fmt.Sprintf("/testcases/%d/in/%s", taskID, testcase.Name), testcase.In); err != nil {
		return err
	}
	if err := upload(ctx, bucket, fmt.Sprintf("/testcases/%d/out/%s", taskID, testcase.Name), testcase.Out); err != nil {
		return err
	}
	return nil
}

func download(ctx context.Context, bucketHandle *storage.BucketHandle, name string) ([]byte, error) {
	r, err := bucketHandle.Object(name).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

func upload(ctx context.Context, bucketHandle *storage.BucketHandle, name string, b []byte) error {
	w := bucketHandle.Object(name).NewWriter(ctx)
	defer w.Close()
	if _, err := io.Copy(w, bytes.NewReader(b)); err != nil {
		return err
	}
	return nil
}

var _ Repository = (*gcsImpl)(nil)
