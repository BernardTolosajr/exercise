package services

import (
	"github.com/exercise/models"
	"github.com/stretchr/testify/mock"
)

type CommentServiceMock struct {
	mock.Mock
}

func (o *CommentServiceMock) Create(comment *models.Comment) (string, error) {
	args := o.Called(comment)
	return args.Get(0).(string), args.Error(1)
}

func (o *CommentServiceMock) DeleteAll(org string) (int64, error) {
	args := o.Called(org)
	return args.Get(0).(int64), args.Error(1)
}

func (o *CommentServiceMock) GetAllBy(org string) ([]*Comment, error) {
	args := o.Called(org)
	return args.Get(0).([]*Comment), args.Error(1)
}
