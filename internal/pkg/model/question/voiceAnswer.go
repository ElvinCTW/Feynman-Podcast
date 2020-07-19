package question

import (
	"go.mongodb.org/mongo-driver/bson"
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

func (c *VoiceAnswerCollection) CreateData(questionId, userId, uri string) (*string, error) {
	questionObjectId, err := primitive.ObjectIDFromHex(questionId)
	if err != nil {
		return nil, err
	}

	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	comment := make([]Comment, 0)
	likers := make([]string, 0)
	va := &VoiceAnswer{
		QuestionId: &questionObjectId,
		UserId:     &userObjectId,
		URI:        uri,
		Comments:   &comment,
		Likers:     &likers,
	}

	r, err := c.col.InsertOne(nil, va)
	if err != nil {
		panic(err)
	}

	insertedId := r.InsertedID.(primitive.ObjectID).Hex()

	return &insertedId, nil
}

func (c *VoiceAnswerCollection) ListData(questionId string) *[]VoiceAnswer {
	questionObjectId, err := primitive.ObjectIDFromHex(questionId)
	if err != nil {
		return nil
	}

	filter := bson.M{"questionId": questionObjectId}
	cur, err := c.col.Find(nil, filter)
	if err != nil {
		panic(err)
	}

	defer cur.Close(nil)

	list := make([]VoiceAnswer, 0)
	if err = cur.All(nil, &list); err != nil {
		panic(err)
	}

	return &list
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
