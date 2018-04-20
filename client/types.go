package client

import (
	"context"
	"io"
)

type BlobStore interface {
	GetObject(bucket, key string) ([]byte, error)
	GetObjectWithContext(ctx context.Context, bucket, key string) ([]byte, error)

	PutObject(bucket, key string, body io.Reader) error
	PutObjectWithContext(ctx context.Context, bucket, key string, body io.Reader) error

	DeleteObject(bucket, key string) error
	DeleteObjectWithContext(ctx context.Context, bucket, key string) error
}
