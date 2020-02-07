package services

import (
	"errors"
	"testing"

	"github.com/exercise/models"
	"github.com/exercise/repositories"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestServiceReturnCreatedId(t *testing.T) {
	login := "login name"
	mock := repositories.OrganizationRepositoryMock{}

	mock.On("FindOne", login).Return(nil, nil)
	mock.On("Create", &models.Organization{
		Login: login,
	}).Return(primitive.NewObjectID(), nil)

	service := NewOrganization(&mock)

	id, _ := service.Create(&models.Organization{
		Login: login,
	})

	mock.AssertCalled(t, "Create", &models.Organization{
		Login: login,
	})

	assert.Equal(t, 24, len(id))
}

func TestServiceReturnCreatedReturnError(t *testing.T) {
	login := "login name"
	mock := repositories.OrganizationRepositoryMock{}

	mock.On("FindOne", login).Return(nil, nil)
	mock.On("Create", &models.Organization{
		Login: login,
	}).Return(nil, errors.New("ops"))

	service := NewOrganization(&mock)

	_, err := service.Create(&models.Organization{
		Login: login,
	})

	assert.Equal(t, "ops", err.Error())
}

func TestServiceReturnCreatedReturnWrongObjectId(t *testing.T) {
	login := "login name"
	mock := repositories.OrganizationRepositoryMock{}

	mock.On("FindOne", login).Return(nil, nil)
	mock.On("Create", &models.Organization{
		Login: login,
	}).Return("1", nil)

	service := NewOrganization(&mock)

	id, _ := service.Create(&models.Organization{
		Login: login,
	})

	assert.Equal(t, "", id)
}

func TestServiceReturnAlreadExist(t *testing.T) {
	login := "login"
	mock := repositories.OrganizationRepositoryMock{}

	mock.On("FindOne", login).Return(&models.Organization{Login: login}, nil)

	mock.On("Create", &models.Organization{
		Login: login,
	}).Return(primitive.NewObjectID(), nil)

	service := NewOrganization(&mock)

	_, err := service.Create(&models.Organization{
		Login: login,
	})

	assert.Equal(t, "Organization already exist.", err.Error())
}

func TestServiceWillNotCreateId(t *testing.T) {
	login := "login name"
	mock := repositories.OrganizationRepositoryMock{}

	mock.On("FindOne", login).Return(nil, nil)
	mock.On("Create", &models.Organization{
		Login: login,
	}).Return(nil, errors.New("some error found"))

	service := NewOrganization(&mock)

	id, _ := service.Create(&models.Organization{
		Login: login,
	})

	mock.AssertCalled(t, "Create", &models.Organization{
		Login: login,
	})
	assert.Equal(t, "", id)
}
