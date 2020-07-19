package answer

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	AnswerData = "answerData"
)

type DataCollection struct {
	col *mongo.Collection
}

func NewDataCollection(collection *mongo.Collection) *DataCollection {
	return &DataCollection{
		col: collection,
	}
}

func (c *DataCollection) CreateData(questionId, userId, uri string) (*string, error) {
	questionObjectId, err := primitive.ObjectIDFromHex(questionId)
	if err != nil {
		return nil, err
	}

	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	likers := make([]string, 0)
	va := &Data{
		QuestionId: &questionObjectId,
		UserId:     &userObjectId,
		URI:        uri,
		Likers:     &likers,
	}

	r, err := c.col.InsertOne(nil, va)
	if err != nil {
		panic(err)
	}

	insertedId := r.InsertedID.(primitive.ObjectID).Hex()

	return &insertedId, nil
}

func (c *DataCollection) GetData(id string) *Data {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	data := new(Data)
	filter := bson.M{"_id": objectId}
	if err = c.col.FindOne(nil, filter).Decode(data); err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		panic(err)
	}

	return data
}

func (c *DataCollection) ListData(questionId string) *[]Data {
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

	list := make([]Data, 0)
	if err = cur.All(nil, &list); err != nil {
		panic(err)
	}

	return &list
}

// call this function after confirm answer exist with get Data()
func (c *DataCollection) Updatelike(id, likerId string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// check user in likers or not
	filter := bson.M{"_id": objectId, "likers": bson.M{"$in": bson.A{likerId}}}
	r := c.col.FindOne(nil, filter, options.FindOne().SetProjection(bson.M{"_id": 1}))
	if r.Err() != nil && r.Err() != mongo.ErrNoDocuments {
		panic(err)
	}

	// update likers and like
	filter = bson.M{"_id": objectId}
	noUpsert := options.Update().SetUpsert(false)
	update := bson.M{"$inc": bson.M{"likeCount": -1}, "$pull": bson.M{"likers": likerId}}
	if r.Err() == mongo.ErrNoDocuments {
		update = bson.M{"$inc": bson.M{"likeCount": 1}, "$addToSet": bson.M{"likers": likerId}}
	}

	if _, err := c.col.UpdateOne(nil, filter, update, noUpsert); err != nil {
		panic(err)
	}

	return nil
}

type Data struct {
	Id         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	QuestionId *primitive.ObjectID `json:"questionId" bson:"questionId"`
	UserId     *primitive.ObjectID `json:"userId" bson:"userId" `
	URI        string              `json:"uri" bson:"uri"`
	Likers     *[]string           `json:"likers" bson:"likers"`
	LikeCount  int                 `json:"likeCount" bson:"likeCount"`
}
