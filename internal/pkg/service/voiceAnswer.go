package service

import "feynman-podcast/internal/pkg/model/question"

func (c *Client) CreateVoiceAnswer(va *question.VoiceAnswer) error {
	// todo upload file to s3 and get uri
	return c.VoiceAnswerCollection.CreateData(va)
}

func (c *Client) CreateComment(voiceAnswerId string, comment *question.Comment) error {
	return c.VoiceAnswerCollection.CreateComment(voiceAnswerId, comment)
}
