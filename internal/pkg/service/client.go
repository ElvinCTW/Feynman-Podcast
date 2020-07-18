package service

import (
	"feynman-podcast/internal/pkg/config"
	"feynman-podcast/internal/pkg/model"
	"feynman-podcast/internal/pkg/upload"
	"sync"
)

var (
	once   sync.Once
	client *Client
)

func NewClient(c *config.Config) *Client {
	once.Do(func() {
		client = &Client{
			ModelClient:  model.NewClient(c.Mongo.Database, c.Mongo.URI),
			UploadClient: upload.NewClient(c.Upload.AwsAccessKey, c.Upload.AwsAccessSecret, c.Upload.AwsRegion, c.Upload.AwsBucket),
		}
	})

	return client
}

type Client struct {
	*model.ModelClient
	UploadClient *upload.Client
}
