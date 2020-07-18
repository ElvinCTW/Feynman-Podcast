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

	va := &VoiceAnswer{
		QuestionId: &questionObjectId,
		UserId:     &userObjectId,
		URI:        uri,
	}

	r, err := c.col.InsertOne(nil, va)
	if err != nil {
		panic(err)
	}

	insertedId := r.InsertedID.(primitive.ObjectID).Hex()

	return &insertedId, nil
}

func (c *VoiceAnswerCollection) ListData(questionId string) []*VoiceAnswer {
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

	var list []*VoiceAnswer
	for cur.TryNext(nil) {
		va := new(VoiceAnswer)
		if err := cur.Decode(va); err != nil {
			continue
		}
		list = append(list, va)
	}

	return list
}

type VoiceAnswer struct {
	Id         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	QuestionId *primitive.ObjectID `json:"questionId,omitempty" bson:"questionId,omitempty"`
	UserId     *primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty" `
	URI        string              `json:"uri" bson:"uri"`
	Comments   *[]Comment          `json:"comments,omitempty" bson:"comments,omitempty"`
	Likers     *[]string           `json:"likers,omitempty" bson:"likers,omitempty"`
	LikeCount  int                 `json:"likeCount" bson:"likeCount"`
}
