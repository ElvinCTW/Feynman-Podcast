package service

import (
	"feynman-podcast/internal/pkg/config"
	"feynman-podcast/internal/pkg/crawler"
	"feynman-podcast/internal/pkg/model"
	"feynman-podcast/internal/pkg/upload"
	"net/http"
	"sync"
)

var (
	once   sync.Once
	client *Client
)

func NewClient(c *config.Config) *Client {
	once.Do(func() {
		httpClient := &http.Client{}
		client = &Client{
			ModelClient:   model.NewClient(c.Mongo.Database, c.Mongo.URI),
			UploadClient:  upload.NewClient(c.Upload.AwsAccessKey, c.Upload.AwsAccessSecret, c.Upload.AwsRegion, c.Upload.AwsBucket),
			CrawlerClient: crawler.NewClient(httpClient),
		}
	})

	return client
}

type Client struct {
	*model.ModelClient
	UploadClient  *upload.Client
	CrawlerClient *crawler.Client
}
