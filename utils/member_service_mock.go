package utils

import (
	"github.com/exercise/models"
	"github.com/exercise/services"
	"github.com/stretchr/testify/mock"
)

type MemberServiceMock struct {
	mock.Mock
}

func (m *MemberServiceMock) GetAllBy(org string) ([]*services.MemberView, error) {
	args := m.Called(org)
	return args.Get(0).([]*services.MemberView), args.Error(1)
}

func (m *MemberServiceMock) Create(member *models.Member) (string, error) {
	args := m.Called(member)
	return args.Get(0).(string), args.Error(1)
}
