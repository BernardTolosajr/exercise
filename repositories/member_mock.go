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

func (m *MemberRepositoryMock) GetAllby(org string) ([]*models.Member, error) {
	args := m.Called(org)
	return args.Get(0).([]*models.Member), args.Error(1)
}
