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
	mock := repositories.OrganizationRepositoryMock{}

	mock.On("Create", &models.Organization{
		Login: "login name",
	}).Return(primitive.NewObjectID(), nil)

	service := NewOrganization(&mock)

	id, _ := service.Create(&models.Organization{
		Login: "login name",
	})

	mock.AssertCalled(t, "Create", &models.Organization{
		Login: "login name",
	})

	assert.Equal(t, 24, len(id))
}

func TestServiceWillNotCreateId(t *testing.T) {
	mock := repositories.OrganizationRepositoryMock{}

	mock.On("Create", &models.Organization{
		Login: "login name",
	}).Return(nil, errors.New("some error found"))

	service := NewOrganization(&mock)

	id, _ := service.Create(&models.Organization{
		Login: "login name",
	})

	mock.AssertCalled(t, "Create", &models.Organization{
		Login: "login name",
	})
	assert.Equal(t, "", id)
}
