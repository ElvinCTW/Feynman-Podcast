package service

import (
	"sync"
	"feynman-podcast/internal/pkg/config"
)

type Client struct {
}

var (
	once   sync.Once
	client *Client
)

func NewClient(c *config.Config) *Client {
	once.Do(func() {
		client = &Client{}
	})

	return client
}
