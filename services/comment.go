package services

import (
	"errors"

	"github.com/exercise/models"
	"github.com/exercise/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICommentService interface {
	Create(comment *models.Comment) (string, error)
}

type CommentService struct {
	CommentRepository repositories.ICommentsRepository
}

func NewCommentService(commentRepository repositories.ICommentsRepository) *CommentService {
	return &CommentService{
		CommentRepository: commentRepository,
	}
}

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
