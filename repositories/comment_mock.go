package repositories

import (
	"github.com/exercise/models"
	"github.com/stretchr/testify/mock"
)

type CommentRepositoryMock struct {
	mock.Mock
}

func (m *CommentRepositoryMock) Create(comment *models.Comment) (interface{}, error) {
	args := m.Called(comment)
	return args.Get(0), args.Error(1)
}

func (m *CommentRepositoryMock) DeleteAll(org string) (interface{}, error) {
	args := m.Called(org)
	return args.Get(0), args.Error(1)
}
