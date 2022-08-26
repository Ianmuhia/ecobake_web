package internal

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// NewMinioConnection MinioConnection func for opening file storage connection.
func NewMinioConnection(
	accessKeyID string,
	accessKey string,
	endpoint string,
	bucketName string,
	logger *log.Logger) (*minio.Client, string, error) {
	ctx := context.Background()

	// Initialize file storage client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, accessKey, ""),
		Secure: false,
	})
	if err != nil {
		logger.Fatalln(err)
	}
	// Make a new bucket called dev-file storage.

	location := "us-east-1"

	createBktError := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if createBktError != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			logger.Printf("Bucket %s is already created ...Skipping creation", bucketName)
		} else {
			logger.Fatalln(err)
		}
	} else {
		logger.Printf("Successfully created %s\n", bucketName)
	}

	return minioClient, bucketName, err
}
