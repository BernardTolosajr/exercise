package services

import (
	"errors"
	"testing"

	"github.com/exercise/models"
	"github.com/exercise/repositories"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCommentServiceCreateNewCommentWhenSuccess(t *testing.T) {
	mock := repositories.CommentRepositoryMock{}

	comment := &models.Comment{
		Org:     "foo",
		Comment: "bar",
	}

	mock.On("Create", comment).Return(primitive.NewObjectID(), nil)

	service := NewCommentService(&mock)

	id, _ := service.Create(comment)

	mock.AssertCalled(t, "Create", comment)

	assert.Equal(t, 24, len(id))
}

func TestCommentServiceCreateNewCommentWhenFailed(t *testing.T) {
	mock := repositories.CommentRepositoryMock{}

	comment := &models.Comment{
		Org:     "foo",
		Comment: "bar",
	}

	mock.On("Create", comment).Return(nil, errors.New("some error found"))

	service := NewCommentService(&mock)

	id, _ := service.Create(comment)

	mock.AssertCalled(t, "Create", comment)

	assert.Equal(t, "", id)
}
