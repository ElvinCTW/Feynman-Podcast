package upload

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	once   sync.Once
	client *Client
)

const (
	audioMpeg      = "audio/mpeg"
	contentType    = "Content-Type"
	filePathFormat = "https://%s.s3-%s.amazonaws.com/%s"
)

type Client struct {
	bucket      string
	region      string
	minioClient *minio.Client
}

func NewClient(accessKey, accessSecret, region, bucket string) *Client {
	once.Do(func() {
		minioClient, err := minio.New("s3.amazonaws.com", &minio.Options{
			Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
			Secure: false,
		})
		if err != nil {
			log.Fatalln(err)
		}

		client = &Client{
			minioClient: minioClient,
			bucket:      bucket,
			region:      region,
		}
	})

	return client
}

func (c *Client) UploadMp3(files []*multipart.FileHeader, name string) (*string, error) {
	if t := files[0].Header.Get(contentType); t != audioMpeg {
		return nil, errors.New("not audio/mpeg type file")
	}

	f, err := files[0].Open()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	option := minio.PutObjectOptions{ContentType: audioMpeg}
	if _, err := c.minioClient.PutObject(context.Background(), c.bucket, name, f, files[0].Size, option); err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(filePathFormat, c.bucket, c.region, name)

	return &uri, nil
}

func (c *Client) UploadText(r io.Reader, name string) (*string, error) {
	option := minio.PutObjectOptions{ContentType: "text/plain"}
	if _, err := c.minioClient.PutObject(context.Background(), c.bucket, name, r, -1, option); err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(filePathFormat, c.bucket, c.region, name)

	return &uri, nil
}
