package model

import (
	"context"
	"feynman-podcast/internal/pkg/model/answer"
	"feynman-podcast/internal/pkg/model/comment"
	"feynman-podcast/internal/pkg/model/question"
	"feynman-podcast/internal/pkg/model/user"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
