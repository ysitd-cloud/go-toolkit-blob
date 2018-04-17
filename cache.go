package blob

import (
	"context"
	"io/ioutil"

	"github.com/golang/groupcache"
	"github.com/minio/minio-go"
)

type Logger interface {
	Debugf(format string, vals ...interface{})
}

type CachedBlobStore struct {
	Group *groupcache.Group
}

func New(client *minio.Client, bucket string, name string, size int64, logger Logger) *CachedBlobStore {
	return &CachedBlobStore{
		Group: groupcache.NewGroup(name, size, &getter{
			client: client,
			bucket: bucket,
			logger: logger,
		}),
	}
}

func (s *CachedBlobStore) Get(key string) (dest []byte, err error) {
	return s.GetWithContext(context.Background(), key)
}

func (s *CachedBlobStore) GetWithContext(ctx context.Context, key string) (dest []byte, err error) {
	err = s.Group.Get(ctx, key, groupcache.AllocatingByteSliceSink(&dest))
	return
}

type getter struct {
	client *minio.Client
	bucket string
	logger Logger
}

func (g *getter) Get(gctx groupcache.Context, key string, dest groupcache.Sink) (err error) {
	ctx := gctx.(context.Context)

	g.logger.Debugf("Get %s from s3", key)

	obj, err := g.client.GetObjectWithContext(ctx, g.bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return
	}

	defer obj.Close()

	buffer, err := ioutil.ReadAll(obj)

	if err != nil {
		return
	}

	return dest.SetBytes(buffer)
}
