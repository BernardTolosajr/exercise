package repositories

import (
	"context"
	"log"
	"time"

	"github.com/exercise/db"
	"github.com/exercise/models"
)

type ICommentsRepository interface {
	Create(comment *models.Comment) (interface{}, error)
}

type CommentsRepository struct {
	MongoDB *db.MongoDB
}

func (o *CommentsRepository) Create(comment *models.Comment) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	comment.DateCreated = time.Now().String()
	comment.DateModified = time.Now().String()

	result, err := o.MongoDB.CommentCollection.InsertOne(ctx, comment)

	if err != nil {
		log.Printf("error on create %v", err)
		return nil, err
	}

	return result.InsertedID, nil
}
