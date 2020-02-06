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
	GetAll(org string) ([]*models.Comment, error)
}

type CommentsRepository struct {
	MongoDB *db.MongoDB
}

// Get all coments by or
func (o *CommentsRepository) GetAll(org string) ([]*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var results []*models.Comment

	filter := orgAndFilter(org)

	cur, err := o.MongoDB.CommentCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem models.Comment
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(ctx)

	return results, nil
}

// Create comment
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

// Delete all comment by org
func (o *CommentsRepository) DeleteAll(org string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// filter only by org and datedeleted is empty
	// https://docs.mongodb.com/manual/reference/operator/query/and/
	filter := orgAndFilter(org)

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

func orgAndFilter(org string) bson.D {
	return bson.D{
		{"$and", bson.A{bson.D{{"org", org}}, bson.D{{"datedeleted", ""}}}},
	}
}
