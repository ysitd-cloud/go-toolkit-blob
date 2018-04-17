package cache

import (
	"context"

	"github.com/golang/groupcache"
)

// CachedBlobStore is the blob store for providing cached blob storage
type CachedBlobStore struct {
	Group *groupcache.Group
}

// GetBlob get blob file without context
func (s *CachedBlobStore) GetBlob(key string) (dest []byte, err error) {
	return s.GetBlobWithContext(context.Background(), key)
}

// GetBlobWithContext blob file from either cache or S3
func (s *CachedBlobStore) GetBlobWithContext(ctx context.Context, key string) (dest []byte, err error) {
	err = s.Group.Get(ctx, key, groupcache.AllocatingByteSliceSink(&dest))
	return
}
