package user

//
//import (
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//const (
//	UserData = "userData"
//)
//
//type DataCollection struct {
//	col *mongo.Collection
//}
//
//func NewUserDataCollection(collection *mongo.Collection) *DataCollection {
//	return &DataCollection{
//		col: collection,
//	}
//}
//
//type Data struct {
//	Id       *primitive.ObjectID `json:"id", bson:"_id"`
//	Email    string              `json:"email" bson:"email"`
//	PassWord string              `json:"password" bson:"password"`
//	Name     string              `json:"name" bson:"name"`
//}
//
//func (c *DataCollection) GetData(email string) *Data {
//	data := new(Data)
//	filter := bson.M{"email": email}
//	project := bson.M{"password": 0}
//
//	r := c.col.FindOne(nil, filter, options.FindOne().SetProjection(project))
//	if err := r.Decode(data); err != nil && err == mongo.ErrNoDocuments {
//		return nil
//	} else if err != nil {
//		panic(err)
//	}
//
//	return data
//}
//
//func (c *DataCollection) GetDataWithHash(email, hash string) *Data {
//	data := new(Data)
//	filter := bson.M{"email": email, "password": hash}
//	project := bson.M{"password": 0}
//
//	r := c.col.FindOne(nil, filter, options.FindOne().SetProjection(project))
//	if err := r.Decode(data); err != nil && err == mongo.ErrNoDocuments {
//		return nil
//	} else if err != nil {
//		panic(err)
//	}
//
//	return data
//}
//
//func (c *DataCollection) CreateData(data *Data) error {
//	if _, err := c.col.InsertOne(nil, data); err != nil {
//		panic(err)
//	}
//
//	return nil
//}
