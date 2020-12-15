package question

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

const (
	CollectionName = "question"
	Domain         = "domain"
	Title          = "title"
	Content        = "content"
	Options        = "options"
)

type Question struct {
	Id    string `json:"id,omitempty" bson:"_id,omitempty"`
	Class struct {
		TestType string `json:"testType" bson:"testType"`
		Domain   string `json:"domain" bson:"domain"`
	} `json:"class,inline" bson:"class,inline"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	//Options *[]Option `json:"options" bson:"options"`
}

type InsertData map[string]interface{}

type QuestionCollection struct {
	col *firestore.CollectionRef
}

func NewQuestionCollection(col *firestore.CollectionRef) *QuestionCollection {
	return &QuestionCollection{col: col}
}

func (c *QuestionCollection) CreateQuestion(domain, title, content string) error {
	ctx := context.Background()
	q := InsertData{
		Domain:  domain,
		Title:   title,
		Content: content,
	}

	_, _, err := c.col.Add(ctx, q)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return err
	}

	return nil
}
