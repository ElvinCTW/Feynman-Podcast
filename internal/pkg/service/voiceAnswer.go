package service

import (
	"feynman-podcast/internal/pkg/model/question"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"
)

func (c *Client) CreateVoiceAnswer(questionId string, userId string, file []*multipart.FileHeader) error {
	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	name := fmt.Sprintf("%s-%s-%s", questionId, userId, now)

	if uri, err := c.UploadClient.UploadMp3(file, name); err != nil {
		return err
	} else if insertedId, err := c.VoiceAnswerCollection.CreateData(questionId, userId, *uri); err != nil {
		return err
	} else if err = c.AddQuestionVoiceAnswer(questionId, *insertedId); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Client) ListVoiceAnswer(questionId string) *[]question.VoiceAnswer {
	return c.VoiceAnswerCollection.ListData(questionId)
}
