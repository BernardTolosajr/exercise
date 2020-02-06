package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client                 *mongo.Client
	OrganizationCollection *mongo.Collection
}

// New mongoDB setup new mongo client
func NewMongoDB(host string, database string) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(host)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// setup database and collections
	organizations := client.Database(database).Collection("organizations")

	return &MongoDB{
		Client:                 client,
		OrganizationCollection: organizations,
	}
}
