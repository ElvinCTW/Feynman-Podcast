package service

import "feynman-podcast/internal/pkg/model/question"

func (c *Client) CreateComment(voiceAnswerId string, comment *question.Comment) error {
	return c.VoiceAnswerCollection.CreateComment(voiceAnswerId, comment)
}
