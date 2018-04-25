package cache

import (
	"code.ysitd.cloud/toolkit/blob/client"
	"code.ysitd.cloud/toolkit/blob/client/test"
)

func newMockClient() client.BlobStore {
	return &test.MockBlobStore{
		Blob: make(map[string][]byte),
	}
}
