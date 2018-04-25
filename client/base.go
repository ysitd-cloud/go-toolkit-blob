package client

import (
	"context"
	"errors"
	"io"
)

var ErrNotImplements = errors.New("not impleaments")

type BaseBlobStore struct{}

func (s *BaseBlobStore) GetObject(bucket, key string) ([]byte, error) {
	return s.GetObjectWithContext(context.Background(), bucket, key)
}

func (s *BaseBlobStore) GetObjectWithContext(ctx context.Context, bucket, key string) (data []byte, err error) {
	return nil, ErrNotImplements
}

func (s *BaseBlobStore) PutObject(bucket, key string, body io.Reader) error {
	return s.PutObjectWithContext(context.Background(), bucket, key, body)
}

func (s *BaseBlobStore) PutObjectWithContext(ctx context.Context, bucket, key string, body io.Reader) error {
	return ErrNotImplements
}

func (s *BaseBlobStore) DeleteObject(bucket, key string) error {
	return s.DeleteObjectWithContext(context.Background(), bucket, key)
}

func (s *BaseBlobStore) DeleteObjectWithContext(ctx context.Context, bucket, key string) error {
	return ErrNotImplements
}
