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
