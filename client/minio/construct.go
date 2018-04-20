// Copyright 2018 Tony Yip. All rights reserved.
// Use of this source code is governed by MIT license.

package minio

import (
	"os"

	"github.com/minio/minio-go"
)

// NewMinioClientFromEnv create minio client from environment variables
func NewMinioClientFromEnv() (*minio.Client, error) {
	endpoint := os.Getenv("S3_ENDPOINT")
	if endpoint == "" {
		endpoint = "s3.amazonaws.com"
	}

	return minio.New(
		endpoint,
		os.Getenv("S3_ACCESS_KEY_ID"),
		os.Getenv("S3_SECRET_ACCESS_KEY"),
		os.Getenv("S3_INSECURE") == "",
	)
}
