package question

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (c *VoiceAnswerCollection) CreateData(va *VoiceAnswer) error {
	if _, err := c.col.InsertOne(nil, va); err != nil {
		panic(err)
	}

	return nil
}

func (c *VoiceAnswerCollection) CreateComment(voiceAnswerId string, comment *Comment) error {
	vaObjectId, err := primitive.ObjectIDFromHex(voiceAnswerId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": vaObjectId}
	update := bson.M{
		"$push": bson.M{
			"comments": bson.M{
				"$each": bson.A{comment},
			},
		}}
	noUpsert := options.Update().SetUpsert(false)

	if _, err := c.col.UpdateOne(nil, filter, update, noUpsert); err != nil {
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

type Comment struct {
	Content   string              `json:"content" bson:"content"`
	UserId    *primitive.ObjectID `json:"userId" bson:"userId" `
	Likers    *[]string           `json:"likers" bson:"likers"`
	LikeCount int                 `json:"likeCount" bson:"likeCount"`
}