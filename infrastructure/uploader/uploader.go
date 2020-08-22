package uploader

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"image"
	"image/png"
	"net/url"
	"os"
)

type Uploader interface {
	UploadFromBytes(bucketName string, filePath string, byte []byte) (string, error)
	UploadFromImage(bucketName string, filePath string, imageData image.Image) (string, error)
}

func NewStorageUploader() (Uploader, error) {
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(os.Getenv("STORAGE_SERVICE")))
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		return nil, err
	}
	return NewGCSUploader(client), nil
}

type GCSUploader struct {
	*storage.Client
}

func NewGCSUploader(client *storage.Client) *GCSUploader {
	return &GCSUploader{client}
}

func (u GCSUploader) UploadFromImage(bucketName string, filePath string, imageData image.Image) (string, error) {
	buf := &bytes.Buffer{}
	if err := png.Encode(buf, imageData); err != nil {
		return "", err
	}
	return u.UploadFromBytes(bucketName, filePath, buf.Bytes())
}

func (u GCSUploader) UploadFromBytes(bucketName string, filePath string, b []byte) (string, error) {
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(os.Getenv("STORAGE_SERVICE")))
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		return "", err
	}

	writer := client.Bucket(bucketName).Object(filePath).NewWriter(context.Background())
	_, err = writer.Write(b)
	if err != nil {
		return "", err
	}
	writer.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	writer.ContentType = fmt.Sprintf("image/png")
	writer.CacheControl = "public, max-age=86400"
	if err := writer.Close(); err != nil {
		return "", err
	}
	url := url.URL{Path: fmt.Sprintf("%s/%s", bucketName, filePath)}
	return url.String(), nil
}
