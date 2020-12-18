package question

//
//import (
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//)
//
//const (
//	QuestionData = "questionData"
//)
//
//type DataCollection struct {
//	col *mongo.Collection
//}
//
//func NewDataCollection(collection *mongo.Collection) *DataCollection {
//	return &DataCollection{
//		col: collection,
//	}
//}
//
//func (c *DataCollection) CreateData(data *Data) error {
//	d := &Data{
//		Class:   data.Class,
//		Title:   data.Title,
//		Content: data.Content,
//		Options: data.Options,
//	}
//
//	if _, err := c.col.InsertOne(nil, d); err != nil {
//		panic(err)
//	}
//
//	return nil
//}
//
//func (c *DataCollection) GetData(id string) *Data {
//	objectId, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return nil
//	}
//
//	filter := bson.M{"_id": objectId}
//
//	data := new(Data)
//	if err := c.col.FindOne(nil, filter).Decode(data); err == mongo.ErrNoDocuments {
//		return nil
//	} else if err != nil {
//		panic(err)
//	}
//
//	return data
//}
//
//type Data struct {
//	Id    *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
//	Class struct {
//		TestType string `json:"testType" bson:"testType"`
//		Domain   string `json:"domain" bson:"domain"`
//	} `json:"class,inline" bson:"class,inline"`
//	Title   string    `json:"title" bson:"title"`
//	Content string    `json:"content" bson:"content"`
//	Options *[]Option `json:"options" bson:"options"`
//}
//
//type Option struct {
//	Num     string `json:"num" bson:"num"`
//	Content string `json:"content" bson:"content"`
//	Answer  bool   `json:"answer" bson:"answer"`
//}
