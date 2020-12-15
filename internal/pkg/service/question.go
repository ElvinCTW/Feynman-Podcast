package service

import "feynman-podcast/internal/pkg/model/question"

func (c *Client) CreateQuestion(q *question.Question) error {
	return c.QuestionCollection.CreateQuestion(q.Class.Domain, q.Title, q.Content)
}

//func (c *Client) GetQuestion(questionId string) *question.Data {
//	return c.QuestionCollection.GetData(questionId)
//}
