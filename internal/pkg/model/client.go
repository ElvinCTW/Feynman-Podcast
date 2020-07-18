package model

import (
	"feynman-podcast/internal/pkg/model/user"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModelClient struct {
	UserDataCollection *user.DataCollection
}

func NewClient(database, uri string) *ModelClient {
	if client, err := mongo.NewClient(options.Client().ApplyURI(uri)); err != nil {
		panic(err)
	} else {
		db := client.Database(database)
		return &ModelClient{
			UserDataCollection: user.NewCollection(db.Collection(user.UserData)),
		}
	}
}
