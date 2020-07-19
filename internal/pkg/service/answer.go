package service

import (
	"feynman-podcast/internal/pkg/fperr"
	"feynman-podcast/internal/pkg/model/answer"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"
)

func (c *Client) CreateAnswer(questionId string, userId string, file []*multipart.FileHeader) error {
	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	name := fmt.Sprintf("%s-%s-%s", questionId, userId, now)

	if uri, err := c.UploadClient.UploadMp3(file, name); err != nil {
		return err
	} else if _, err := c.AnswerCollection.CreateData(questionId, userId, *uri); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Client) GetAnswer(id string) *answer.Data {
	return c.AnswerCollection.GetData(id)
}

func (c *Client) ListAnswer(questionId string) *[]answer.Data {
	return c.AnswerCollection.ListData(questionId)
}

func (c *Client) UpdateAnswerLike(id, likerId string) error {
	if va := c.GetAnswer(id); va == nil {
		return fperr.New(fperr.NoContent)
	}
	return c.AnswerCollection.Updatelike(id, likerId)
}
