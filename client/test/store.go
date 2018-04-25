package test

import (
	"code.ysitd.cloud/toolkit/blob/client"
	"context"
	"errors"
	"io"
	"io/ioutil"
)

var ErrNotFound = errors.New("not found")

type MockBlobStore struct {
	client.BlobStore
	Blob map[string][]byte
}

func makeKey(bucket, key string) string {
	return bucket + "@" + key
}

func (s *MockBlobStore) GetObjectWithContext(_ context.Context, bucket, key string) (data []byte, err error) {
	data, exists := s.Blob[makeKey(bucket, key)]
	if !exists {
		return nil, ErrNotFound
	}
	return data, nil
}

func (s *MockBlobStore) PutObjectWithContext(_ context.Context, bucket, key string, body io.Reader) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	s.Blob[makeKey(bucket, key)] = data
	return nil
}

func (s *MockBlobStore) DeleteObjectWithContext(_ context.Context, bucket, key string) error {
	delete(s.Blob, makeKey(bucket, key))
	return nil
}
