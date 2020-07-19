package service

import (
	"feynman-podcast/internal/pkg/model/comment"
)

func (c *Client) CreateComment(answerId string, comment *comment.Data) error {
	return c.CommentCollection.CreateData(answerId, comment)
}

func (c *Client) ListComment(awswerId string) *[]comment.Data {
	// todo
	return nil
}

func (c *Client) UpdateCommentLike(id, likerId string) error {
	// todo
	return nil
}

func (c *Client) DeleteComment(id string) {
	// todo
}
