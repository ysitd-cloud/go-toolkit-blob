package blob

import (
	"github.com/minio/minio-go"
)

func RemovePrefix(client *minio.Client, bucketName, objectPrefix string, recursive bool) []error {
	objectsCh := make(chan string)

	errs := make([]error, 0)

	// Send object names that are needed to be removed to objectsCh
	go func() {
		defer close(objectsCh)
		// List all objects from a bucket-name with a matching prefix.
		for object := range client.ListObjects(bucketName, objectPrefix, true, nil) {
			if object.Err != nil {
				errs = append(errs, object.Err)
			} else {
				objectsCh <- object.Key
			}
		}
	}()
	for rErr := range client.RemoveObjects(bucketName, objectsCh) {
		errs = append(errs, rErr.Err)
	}
	return errs
}
