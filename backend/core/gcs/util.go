package gcs

import (
	"bytes"
	"context"
	"io"

	"cloud.google.com/go/storage"
)

func DownloadFile(ctx context.Context, bucketHandle *storage.BucketHandle, name string) (b []byte, err error) {
	var r *storage.Reader
	r, err = bucketHandle.Object(name).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = r.Close(); err != nil {
			b = nil
		}
	}()
	return io.ReadAll(r)
}

func UploadFile(ctx context.Context, bucketHandle *storage.BucketHandle, name string, b []byte) (err error) {
	w := bucketHandle.Object(name).NewWriter(ctx)
	defer func() {
		err = w.Close()
	}()
	if _, err := io.Copy(w, bytes.NewReader(b)); err != nil {
		return err
	}
	return nil
}
