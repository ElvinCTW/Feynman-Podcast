package comment

//
//import (
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//const (
//	CommentData = "commentData"
//)
//
//type Data struct {
//	Id        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
//	AnswerId  *primitive.ObjectID `json:"answerId,omitempty" bson:"answerId,omitempty"`
//	Content   string              `json:"content" bson:"content"`
//	UserId    *primitive.ObjectID `json:"userId" bson:"userId"`
//	Likers    *[]string           `json:"likers" bson:"likers"`
//	LikeCount int                 `json:"likeCount" bson:"likeCount"`
//}
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
//func (c *DataCollection) CreateData(answerId string, userId string, data *Data) error {
//	answerObjectId, err := primitive.ObjectIDFromHex(answerId)
//	if err != nil {
//		return err
//	}
//
//	userObjectId, err := primitive.ObjectIDFromHex(userId)
//	if err != nil {
//		return err
//	}
//
//	likers := make([]string, 0)
//	d := &Data{
//		AnswerId:  &answerObjectId,
//		Content:   data.Content,
//		UserId:    &userObjectId,
//		Likers:    &likers,
//		LikeCount: 0,
//	}
//
//	if _, err := c.col.InsertOne(nil, d); err != nil {
//		panic(err)
//	}
//
//	return nil
//}
//
//func (c *DataCollection) ListData(answerId string) *[]Data {
//	answerObjectId, err := primitive.ObjectIDFromHex(answerId)
//	if err != nil {
//		return nil
//	}
//
//	filter := bson.M{"answerId": answerObjectId}
//	cur, err := c.col.Find(nil, filter)
//	if err != nil {
//		panic(err)
//	}
//
//	defer cur.Close(nil)
//
//	list := make([]Data, 0)
//	if err = cur.All(nil, &list); err != nil {
//		panic(err)
//	}
//
//	return &list
//}
//
//func (c *DataCollection) Updatelike(id, likerId string) error {
//	objectId, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return err
//	}
//
//	// check user in likers or not
//	filter := bson.M{"_id": objectId, "likers": bson.M{"$in": bson.A{likerId}}}
//	r := c.col.FindOne(nil, filter, options.FindOne().SetProjection(bson.M{"_id": 1}))
//	if r.Err() != nil && r.Err() != mongo.ErrNoDocuments {
//		panic(err)
//	}
//
//	// update likers and like
//	filter = bson.M{"_id": objectId}
//	noUpsert := options.Update().SetUpsert(false)
//	update := bson.M{"$inc": bson.M{"likeCount": -1}, "$pull": bson.M{"likers": likerId}}
//	if r.Err() == mongo.ErrNoDocuments {
//		update = bson.M{"$inc": bson.M{"likeCount": 1}, "$addToSet": bson.M{"likers": likerId}}
//	}
//
//	if _, err := c.col.UpdateOne(nil, filter, update, noUpsert); err != nil {
//		panic(err)
//	}
//
//	return nil
//}
//
//func (c *DataCollection) DeleteData(id, userId string) {
//	objectId, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return
//	}
//
//	userObjectId, err := primitive.ObjectIDFromHex(userId)
//	if err != nil {
//		return
//	}
//
//	filter := bson.M{"_id": objectId, "userId": userObjectId}
//	if _, err := c.col.DeleteOne(nil, filter); err != nil {
//		panic(err)
//	}
//}
