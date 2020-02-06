package repositories

import (
	"context"
	"log"
	"time"

	"github.com/exercise/db"
	"github.com/exercise/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ICommentsRepository interface {
	Create(comment *models.Comment) (interface{}, error)
	DeleteAll(org string) (interface{}, error)
}

type CommentsRepository struct {
	MongoDB *db.MongoDB
}

func (o *CommentsRepository) Create(comment *models.Comment) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	comment.DateCreated = time.Now()
	comment.DateModified = time.Now()

	result, err := o.MongoDB.CommentCollection.InsertOne(ctx, comment)

	if err != nil {
		log.Printf("error on create %v", err)
		return nil, err
	}

	return result.InsertedID, nil
}

func (o *CommentsRepository) DeleteAll(org string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// filter only by org and datedeleted is empty
	// https://docs.mongodb.com/manual/reference/operator/query/and/
	filter := bson.D{
		{"$and", bson.A{bson.D{{"org", org}}, bson.D{{"datedeleted", ""}}}},
	}

	result, err := o.MongoDB.CommentCollection.UpdateMany(ctx, filter, bson.D{
		{"$set", bson.D{
			{"datedeleted", time.Now()},
		}},
	})

	if err != nil {
		return nil, err
	}

	return result.ModifiedCount, nil
}
