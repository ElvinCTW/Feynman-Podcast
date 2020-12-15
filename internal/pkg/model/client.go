package model

import (
	"cloud.google.com/go/firestore"
	"context"
	"feynman-podcast/internal/pkg/model/answer"
	"feynman-podcast/internal/pkg/model/comment"
	"feynman-podcast/internal/pkg/model/question"
	"feynman-podcast/internal/pkg/model/user"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once            = sync.Once{}
	fireStoreClient *FireStoreClient
)

type ModelClient struct {
	UserDataCollection *user.DataCollection
	QuestionCollection *question.DataCollection
	AnswerCollection   *answer.DataCollection
	CommentCollection  *comment.DataCollection
}

func NewClient(database, uri string) *ModelClient {
	if client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri)); err != nil {
		panic(err)
	} else {
		db := client.Database(database)
		return &ModelClient{
			UserDataCollection: user.NewUserDataCollection(db.Collection(user.UserData)),
			QuestionCollection: question.NewDataCollection(db.Collection(question.QuestionData)),
			AnswerCollection:   answer.NewDataCollection(db.Collection(answer.AnswerData)),
			CommentCollection:  comment.NewDataCollection(db.Collection(comment.CommentData)),
		}
	}
}

type FireStoreClient struct {
	firestore *firestore.Client
}

func NewFireStore(app *firebase.App) *FireStoreClient {
	once.Do(func() {
		ctx := context.Background()
		c, err := app.Firestore(ctx)
		if err != nil {
			fmt.Println("cannot init FireStore")
		}
		fireStoreClient = &FireStoreClient{firestore: c}
	})

	if fireStoreClient.firestore != nil {
		fmt.Println("init firestore success")
	}
	return fireStoreClient
}
