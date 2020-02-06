package repositories

import (
	"context"
	"log"
	"time"

	"github.com/exercise/db"
	"github.com/exercise/models"
	"go.mongodb.org/mongo-driver/bson"
)

type IOrganizationRepository interface {
	Create(organization *models.Organization) (interface{}, error)
	FindOne(login string) (*models.Organization, error)
}

type OrganizationRepository struct {
	MongoDB *db.MongoDB
}

func (o *OrganizationRepository) FindOne(login string) (*models.Organization, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result := &models.Organization{}

	err := o.MongoDB.OrganizationCollection.FindOne(ctx, bson.M{"login": login}).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (o *OrganizationRepository) Create(organization *models.Organization) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	organization.DateCreated = time.Now().String()
	organization.DateModified = time.Now().String()

	result, err := o.MongoDB.OrganizationCollection.InsertOne(ctx, organization)

	if err != nil {
		log.Printf("error on create %v", err)
		return nil, err
	}

	return result.InsertedID, nil
}
