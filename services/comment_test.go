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

func TestCommentServiceDeleteAllWhenSuccess(t *testing.T) {
	org := "foo"

	mock := repositories.CommentRepositoryMock{}

	mock.On("DeleteAll", org).Return(int64(1), nil)

	service := NewCommentService(&mock)

	result, _ := service.DeleteAll(org)

	assert.Equal(t, int64(1), result)
}

func TestCommentServiceDeleteAllWhenFailed(t *testing.T) {
	org := "foo"

	mock := repositories.CommentRepositoryMock{}

	mock.On("DeleteAll", org).Return(0, errors.New("ops"))

	service := NewCommentService(&mock)

	_, err := service.DeleteAll(org)

	assert.Equal(t, "ops", err.Error())
}
