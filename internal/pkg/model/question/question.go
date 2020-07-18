package question

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

type Data struct {
	Id    *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Class struct {
		TestType string `json:"testType" bson:"testType"`
		Domain   string `json:"domain" bson:"domain"`
	} `json:"class" bson:"class"`
	Title        string    `json:"title" bson:"title"`
	Content      string    `json:"content" bson:"content"`
	Options      *[]Option `json:"options" bson:"options"`
	VoiceAnswers *[]string `json:"voiceAnswers" bson:"voiceAnswers"`
}

type Option struct {
	Num     string `json:"num" bson:"num"`
	Content string `json:"content" bson:"content"`
	Answer  bool   `json:"answer" bson:"answer"`
}
