package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/minio/minio-go/v7"
)

type fileStorageService struct {
	bucketName string
	client     *minio.Client
}

func NewFileStorageService(bucketName string, client *minio.Client) *fileStorageService {
	return &fileStorageService{bucketName: bucketName, client: client}
}

type FileStorageService interface {
	UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error)
	MultipleFileUpload(filenames []*multipart.FileHeader) ([]string, error)
}

func (s *fileStorageService) UploadFile(objectName string, fileBuffer multipart.File, fileSize int64, contentType string) (minio.UploadInfo, error) {
	ctx := context.Background()
	// Upload the zip file with PutObject
	info, err := s.client.PutObject(
		ctx,
		s.bucketName,
		objectName,
		fileBuffer,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return info, err
	}
	return info, nil
}

func (s *fileStorageService) MultipleFileUpload(filenames []*multipart.FileHeader) ([]string, error) {
	type item struct {
		thumbnail minio.UploadInfo
		err       error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(fh *multipart.FileHeader) {
			ctx := context.Background()
			ctt := fh.Header["Content-Type"][0]
			v, _ := fh.Open()
			var it item
			fname := fmt.Sprintf("%s/%s", "testfolder", fh.Filename)

			it.thumbnail, it.err = s.client.PutObject(ctx, s.bucketName, fname, v, fh.Size, minio.PutObjectOptions{ContentType: ctt})

			ch <- it
		}(f)
	}
	dt := make([]string, 0)
	for range filenames {
		it := <-ch
		if it.err != nil {
			return dt, it.err
		}
		dt = append(dt, it.thumbnail.Key)
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := s.client.ListObjects(ctx, "fainda", minio.ListObjectsOptions{
		WithVersions: false,
		WithMetadata: false,
		Prefix:       "testfolder",
		Recursive:    true,
		MaxKeys:      0,
		StartAfter:   "",
		UseV1:        false,
	})
	for object := range objectCh {
		if object.Err != nil {
			//fmt.Println(object.Err)
			//return
		}
		g := strings.Split(object.Key, "/")
		fmt.Println(g[1])
		newfolder := fmt.Sprintf("%s/%s", "testfolder2", g[1])

		srcOpts := minio.CopySrcOptions{
			Bucket: "fainda",
			Object: object.Key,
		}

		// Destination object
		dstOpts := minio.CopyDestOptions{
			Bucket: "fainda",
			Object: newfolder,
		}

		// Copy object call
		uploadInfo, err := s.client.CopyObject(context.Background(), dstOpts, srcOpts)
		if err != nil {
			fmt.Println(err)
			//return
		}

		fmt.Println("Successfully copied object:", uploadInfo)

	}

	return dt, nil
}
