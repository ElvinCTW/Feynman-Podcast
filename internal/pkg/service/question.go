package service

import "feynman-podcast/internal/pkg/model/question"

func (c *Client) CreateQuestion(question *question.Data) error {
	return c.QuestionCollection.CreateData(question)
}

func (c *Client) GetQuestion(questionId string) *question.Data {
	return c.QuestionCollection.GetData(questionId)
}
