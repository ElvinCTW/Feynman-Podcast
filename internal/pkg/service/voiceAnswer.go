package service

import (
	"feynman-podcast/internal/pkg/fperr"
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

func (c *Client) GetVoiceAnswer(id string) *question.VoiceAnswer {
	return c.VoiceAnswerCollection.GetData(id)
}

func (c *Client) ListVoiceAnswer(questionId string) *[]question.VoiceAnswer {
	return c.VoiceAnswerCollection.ListData(questionId)
}

func (c *Client) UpdateVoiceAnswerLike(id, likerId string) error {
	if va := c.GetVoiceAnswer(id); va == nil {
		return fperr.New(fperr.NoContent)
	}
	return c.VoiceAnswerCollection.Updatelike(id, likerId)
}
