package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go"
	"io"
	"io/ioutil"
)

type Store struct {
	Client *minio.Client
}

func (s *Store) GetObject(bucket, key string) ([]byte, error) {
	return s.GetObjectWithContext(context.Background(), bucket, key)
}

func (s *Store) GetObjectWithContext(ctx context.Context, bucket, key string) (data []byte, err error) {
	obj, err := s.Client.GetObjectWithContext(ctx, bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return
	}
	defer obj.Close()

	return ioutil.ReadAll(obj)
}

func (s *Store) PutObject(bucket, key string, body io.Reader) error {
	return s.PutObjectWithContext(context.Background(), bucket, key, body)
}

func (s *Store) PutObjectWithContext(ctx context.Context, bucket, key string, body io.Reader) error {
	var buffer bytes.Buffer
	buffer.ReadFrom(body)
	_, err := s.Client.PutObjectWithContext(ctx, bucket, key, &buffer, int64(buffer.Len()), minio.PutObjectOptions{})
	return err
}

func (s *Store) DeleteObject(bucket, key string) error {
	return s.DeleteObjectWithContext(context.Background(), bucket, key)
}

func (s *Store) DeleteObjectWithContext(ctx context.Context, bucket, key string) error {
	objs := make(chan string)
	defer close(objs)
	errs := s.Client.RemoveObjectsWithContext(ctx, bucket, objs)
	objs <- key
	err := <-errs
	return err.Err
}
