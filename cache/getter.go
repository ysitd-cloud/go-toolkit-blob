package cache

import (
	"context"
	"io/ioutil"

	"github.com/golang/groupcache"
	"github.com/minio/minio-go"
)

// Logger interface of getter logging
type Logger interface {
	Debugf(format string, vals ...interface{})
}

type getter struct {
	client *minio.Client
	bucket string
	logger Logger
}

// NewGetter create a groupgcache.Getter for CachedBlobStore
func NewGetter(client *minio.Client, bucket string, logger Logger) groupcache.Getter {
	return &getter{
		client: client,
		bucket: bucket,
		logger: logger,
	}
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
