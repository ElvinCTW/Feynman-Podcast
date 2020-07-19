package comment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CommentData = "commentData"
)

type Data struct {
	Id        *primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	AnswerId  *primitive.ObjectID `json:"answerId,omitempty" bson:"answerId,omitempty"`
	Content   string              `json:"content" bson:"content"`
	UserId    *primitive.ObjectID `json:"userId" bson:"userId"`
	Likers    *[]string           `json:"likers" bson:"likers"`
	LikeCount int                 `json:"likeCount" bson:"likeCount"`
}

type DataCollection struct {
	col *mongo.Collection
}

func NewDataCollection(collection *mongo.Collection) *DataCollection {
	return &DataCollection{
		col: collection,
	}
}

func (c *DataCollection) CreateData(answerId string, data *Data) error {
	answerObjectId, err := primitive.ObjectIDFromHex(answerId)
	if err != nil {
		return err
	}

	likers := make([]string, 0)
	d := &Data{
		AnswerId:  &answerObjectId,
		Content:   data.Content,
		UserId:    data.UserId,
		Likers:    &likers,
		LikeCount: 0,
	}

	if _, err := c.col.InsertOne(nil, d); err != nil {
		panic(err)
	}

	return nil
}
