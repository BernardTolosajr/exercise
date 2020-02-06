package router

import (
	"github.com/exercise/db"
	"github.com/exercise/handler"
	"github.com/exercise/repositories"
	"github.com/exercise/services"
	"github.com/gorilla/mux"
)

// New create instance of mux router with mongoDB
func New(mongoDB *db.MongoDB) *mux.Router {
	r := mux.NewRouter()

	organizationService := services.NewOrganizationService(&repositories.OrganizationRepository{
		MongoDB: mongoDB,
	})

	r.Handle("/api/organizations", handler.OrganizationHandler(organizationService))

	return r
}
