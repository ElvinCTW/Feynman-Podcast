package upload

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
)

const (
	Question = "question"
)

type StorageClient struct {
	qBucket *storage.BucketHandle
}

func NewStorage(app *firebase.App) *StorageClient {
	client, err := app.Storage(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	qBucket, err := client.Bucket(Question)
	if err != nil {
		fmt.Println(err)
	}
	return &StorageClient{
		qBucket: qBucket,
	}
}
