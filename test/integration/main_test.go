package integration

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/exercise/db"
	"github.com/exercise/models"
	"github.com/exercise/router"
	"github.com/gorilla/mux"
)

var r *mux.Router
var mongo *db.MongoDB

// Setup test and seed data
func TestMain(m *testing.M) {
	mongo = db.NewMongoDB("mongodb://localhost:27017", "test")

	seedsComments(mongo)

	r = router.New(mongo)

	m.Run()
}

func seedsComments(db *db.MongoDB) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	comment := &models.Comment{
		Org:     "foo",
		Comment: "test",
	}

	result, err := db.CommentCollection.InsertOne(ctx, comment)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("inserted %v", result)
}

func DropCommentCollection(db *db.MongoDB) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	db.CommentCollection.Drop(ctx)
}
