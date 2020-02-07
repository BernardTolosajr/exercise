package repositories

import (
	"github.com/exercise/models"
	"github.com/stretchr/testify/mock"
)

type OrganizationRepositoryMock struct {
	mock.Mock
}

func (m *OrganizationRepositoryMock) Create(organization *models.Organization) (interface{}, error) {
	args := m.Called(organization)
	return args.Get(0), args.Error(1)
}

func (m *OrganizationRepositoryMock) FindOne(login string) (*models.Organization, error) {
	args := m.Called(login)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Organization), args.Error(1)
	}
	return nil, args.Error(1)
}
