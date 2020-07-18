package service

import (
	"feynman-podcast/internal/pkg/config"
	"feynman-podcast/internal/pkg/model"
	"sync"
)

var (
	once   sync.Once
	client *Client
)

func NewClient(c *config.Config) *Client {
	once.Do(func() {
		client = &Client{}

		client.ModelClient = model.NewClient(c.Mongo.Database, c.Mongo.URI)
	})

	return client
}

type Client struct {
	*model.ModelClient
}
