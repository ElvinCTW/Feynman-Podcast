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
	Id    string `json:"id,omitempty"`
	Class struct {
		TestType string `json:"testType"`
		Domain   string `json:"domain"`
	} `json:"class,inline"`
	Title   string `json:"title"`
	Content string `json:"content"`
	//Options *[]Option `json:"options" bson:"options"`
}

type InsertData map[string]interface{}

type QuestionCollection struct {
	col *firestore.CollectionRef
}

func NewQuestionCollection(col *firestore.CollectionRef) *QuestionCollection {
	return &QuestionCollection{col: col}
}

func (c *QuestionCollection) CreateQuestion(q *Question) (*string, error) {
	ctx := context.Background()

	doc := c.col.NewDoc()
	_, err := doc.Set(ctx, q)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return nil, err
	}

	return &doc.ID, nil
}

func (c *QuestionCollection) GetQuestion(id string) *Question {
	ctx := context.Background()
	dsnap, err := c.col.Doc(id).Get(ctx)
	if err != nil {
		return nil
	}

	q := new(Question)
	if err := dsnap.DataTo(q); err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	q.Id = dsnap.Ref.ID
	return q
}
