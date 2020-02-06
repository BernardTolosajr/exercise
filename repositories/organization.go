package repositories

import (
	"context"
	"log"
	"time"

	"github.com/exercise/db"
	"github.com/exercise/models"
)

type IOrganizationRepository interface {
	Create(organization *models.Organization) (interface{}, error)
}

type OrganizationRepository struct {
	MongoDB *db.MongoDB
}

func (o *OrganizationRepository) Create(organization *models.Organization) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := o.MongoDB.OrganizationCollection.InsertOne(ctx, organization)

	if err != nil {
		log.Printf("error on create %v", err)
		return nil, err
	}

	return result.InsertedID, nil
}
