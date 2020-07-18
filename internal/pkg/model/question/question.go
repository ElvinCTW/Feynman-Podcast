package question

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	QuestionData = "questionData"
)

type QuestionCollection struct {
	col *mongo.Collection
}

func NewQuestionCollection(collection *mongo.Collection) *QuestionCollection {
	return &QuestionCollection{
		col: collection,
	}
}

func (c *QuestionCollection) CreateData(data *Data) error {
	if _, err := c.col.InsertOne(nil, data); err != nil {
		panic(err)
	}

	return nil
}

func (c *QuestionCollection) GetData(id string) *Data {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	filter := bson.M{"_id": objectId}

	data := new(Data)
	if err := c.col.FindOne(nil, filter).Decode(data); err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		panic(err)
	}

	return data
}

func (c *QuestionCollection) AddVoiceAnswer(id, VoiceAnswerId string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$push": bson.M{
			"voiceAnswers": bson.M{
				"$each": bson.A{VoiceAnswerId},
			},
		}}
	noUpsert := options.Update().SetUpsert(false)

	if _, err := c.col.UpdateOne(nil, filter, update, noUpsert); err != nil {
		panic(err)
	}

	return nil
}

type Data struct {
	Id    *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Class struct {
		TestType string `json:"testType" bson:"testType"`
		Domain   string `json:"domain" bson:"domain"`
	} `json:"class" bson:"class"`
	Title        string    `json:"title" bson:"title"`
	Content      string    `json:"content" bson:"content"`
	Options      *[]Option `json:"options,omitempty" bson:"options,omitempty"`
	VoiceAnswers *[]string `json:"voiceAnswers,omitempty" bson:"voiceAnswers,omitempty"`
}

type Option struct {
	Num     string `json:"num" bson:"num"`
	Content string `json:"content" bson:"content"`
	Answer  bool   `json:"answer" bson:"answer"`
}
