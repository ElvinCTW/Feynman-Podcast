package service

import (
	"context"
	"feynman-podcast/internal/pkg/config"
	"feynman-podcast/internal/pkg/crawler"
	"feynman-podcast/internal/pkg/model"
	"feynman-podcast/internal/pkg/upload"
	"net/http"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/option"
	"sync"
)

var (
	once   sync.Once
	client *Client
)

func NewClient(c *config.Config) *Client {
	once.Do(func() {
		httpClient := &http.Client{}
		opt := option.WithCredentialsFile("test-helper-key.json")
		config := &firebase.Config{
			StorageBucket: "gs://sapient-torch-298615.appspot.com",
		}

		app, err := firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			fmt.Println("unable to init app engine")
		}

		fmt.Println(app)

		client = &Client{
			//ModelClient:     model.NewClient(c.Mongo.Database, c.Mongo.URI),
			FireStoreClient: model.NewFireStore(app),
			UploadClient:    upload.NewStorage(app),
			CrawlerClient: crawler.NewClient(httpClient),
			App:          app,
		}
	})

	return client
}

type Client struct {
	//*model.ModelClient
	*model.FireStoreClient
	UploadClient *upload.StorageClient
	CrawlerClient *crawler.Client
	App          *firebase.App
}
