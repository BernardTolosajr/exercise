package services

import (
	"github.com/exercise/models"
	"github.com/stretchr/testify/mock"
)

type MemberServiceMock struct {
	mock.Mock
}

func (m *MemberServiceMock) GetAllBy(org string) ([]*MemberView, error) {
	args := m.Called(org)
	return args.Get(0).([]*MemberView), args.Error(1)
}

func (m *MemberServiceMock) Create(member *models.Member) (string, error) {
	args := m.Called(member)
	return args.Get(0).(string), args.Error(1)
}
