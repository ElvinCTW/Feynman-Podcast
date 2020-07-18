package question

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	VoiceAnswerData = "voiceAnswerData"
)

type VoiceAnswerCollection struct {
	col *mongo.Collection
}

func NewVoiceAnswerCollection(collection *mongo.Collection) *VoiceAnswerCollection {
	return &VoiceAnswerCollection{
		col: collection,
	}
}

func (c *VoiceAnswerCollection) CreateData(questionId, userId, uri string) error {
	questionObjectId, err := primitive.ObjectIDFromHex(questionId)
	if err != nil {
		return err
	}

	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	va := &VoiceAnswer{
		QuestionId: &questionObjectId,
		UserId:     &userObjectId,
		URI:        uri,
	}

	if _, err := c.col.InsertOne(nil, va); err != nil {
		panic(err)
	}

	return nil
}

type VoiceAnswer struct {
	Id         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	QuestionId *primitive.ObjectID `json:"questionId" bson:"questionId"`
	UserId     *primitive.ObjectID `json:"userId" bson:"userId" `
	URI        string              `json:"uri" bson:"uri"`
	Comments   *[]Comment          `json:"comments" bson:"comments"`
	Likers     *[]string           `json:"likers" bson:"likers"`
	LikeCount  int                 `json:"likeCount" bson:"likeCount"`
}
