package testcases

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

const tasksFolderID = "1d4TxRKqttep964mM-k60lsNjxhVVY8HI"

type gcsImpl struct {
	client *storage.Client
}

func NewRepository(client *storage.Client) Repository {
	return &gcsImpl{client}
}

func (i *gcsImpl) DownloadTestcases(ctx context.Context, taskID int) ([]*Testcase, error) {
	bucket := i.client.Bucket("szpp-judge")
	itr := bucket.Objects(ctx, &storage.Query{
		Prefix: fmt.Sprintf("/testcases/%d/in", taskID),
	})

	testcases := make([]*Testcase, 0)
	for {
		attrs, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		testcaseName := filepath.Base(attrs.Name)
		inBytes, err := download(ctx, bucket, fmt.Sprintf("/testcases/%d/in/%s", taskID, testcaseName))
		if err != nil {
			return nil, err
		}
		outBytes, err := download(ctx, bucket, fmt.Sprintf("/testcases/%d/out/%s", taskID, testcaseName))
		if err != nil {
			return nil, err
		}
		testcases = append(testcases, &Testcase{
			TaskID: taskID,
			Name:   testcaseName,
			In:     inBytes,
			Out:    outBytes,
		})
	}
	return testcases, nil
}

func download(ctx context.Context, bucketHandle *storage.BucketHandle, name string) ([]byte, error) {
	r, err := bucketHandle.Object(name).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}
