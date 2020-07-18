package question

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct {
	Content   string              `json:"content" bson:"content"`
	UserId    *primitive.ObjectID `json:"userId" bson:"userId" `
	Likers    *[]string           `json:"likers" bson:"likers"`
	LikeCount int                 `json:"likeCount" bson:"likeCount"`
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
