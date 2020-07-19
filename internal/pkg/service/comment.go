package service

import (
	"feynman-podcast/internal/pkg/model/comment"
)

func (c *Client) CreateComment(answerId string, userId string, comment *comment.Data) error {
	return c.CommentCollection.CreateData(answerId, userId, comment)
}

func (c *Client) ListComment(awswerId string) *[]comment.Data {
	return c.CommentCollection.ListData(awswerId)
}

func (c *Client) UpdateCommentLike(id, likerId string) error {
	return c.CommentCollection.Updatelike(id, likerId)
}

func (c *Client) DeleteComment(id, userId string) {
	c.CommentCollection.DeleteData(id, userId)
}
