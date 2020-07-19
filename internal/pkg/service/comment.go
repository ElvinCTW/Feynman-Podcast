package service

import (
	"feynman-podcast/internal/pkg/model/comment"
)

func (c *Client) CreateComment(answerId string, comment *comment.Data) error {
	return c.CommentCollection.CreateData(answerId, comment)
}
