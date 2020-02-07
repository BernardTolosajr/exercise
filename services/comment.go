package services

import (
	"errors"

	"github.com/exercise/models"
	"github.com/exercise/repositories"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICommentService interface {
	Create(comment *models.Comment) (string, error)
	DeleteAll(org string) (int64, error)
	GetAllBy(org string) ([]*Comment, error)
}

type CommentService struct {
	CommentRepository repositories.ICommentsRepository
}

// Comment view
type Comment struct {
	Comment string
}

func NewCommentService(commentRepository repositories.ICommentsRepository) *CommentService {
	return &CommentService{
		CommentRepository: commentRepository,
	}
}

// Get all by org return Array of comment and error
func (o *CommentService) GetAllBy(org string) ([]*Comment, error) {
	results, err := o.CommentRepository.GetAll(org)

	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}

	// return empty array instead of null
	if len(results) == 0 {
		return make([]*Comment, 0), nil
	}

	var comments []*Comment

	for _, model := range results {
		comments = append(comments, &Comment{
			Comment: model.Comment,
		})
	}

	return comments, nil
}

// Create comment return doc id and error
func (o *CommentService) Create(comment *models.Comment) (string, error) {
	result, err := o.CommentRepository.Create(comment)

	if err != nil {
		return "", err
	}

	if result != nil {
		if oid, ok := result.(primitive.ObjectID); ok {
			return oid.Hex(), nil
		}
		return "", errors.New("Unable to cast ObjectId")
	}

	return "", nil
}

// Delete all comment by org return total number of updated docs and error
func (o *CommentService) DeleteAll(org string) (int64, error) {
	result, err := o.CommentRepository.DeleteAll(org)

	if err != nil {
		log.Errorf("error on updating comment %v", err)
		return 0, err
	}

	return result.(int64), err
}
