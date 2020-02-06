package services

import (
	"errors"

	"github.com/exercise/models"
	"github.com/exercise/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrganizationService interface {
	Create(org *models.Organization) (string, error)
}

type OrganizationService struct {
	OrganizationRepository repositories.IOrganizationRepository
}

func NewOrganizationService(organizationRepository repositories.IOrganizationRepository) *OrganizationService {
	return &OrganizationService{
		OrganizationRepository: organizationRepository,
	}
}

func (o *OrganizationService) Create(org *models.Organization) (string, error) {
	result, err := o.OrganizationRepository.Create(org)

	if err != nil {
		return "", err
	}

	if result != nil {
		if oid, ok := result.(primitive.ObjectID); ok {
			return oid.Hex(), nil
		}
		return "", errors.New("Unable to cast ObjectId")
	}

	return "", nil
}
