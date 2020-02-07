package repositories

import (
	"context"
	"log"
	"time"

	"github.com/exercise/db"
	"github.com/exercise/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMembersRepository interface {
	Create(comment *models.Member) (interface{}, error)
	GetAllby(org string) ([]*models.Member, error)
}

type MembersRepository struct {
	MongoDB *db.MongoDB
}

// Create member return interface and error
func (o *MembersRepository) Create(comment *models.Member) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	comment.DateCreated = time.Now()
	comment.DateModified = time.Now()

	result, err := o.MongoDB.MemberCollection.InsertOne(ctx, comment)

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

// GetallBy accept org paramter and return an array of member and error
func (o *MembersRepository) GetAllby(org string) ([]*models.Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var results []*models.Member

	options := options.Find()
	// for now set limit 100 by default
	// maybe in the future we can add pagination
	options.SetLimit(100)
	// Sort field -1 for descending
	options.SetSort(bson.D{{"followers", -1}})
	filter := orgAndFilter(org)

	cur, err := o.MongoDB.MemberCollection.Find(ctx, filter, options)

	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem models.Member
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
