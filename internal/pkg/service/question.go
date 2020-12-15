package service

import "feynman-podcast/internal/pkg/model/question"

func (c *Client) CreateQuestion(q *question.Question) (*string, error) {
	return c.QuestionCollection.CreateQuestion(q)
}

func (c *Client) GetQuestion(questionId string) *question.Question {
	return c.QuestionCollection.GetQuestion(questionId)
}
