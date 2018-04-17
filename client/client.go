// Copyright 2018 Tony Yip. All rights reserved.
// Use of this source code is governed by MIT license.

package client

import (
	"os"

	"github.com/minio/minio-go"
)

// NewFromEnv create minio client from environment variables
func NewFromEnv() (*minio.Client, error) {
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
