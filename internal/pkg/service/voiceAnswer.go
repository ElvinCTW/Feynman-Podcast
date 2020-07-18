package service

import (
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
	} else if err = c.VoiceAnswerCollection.CreateData(questionId, userId, *uri); err != nil {
		return err
	} else {
		return nil
	}
}
