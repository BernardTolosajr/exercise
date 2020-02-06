package services

import (
	"github.com/exercise/models"
	"github.com/stretchr/testify/mock"
)

type OrganizationServiceMock struct {
	mock.Mock
}

func (o *OrganizationServiceMock) Create(org *models.Organization) (string, error) {
	args := o.Called(org)
	return args.Get(0).(string), args.Error(1)
}
