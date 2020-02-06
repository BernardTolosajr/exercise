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
