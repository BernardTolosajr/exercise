package repositories

import (
	"github.com/exercise/models"
	"github.com/stretchr/testify/mock"
)

type MemberRepositoryMock struct {
	mock.Mock
}

// Create accept member model and return interface and error
func (m *MemberRepositoryMock) Create(comment *models.Member) (interface{}, error) {
	args := m.Called(comment)
	return args.Get(0), args.Error(1)
}
