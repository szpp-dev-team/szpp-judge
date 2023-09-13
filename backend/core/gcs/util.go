package gcs

import (
	"bytes"
	"context"
	"io"

	"cloud.google.com/go/storage"
)

func DownloadFile(ctx context.Context, bucketHandle *storage.BucketHandle, name string) ([]byte, error) {
	r, err := bucketHandle.Object(name).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

func UploadFile(ctx context.Context, bucketHandle *storage.BucketHandle, name string, b []byte) error {
	w := bucketHandle.Object(name).NewWriter(ctx)
	defer w.Close()
	if _, err := io.Copy(w, bytes.NewReader(b)); err != nil {
		return err
	}
	return nil
}
