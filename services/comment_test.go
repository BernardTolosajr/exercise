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

func TestCommentServiceCreateNewCommentEmptyResult(t *testing.T) {
	mock := repositories.CommentRepositoryMock{}

	comment := &models.Comment{
		Org:     "foo",
		Comment: "bar",
	}

	mock.On("Create", comment).Return(nil, nil)

	service := NewCommentService(&mock)

	id, _ := service.Create(comment)
	assert.Equal(t, "", id)
}

func TestCommentServiceCreateNewCommentWrongObjectId(t *testing.T) {
	mock := repositories.CommentRepositoryMock{}

	comment := &models.Comment{
		Org:     "foo",
		Comment: "bar",
	}

	mock.On("Create", comment).Return("some", nil)

	service := NewCommentService(&mock)

	id, _ := service.Create(comment)
	assert.Equal(t, "", id)
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

func TestCommentServiceGetAllWhenSuccessReturnArrayOfComment(t *testing.T) {
	org := "foo"

	mock := repositories.CommentRepositoryMock{}

	comments := []*models.Comment{&models.Comment{Comment: "foo"}}

	mock.On("GetAll", org).Return(comments, nil)

	service := NewCommentService(&mock)

	results, _ := service.GetAllBy(org)

	assert.Equal(t, 1, len(results))
}

func TestCommentServiceGetAllWhenFailedReturnEmptyArray(t *testing.T) {
	org := "foo"

	mock := repositories.CommentRepositoryMock{}

	comments := []*models.Comment{}

	mock.On("GetAll", org).Return(comments, nil)

	service := NewCommentService(&mock)

	results, _ := service.GetAllBy(org)

	assert.Equal(t, 0, len(results))
}

func TestCommentServiceGetAllWhenFailedReturnError(t *testing.T) {
	org := "foo"

	mock := repositories.CommentRepositoryMock{}

	mock.On("GetAll", org).Return([]*models.Comment{}, errors.New("ops"))

	service := NewCommentService(&mock)

	_, err := service.GetAllBy(org)

	assert.Equal(t, "ops", err.Error())
}
