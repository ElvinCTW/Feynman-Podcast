package model

import (
	"context"
	"feynman-podcast/internal/pkg/model/question"
	firebase "firebase.google.com/go/v4"
	"fmt"
)

//type ModelClient struct {
//	UserDataCollection *user.DataCollection
//	QuestionCollection *question.DataCollection
//	AnswerCollection   *answer.DataCollection
//	CommentCollection  *comment.DataCollection
//}

//func NewClient(database, uri string) *ModelClient {
//	if client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri)); err != nil {
//		panic(err)
//	} else {
//		db := client.Database(database)
//		return &ModelClient{
//			UserDataCollection: user.NewUserDataCollection(db.Collection(user.UserData)),
//			QuestionCollection: question.NewDataCollection(db.Collection(question.QuestionData)),
//			AnswerCollection:   answer.NewDataCollection(db.Collection(answer.AnswerData)),
//			CommentCollection:  comment.NewDataCollection(db.Collection(comment.CommentData)),
//		}
//	}
//}

type FireStoreClient struct {
	QuestionCollection *question.QuestionCollection
}

func NewFireStore(app *firebase.App) *FireStoreClient {
	ctx := context.Background()
	c, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("cannot init FireStore")
		return nil
	}

	fmt.Println("init firestore success")
	return &FireStoreClient{
		QuestionCollection: question.NewQuestionCollection(c.Collection(question.CollectionName)),
	}
}
