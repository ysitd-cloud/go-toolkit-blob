// Copyright 2018 Tony Yip. All rights reserved.
// Use of this source code is governed by MIT license.

package cache

import (
	"context"

	"code.ysitd.cloud/toolkit/blob/client"

	"github.com/golang/groupcache"
)

type getter struct {
	client client.BlobStore
	bucket string
	logger Logger
}

// NewGetter create a groupgcache.Getter for CachedBlobStore
func NewGetter(client client.BlobStore, bucket string, logger Logger) groupcache.Getter {
	if logger == nil {
		logger = &NullLogger{}
	}
	return &getter{
		client: client,
		bucket: bucket,
		logger: logger,
	}
}

func (g *getter) Get(gctx groupcache.Context, key string, dest groupcache.Sink) (err error) {
	ctx := gctx.(context.Context)

	g.logger.Debugf("Get %s from s3", key)

	buffer, err := g.client.GetObjectWithContext(ctx, g.bucket, key)
	if err != nil {
		return
	}

	return dest.SetBytes(buffer)
}
